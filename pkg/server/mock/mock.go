package mock

import (
	"net/http"

	"github.com/gorilla/mux"
	cw "github.com/techmexdev/competitive_writing_api"
	"github.com/unrolled/render"
)

type mock struct {
	store  cw.Storage
	router *mux.Router
	render *render.Render
}

// New creates an HTTP Server that always responds with hard-coded data.
func New(store cw.Storage) http.Handler {
	return &mock{store: store, router: mux.NewRouter(), render: render.New()}
}

func (m *mock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// cors.AllowAll().Handler(m.router)
	m.routes()
	m.router.ServeHTTP(w, r)
}

func (m *mock) routes() {
	m.router.HandleFunc("/signup", m.handleAuthSignup()).Methods("POST")
	m.router.HandleFunc("/login", m.handleAuthLogin()).Methods("POST")
	m.router.HandleFunc("/passage", m.handlePassages()).Methods("GET")
	m.router.HandleFunc("/analysis", m.handleAnalysiss()).Methods("GET")
	m.router.HandleFunc("/selection", m.handleSelections()).Methods("GET")
}

func (m *mock) handleAuthSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (m *mock) handleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (m *mock) handlePassages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pp, _ := m.store.Passage.List()
		m.render.JSON(w, http.StatusOK, pp)
	}
}

func (m *mock) handleAnalysiss() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		aa, _ := m.store.Analysis.List()
		m.render.JSON(w, http.StatusOK, aa)
	}
}

func (m *mock) handleSelections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss, _ := m.store.Selection.List()
		m.render.JSON(w, http.StatusOK, ss)
	}
}

func (m *mock) authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
