package main

import (
	"encoding/json"
	"fmt"
	"github.com/xivgear/lodestone-parser/parser"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	p := parser.NewParser()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			if r.URL.Query().Get("lodestoneId") == "" {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("400_StatusBadRequest"))
				return
			}

			if err := p.ParseCharacter(r.URL.Query().Get("lodestoneId")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("500_StatusInternalServerError"))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			data, _ := json.MarshalIndent(p.Character, "", "\t")
			w.Write(data)
			// For debugging purpose,
		} else if r.Method == "HEAD" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("405_StatusMethodNotAllowed"))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
