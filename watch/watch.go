package watch

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/inotify"
)

func Watch(dir string, fn func(path string) (string, error)) {
	watch(dir, fn)
}

func watch(dir string, fn func(path string) (string, error)) {
	if !hasDir(dir) {
		err := errors.New("Directory does not exist.")
		log.Fatal(err)
		return
	}

	// `/tmp/transporter` is a reserved path, allowing
	// this would cause an infinite loop on file addition
	if dir == "/tmp/transporter" {
		err := errors.New("Directory is reserved.")
		log.Fatal(err)
		return
	}

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Watching", dir)
	watcher.Watch(dir)

	for {
		select {
		case ev := <-watcher.Event:
			if ev.Mask == inotify.IN_CLOSE_WRITE {
				_, err = fn(ev.Name)
				if err != nil {
					log.Println(err)
				}
			}
		case err := <-watcher.Error:
			fmt.Println("error:", err)
		}
	}
}

func hasDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}
