package adapter

import (
	"fmt"
	"net/http"

	"github.com/apostrohedottilde/indymorning/api/project"
	"github.com/apostrohedottilde/indymorning/api/shared/middleware/jwt"
	l "github.com/apostrohedottilde/indymorning/api/shared/middleware/logger"
	t "github.com/apostrohedottilde/indymorning/api/shared/middleware/terminator"

	"github.com/gorilla/mux"
)

// HTTPAdapter implementation
type HTTPAdapter struct {
	router *mux.Router
}

// Start http adapter and listen for requests
func (adapter *HTTPAdapter) Start() {
	fmt.Println("Starting HTTP connection...")
	http.ListenAndServe(":8000", adapter.router)
}

// Stop http adapter
func (adapter *HTTPAdapter) Stop() {
	adapter.router = nil
}

// New creates a new instance of HTTPAdapter and returns a pointer to it.
func New(h *project.ProjectController) *HTTPAdapter {
	r := mux.NewRouter()

	r.HandleFunc("/auth/login", l.Log(jwt.Sign(t.End())).ServeHTTP).Methods("GET")

	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(h.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(h.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/projects", l.Log(jwt.Validate(h.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects", l.Log(jwt.Validate(h.Create(t.End()))).ServeHTTP).Methods("POST")
	r.HandleFunc("/projects", l.Log(jwt.Validate(h.FindAll(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(h.Delete(t.End()))).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/projects/{id}/cancel", l.Log(jwt.Validate(h.Cancel(t.End()))).ServeHTTP).Methods("POST")

	http.Handle("/", r)
	return &HTTPAdapter{
		router: r,
	}
}
