package server

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/login", s.handleAuthLogin()).Methods("POST")
	s.router.HandleFunc("/signup", s.handleAuthSignup()).Methods("POST")
}

func (s *server) handleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) handleAuthSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) handleAnalysisList() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *server) authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
