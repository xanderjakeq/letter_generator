package main

import l "letter_generator/pkg/letter"
import "letter_generator/pkg/helpers"

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	letters := make([]l.Letter, 0)

	if len(os.Args) > 1 {
		if os.Args[1] == "serve" {
			log.Print("running server on")
		} else {
			input_path := os.Args[1]

			byte, err := os.ReadFile(input_path)

			if err != nil {
				log.Fatal(err)
			}

			helpers.ReadInput(&letters, byte)
		}
	} else {
		log.Fatal("input file path required. example: './input.txt'")
	}

	var output_path string
	var wg sync.WaitGroup

	for _, letter := range letters {
		wg.Add(1)
		go func(l *l.Letter) {
			defer wg.Done()
			output_path, _ = l.Generate()
		}(&letter)
	}

	wg.Wait()

	cmd := exec.Command("open", fmt.Sprintf("%s", output_path))
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
