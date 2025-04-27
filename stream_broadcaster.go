package util

import (
	"io"
	"net/http"
	"sync"
)

type StreamBroadcaster struct {
	ContentType string
	StartFunc   func() (io.ReadCloser, error)

	clientsMu sync.Mutex
	clients   map[chan []byte]struct{}
	add       chan chan []byte
	remove    chan chan []byte
	broadcast chan []byte

	streaming bool
	stopCh    chan struct{}
}

func NewStreamBroadcaster(contentType string, startFunc func() (io.ReadCloser, error)) *StreamBroadcaster {
	sb := &StreamBroadcaster{
		ContentType: contentType,
		StartFunc:   startFunc,
		clients:     make(map[chan []byte]struct{}),
		add:         make(chan chan []byte),
		remove:      make(chan chan []byte),
		broadcast:   make(chan []byte, 64),
		stopCh:      make(chan struct{}),
	}
	go sb.run()
	return sb
}

func (sb *StreamBroadcaster) run() {
	for {
		select {
		case client := <-sb.add:
			sb.clientsMu.Lock()
			sb.clients[client] = struct{}{}
			startNeeded := !sb.streaming && len(sb.clients) == 1
			sb.clientsMu.Unlock()

			if startNeeded {
				go sb.startStreaming()
			}

		case client := <-sb.remove:
			sb.clientsMu.Lock()
			delete(sb.clients, client)
			close(client)
			shouldStop := sb.streaming && len(sb.clients) == 0
			sb.clientsMu.Unlock()

			if shouldStop {
				close(sb.stopCh)
				sb.stopCh = make(chan struct{})
			}

		case chunk := <-sb.broadcast:
			sb.clientsMu.Lock()
			for client := range sb.clients {
				select {
				case client <- chunk:
				default:
					delete(sb.clients, client)
					close(client)
				}
			}
			sb.clientsMu.Unlock()
		}
	}
}

func (sb *StreamBroadcaster) startStreaming() {
	r, err := sb.StartFunc()
	if err != nil {
		return
	}
	defer r.Close()

	sb.clientsMu.Lock()
	sb.streaming = true
	sb.clientsMu.Unlock()

	buf := make([]byte, 1024)
	for {
		select {
		case <-sb.stopCh:
			sb.clientsMu.Lock()
			sb.streaming = false
			sb.clientsMu.Unlock()
			return
		default:
			n, err := r.Read(buf)
			if err != nil {
				sb.clientsMu.Lock()
				sb.streaming = false
				sb.clientsMu.Unlock()
				return
			}
			chunk := make([]byte, n)
			copy(chunk, buf[:n])
			sb.broadcast <- chunk
		}
	}
}

func (sb *StreamBroadcaster) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", sb.ContentType)
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	client := make(chan []byte, 16)
	sb.add <- client
	defer func() { sb.remove <- client }()

	for chunk := range client {
		_, err := w.Write(chunk)
		if err != nil {
			break
		}
		flusher.Flush()
	}
}
