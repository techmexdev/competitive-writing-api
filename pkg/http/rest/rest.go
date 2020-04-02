package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/techmexdev/competitive_writing_api/pkg/auth"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
	"github.com/unrolled/render"
)

type authSvc auth.Service
type psgSvc passage.Service
type usrSvc user.Service
type wriSvc writing.Service
type selSvc selecting.Service

type server struct {
	authSvc
	psgSvc
	usrSvc
	wriSvc
	selSvc
	router *mux.Router
	render *render.Render
	logger
}

// Config contains app's services
type Config struct {
	Auth auth.Service
	Psg  passage.Service
	Usr  user.Service
	Wri  writing.Service
	Sel  selecting.Service
}

// New creates an HTTP Server that always responds with hard-coded data.
func New(config Config, lgr *log.Logger) http.Handler {
	s := &server{
		authSvc: config.Auth,
		wriSvc:  config.Wri,
		usrSvc:  config.Usr,
		selSvc:  config.Sel,
		logger:  logger{lgr},
		router:  mux.NewRouter(),
		render:  render.New(),
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router.HandleFunc("/health", trace(cors(logRequest(s.logger, health)))).Methods("GET")
	s.router.HandleFunc("/passage", s.handlePassages()).Methods("GET")
	s.router.HandleFunc("/writing", s.handleWritings()).Methods("GET")
	s.router.HandleFunc("/selection", s.handleSelections()).Methods("GET")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *server) writeJSON(w http.ResponseWriter, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed encoding %#v into JSON: %s", data, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
	return nil
}

func writeHeaderAndText(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Add("Content-Type", "text")
	w.WriteHeader(statusCode)
	w.Write([]byte(http.StatusText(statusCode) + " " + msg))
}

func (s *server) serverError(w http.ResponseWriter, r *http.Request, err error) {
	writeHeaderAndText(w, http.StatusInternalServerError, "")
	s.logger.serverError(r, err)
}

func (s *server) handlePassages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pp, err := s.ListPassages()
		if err != nil {
			s.serverError(w, r, err)
			return
		}

		err = s.writeJSON(w, pp)
		if err != nil {
			s.serverError(w, r, fmt.Errorf("failed writing JSON into response: %s", err))
			return
		}
	}
}

func (s *server) handleWritings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) handleSelections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
