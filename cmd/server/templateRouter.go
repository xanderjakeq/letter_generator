package main

import (
    "os"
    "log"
    "strings"
    "fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"letter_generator/cmd/server/views"
)

type templateRouter struct{}

func (rt templateRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rt.Landing)

	return r
}

func (rt templateRouter) Landing(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

    cwd_arr := strings.Split(cwd, "/")
    cwd = strings.Join(cwd_arr[:len(cwd_arr) - 2], "/")

	dir_name := fmt.Sprintf("%s/templates", cwd)
    files, err := os.ReadDir(dir_name)

	if err != nil {
		log.Fatal(err)
	}

    var templates []string

    for _, file := range files {
        templates = append(templates, file.Name())
    }

	templ.Handler(views.Template(&templates)).ServeHTTP(w, r)
}
