package adapter

import (
	"fmt"
	"github.com/apostrophedottilde/go-forum-api/shared/jwt"
	"net/http"

	"github.com/apostrophedottilde/go-forum-api/forum"
	"github.com/apostrophedottilde/go-forum-api/user"

	l "github.com/apostrophedottilde/go-forum-api/shared/middleware/logger"
	t "github.com/apostrophedottilde/go-forum-api/shared/middleware/terminator"

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
func New(u *user.Controller, p *forum.Controller) *HTTPAdapter {
	r := mux.NewRouter()

	r.HandleFunc("/auth/login", l.Log(u.Login(t.End())).ServeHTTP).Methods("POST")
	r.HandleFunc("/auth/register", u.Register(t.End()).ServeHTTP).Methods("POST")

	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/users", l.Log(jwt.Validate(u.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/users/{id}", l.Log(jwt.Validate(u.Delete(t.End()))).ServeHTTP).Methods("DELETE")

	r.HandleFunc("/forums/{id}", l.Log(jwt.Validate(p.FindOne(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/forums/{id}", l.Log(jwt.Validate(p.Update(t.End()))).ServeHTTP).Methods("PUT")
	r.HandleFunc("/forums", l.Log(jwt.Validate(p.FindAll(t.End()))).ServeHTTP).Methods("GET")
	r.HandleFunc("/forums", l.Log(jwt.Validate(p.Create(t.End()))).ServeHTTP).Methods("POST")
	r.HandleFunc("/forums/{id}", l.Log(jwt.Validate(p.Delete(t.End()))).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/forums/{id}/cancel", l.Log(jwt.Validate(p.Cancel(t.End()))).ServeHTTP).Methods("POST")

	http.Handle("/", r)
	return &HTTPAdapter{
		router: r,
	}
}
