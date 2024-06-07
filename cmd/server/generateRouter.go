package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"letter_generator/cmd/server/views"
	"letter_generator/pkg/helpers"
	l "letter_generator/pkg/letter"
)

type generateRouter struct{}

func (rt generateRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", rt.Generate)

	return r
}

func (rt generateRouter) Generate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := []byte(r.Form.Get("input"))

	letters := make([]l.Letter, 0)
	err := helpers.ReadInput(&letters, input)

	if err != nil {
		templ.Handler(views.Error("incomplete input")).ServeHTTP(w, r)
	} else {

		var output_path string

		var wg sync.WaitGroup
		for _, letter := range letters {
			wg.Add(1)
			go func(l *l.Letter) {
				defer wg.Done()
				output_path = l.Generate()
			}(&letter)
		}

		wg.Wait()

		cmd := exec.Command("open", fmt.Sprintf("%s", output_path))
		err = cmd.Run()

		if err != nil {
			log.Fatal(err)
		}

		templ.Handler(views.Generate()).ServeHTTP(w, r)
	}
}
