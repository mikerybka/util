package util

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func NewRecursiveWatcher(dir string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			fmt.Println(path)
			watcher.Add(path)
		}
		return nil
	})
	return watcher, nil
}

func Watch(path string, fn func()) error {
	watcher, err := NewRecursiveWatcher(path)
	if err != nil {
		return err
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case e := <-watcher.Events:
				if e.Has(fsnotify.Create) ||
					e.Has(fsnotify.Write) ||
					e.Has(fsnotify.Remove) ||
					e.Has(fsnotify.Rename) {
					fn()
				}
			case err := <-watcher.Errors:
				fmt.Println("Watcher error:", err)
			}
		}
	}()
	<-done
	return nil
}

// func WatchDir(path string, fn func()) {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer watcher.Close()
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case <-watcher.Events:
// 				fn()
// 			case err := <-watcher.Errors:
// 				panic(err)
// 			}
// 		}
// 	}()
// 	if err := watcher.Add(path); err != nil {
// 		panic(err)
// 	}
// 	<-done
// }
