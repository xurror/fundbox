package server

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"net/http"
)

type Server struct {
	router chi.Router
	config *Config
}

func NewServer(c *Config) (*Server, error) {
	return &Server{
		//router: NewRouter(&RouterConfig{
		//	DisableAuth: c.DisableAuth,
		//}),
		router: NewRouter(),
		config: c,
	}, nil
}

var routes = flag.Bool("routes", false, "Generate router documentation")

func (s *Server) Run() {
	r := s.router

	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	if *routes {
		fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		return
	}

	fmt.Printf("Starting server on %v\n", s.config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), r)
}
