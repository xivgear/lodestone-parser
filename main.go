package main

import (
	"fmt"
	"github.com/xivgear/lodestone-parser/parser"
	"log"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var dataTransmitChan = make(chan string)
	defer close(dataTransmitChan)

	// Create worker in new go routine and give him the channel to push data to it
	w := parser.NewWorker(dataTransmitChan)
	defer func() {
		// Todo: This should go live when doing stuff later
		// sc := make(chan os.Signal, 1)
		// defer close(sc)
		// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		// <-sc
		w.Stop()
		log.Println("finished.")
	}()

	// testing worker foo. todo: remove :-)
	var characters = []string{"11756305", "21541412"}
	characters = []string{"11756305"}
	characters = []string{"9384803"}

	func() {
		for _, id := range characters {
			log.Println("id =", id)
			dataTransmitChan <- id
		}
	}()

	// Todo: This is for debug only! Remove or do only on debug flag...
	<-time.After(500 * time.Millisecond)

	return nil
}
