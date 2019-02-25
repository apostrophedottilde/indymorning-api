package adapter

import (
	"fmt"
	"github.com/apostrophedottilde/indymorning-api/shared/jwt"
	"net/http"

	"github.com/apostrophedottilde/indymorning-api/project"
	"github.com/apostrophedottilde/indymorning-api/user"

	l "github.com/apostrophedottilde/indymorning-api/shared/middleware/logger"
	t "github.com/apostrophedottilde/indymorning-api/shared/middleware/terminator"

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

	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/users", l.Log(jwt.Validate(u.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.Delete(t.End()))).ServeHTTP).Methods("DELETE")

	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(p.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(p.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/projects", l.Log(jwt.Validate(p.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/projects", l.Log(jwt.Validate(p.Create(t.End()))).ServeHTTP).Methods("POST")
	r.HandleFunc("/projects/{id}", l.Log(jwt.Validate(p.Delete(t.End()))).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/projects/{id}/cancel", l.Log(jwt.Validate(p.Cancel(t.End()))).ServeHTTP).Methods("POST")

	http.Handle("/", r)
	return &HTTPAdapter{
		router: r,
	}
}
