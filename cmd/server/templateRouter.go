package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"letter_generator/cmd/server/views"
	"letter_generator/pkg/helpers"
)

type templateRouter struct{}

func (rt templateRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rt.Landing)
	r.Get("/opendir", rt.OpenDir)

	return r
}

func (rt templateRouter) Landing(w http.ResponseWriter, r *http.Request) {
	hx_request := r.Header.Get("HX-Request")
	if len(hx_request) == 0 {
		templ.Handler(views.Index("/template")).ServeHTTP(w, r)
		return
	}

	cwd, err := helpers.GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	dir_name := fmt.Sprintf("%s/templates", cwd)
	files, err := os.ReadDir(dir_name)

	if err != nil {
		log.Fatal(err)
	}

	var templates []string

	for _, file := range files {
		templates = append(templates, file.Name())
	}

	templ.Handler(views.Template(&dir_name, &templates)).ServeHTTP(w, r)
}

func (rt templateRouter) OpenDir(w http.ResponseWriter, r *http.Request) {
	cwd, err := helpers.GetRootDir()

	if err != nil {
		templ.Handler(views.Error(err.Error())).ServeHTTP(w, r)
		return
	}

	cmd := exec.Command("open", fmt.Sprintf("%s/templates", cwd))
	err = cmd.Run()

	if err != nil {
		templ.Handler(views.Error(err.Error())).ServeHTTP(w, r)
		return
	}

	http.Redirect(w, r, "/template", http.StatusTemporaryRedirect)
}
