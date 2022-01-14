package parser

import (
	"github.com/xivgear/lodestone-parser/parser/character"
	"log"
	"sync"
)

type Worker struct {
	DataChan  chan string
	ReturnChan chan *character.Character
	waitGroup sync.WaitGroup
}

func (w *Worker) Listen() {
	go func() {
		defer close(w.DataChan)
		defer close(w.ReturnChan)
		for {
			select {
			case item := <-w.DataChan:
				w.waitGroup.Add(1)
				log.Println("received data", item)
				character := character.NewCharacter(item)
				if err := character.ParseCharacterData(); err != nil {
					log.Println(err)
				}
				w.ReturnChan <-character

				w.waitGroup.Done()
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Println("Stopping...")
	w.waitGroup.Wait()
}

func NewWorker() *Worker {
	var dataChan = make(chan string)
	var returnChan = make(chan *character.Character)
	w := &Worker{
		DataChan: dataChan,
		ReturnChan: returnChan,
	}
	w.Listen()
	return w
}
