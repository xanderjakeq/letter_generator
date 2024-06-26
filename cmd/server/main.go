package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/xanderjakeq/letter_generator/cmd/server/views"
	"github.com/xanderjakeq/letter_generator/pkg/helpers"
)

func main() {

	mode := "dev"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	port := ":3000"
	if mode == "prod" {
		port = ":0"
		if len(os.Args) == 3 {
			port = fmt.Sprintf(":%s", os.Args[2])
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(fmt.Sprintf("can't bind to port: %s", port))
	}
	defer l.Close()

	workDir, _ := helpers.GetRootDir()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/s", filesDir)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index("/input")).ServeHTTP(w, r)
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		hx_request := r.Header.Get("HX-Request")
		if len(hx_request) > 0 {
			templ.Handler(views.About()).ServeHTTP(w, r)
			return
		}
		templ.Handler(views.Index("/about")).ServeHTTP(w, r)
	})

	r.Mount("/input", inputRouter{}.Routes())
	r.Mount("/template", templateRouter{}.Routes())
	r.Mount("/generate", generateRouter{}.Routes())

	addr := strings.Split(l.Addr().String(), ":")[3]

	fmt.Printf("listening at port: localhost:%s\n", addr)
	http.Serve(l, r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
