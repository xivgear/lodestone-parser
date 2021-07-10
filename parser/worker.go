package parser

import (
	"github.com/xivgear/lodestone-parser/parser/character"
	"log"
	"sync"
)

type Worker struct {
	dataChan  chan string
	waitGroup sync.WaitGroup
}

func (w *Worker) Listen() {
	go func() {
		for {
			select {
			case item := <-w.dataChan:
				w.waitGroup.Add(1)
				log.Println("received data", item)
				charParser := character.NewCharacter(item)
				if err := charParser.ParseCharacterData(); err != nil {
					log.Println(err)
				}
				w.waitGroup.Done()
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Println("Stopping...")
	w.waitGroup.Wait()
}

func NewWorker(data chan string) *Worker {
	w := &Worker{
		dataChan: data,
	}
	w.Listen()
	return w
}
