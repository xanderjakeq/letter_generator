package main

import (
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
	templ.Handler(views.Template()).ServeHTTP(w, r)
}
