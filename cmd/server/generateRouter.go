package main

import (
	"fmt"
	"net/http"
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
	helpers.ReadInput(&letters, input)

	fmt.Println(len(letters))
	fmt.Println(len(letters[0].Template_file_name))

	var wg sync.WaitGroup
	for _, letter := range letters {
		wg.Add(1)
		go func(l *l.Letter) {
			defer wg.Done()
			l.Generate()
		}(&letter)
	}

	wg.Wait()

	templ.Handler(views.Generate()).ServeHTTP(w, r)
}
