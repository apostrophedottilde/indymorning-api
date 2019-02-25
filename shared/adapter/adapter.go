package adapter

import (
	"fmt"
	"github.com/apostrohedottilde/indymorning/api/shared"
	"net/http"

	"github.com/apostrohedottilde/indymorning/api/project"
	"github.com/apostrohedottilde/indymorning/api/user"

	l "github.com/apostrohedottilde/indymorning/api/shared/middleware/logger"
	t "github.com/apostrohedottilde/indymorning/api/shared/middleware/terminator"

	"github.com/gorilla/mux"
)

// HTTPAdapter implementation
type HTTPAdapter struct {
	router *mux.Router
}

// Start http adapter and listen for requests
func (adapter *HTTPAdapter) Start() error {
	fmt.Println("Starting HTTP connection...")
	err := http.ListenAndServe(":8000", adapter.router)

	if err != nil {
		return fmt.Errorf("error starting server")
	}
	return nil
}

// Stop http adapter
func (adapter *HTTPAdapter) Stop() {
	adapter.router = nil
}

// New creates a new instance of HTTPAdapter and returns a pointer to it.
func New(u *user.UserController, p *project.ProjectController) *HTTPAdapter {
	r := mux.NewRouter()

	r.HandleFunc("/auth/login", l.Log(u.Login(t.End())).ServeHTTP).Methods("POST")
	r.HandleFunc("/auth/register", u.Register(t.End()).ServeHTTP).Methods("POST")

	r.HandleFunc("/users/{id}", l.Log(shared.Validate(u.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(shared.Validate(u.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/users", l.Log(shared.Validate(u.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(shared.Validate(u.Delete(t.End()))).ServeHTTP).Methods("DELETE")

	r.HandleFunc("/projects/{id}", l.Log(shared.Validate(p.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects/{id}", l.Log(shared.Validate(p.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/projects", l.Log(shared.Validate(p.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects", l.Log(shared.Validate(p.Create(t.End()))).ServeHTTP).Methods("POST")
	r.HandleFunc("/projects/{id}", l.Log(shared.Validate(p.Delete(t.End()))).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/projects/{id}/cancel", l.Log(shared.Validate(p.Cancel(t.End()))).ServeHTTP).Methods("POST")

	http.Handle("/", r)
	return &HTTPAdapter{
		router: r,
	}
}
