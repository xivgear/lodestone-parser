package main

import (
	"encoding/json"
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
	// Create worker which will spawn a new go routine
	w := parser.NewWorker()

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

	func() {
		for _, id := range characters {
			log.Println("id =", id)
			w.DataChan <- id
		}
	}()

	for {
		select {
		case character := <-w.ReturnChan:
			bytes, err := json.Marshal(character)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(bytes))
		}
	}

	// Todo: This is for debug only! Remove or do only on debug flag...
	<-time.After(1 * time.Second)

	return nil
}
