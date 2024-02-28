package util

import "github.com/fsnotify/fsnotify"

func Watch(path string, fn func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-watcher.Events:
				fn()
			case err := <-watcher.Errors:
				panic(err)
			}
		}
	}()
	if err := watcher.Add(path); err != nil {
		panic(err)
	}
	<-done
}
