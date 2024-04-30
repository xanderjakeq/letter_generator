package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	letters := make([]Letter, 0)

	if len(os.Args) > 1 {
		readFileInput(&letters)
	} else {
		log.Fatal("input file path required. example: './input.txt'")
	}

	var wg sync.WaitGroup

	for _, letter := range letters {
		wg.Add(1)
		go func(l *Letter) {
			defer wg.Done()
			l.Generate()
		}(&letter)
	}

	wg.Wait()
}
