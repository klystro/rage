package watcher

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	debounceDelay = 200 * time.Millisecond
	mutex         sync.Mutex
	lastEvent     time.Time
)

func StartWatching(rootDir, entryPoint string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				handleEvent(event, entryPoint)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func handleEvent(event fsnotify.Event, entryPoint string) {
	if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
		now := time.Now()
		mutex.Lock()
		if now.Sub(lastEvent) < debounceDelay {
			mutex.Unlock()
			return
		}
		lastEvent = now
		mutex.Unlock()

		log.Printf("File modified: %s", event.Name)
		restart(entryPoint)
	}
}

func restart(entryPoint string) {
	log.Println("Restarting application...")

	cmd := exec.Command("go", "run", entryPoint)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	go func() {
		err = cmd.Wait()
		if err != nil {
			log.Printf("Application exited with error: %v", err)
		}
	}()
}
