package app

import (
	"log"
	"net/http"

	"github.com/personal/blog-app/config"
	"go.uber.org/dig"
)

// Server represents the app
type Server struct {
	router      Router
	controllers []Controller
	config      *config.Config
}

// ServerDependencies describes dependencies for the server
type ServerDependencies struct {
	dig.In
	Router      Router
	Controllers []Controller `group:"controller"`
	Config      *config.Config
}

// NewServer creates a new server
func NewServer(container *dig.Container) *Server {
	s := &Server{}

	s.initDependencies(container)
	s.initRoutes()

	return s
}

// initDependencies initializes all needed controllers
func (s *Server) initDependencies(container *dig.Container) {
	err := container.Invoke(
		func(dependencies ServerDependencies) {
			s.router = dependencies.Router
			s.controllers = dependencies.Controllers
			s.config = dependencies.Config
		})
	if err != nil {
		panic(err)
	}
}

// initRoutes assigning routes to controllers
func (s *Server) initRoutes() {
	// register all controllers
	for _, c := range s.controllers {
		c.Register(s.router)
	}
}

// Serve listens on the TCP network port. To allow HTTPS connections
// files containing a certificate and matching private key for the server must be provided.
func (s *Server) Serve(port string) error {
	log.Printf("Listening HTTP on port %s...", port)

	return http.ListenAndServe(":"+port, s.router)
}
