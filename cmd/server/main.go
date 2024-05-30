package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"letter_generator/cmd/server/views"
)

func main() {

	mode := "prod"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	port := ":0"
	if mode == "dev" {
		port = ":3000"
		if len(os.Args) == 3 {
			port = fmt.Sprintf(":%s", os.Args[2])
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(fmt.Sprintf("can't bind to port: %s", port))
	}
	defer l.Close()

	r.Get("/templ", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Hello("new testingg")).ServeHTTP(w, r)
	})

	addr := strings.Split(l.Addr().String(), ":")[3]

	fmt.Printf("listening at port: %s", addr)
	http.Serve(l, r)
}
