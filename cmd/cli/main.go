package main

import l "letter_generator/pkg/letter"
import "letter_generator/pkg/helpers"

import (
	"log"
	"os"
	"sync"
)

func main() {
	letters := make([]l.Letter, 0)

	if len(os.Args) > 1 {
		if os.Args[1] == "serve" {
            log.Print("running server on")
		} else {
			helpers.ReadFileInput(&letters)
		}
	} else {
		log.Fatal("input file path required. example: './input.txt'")
	}

	var wg sync.WaitGroup

	for _, letter := range letters {
		wg.Add(1)
		go func(l *l.Letter) {
			defer wg.Done()
			l.Generate()
		}(&letter)
	}

	wg.Wait()
}