package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/xanderjakeq/letter_generator/cmd/server/views"
)

type inputRouter struct{}

func (rt inputRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rt.Landing)

	return r
}

func (rt inputRouter) Landing(w http.ResponseWriter, r *http.Request) {
	hx_request := r.Header.Get("HX-Request")
	if len(hx_request) > 0 {
		templ.Handler(views.Input()).ServeHTTP(w, r)
		return
	}
	templ.Handler(views.Index("/input")).ServeHTTP(w, r)
}
