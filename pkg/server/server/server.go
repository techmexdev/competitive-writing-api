package server

import (
	"net/http"

	"github.com/gorilla/mux"
	cw "github.com/techmexdev/competitive_writing_api"
)

type server struct {
	router  *mux.Router
	auth    *cw.Auth
	storage cw.Storage
}

// New creates an HTTP Server.
func New(mock bool) http.Handler {
	return &server{}
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
