package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/apostrohedottilde/indymorning/api/project/domain"

	s "github.com/apostrohedottilde/indymorning/api/project/service"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Handler interface {
	FindOne(handler *Handler)
	FindAll(handler *Handler)
	Cancel(handler *Handler)
}

type RequestHandler struct {
	service s.Service
}

func (rh *RequestHandler) Create(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: extract this data from the request instead of stubbing
		newProjectStub := domain.GameProject{
			Name:            "Tokyo job hunter - adventure game",
			BriefSynopsis:   "This is an adventure game about moving trying to move to Tokyo",
			FullDescription: "Enter the exciting world of applying for jobs as a software developer in Tokyo with the amzingly fun adventure simulator game!",
			Contributors:    []string{"user1", "user2", "user3"},
			State:           "LIVE",
		}

		res, err := rh.service.Create(newProjectStub)

		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}

		buildResponse(w, 200)
		_, _ = w.Write([]byte(data))
		next.ServeHTTP(w, r)
	})
}

func (rh *RequestHandler) FindOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/projects/")
		res, err := rh.service.FindOne(id)

		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}

		buildResponse(w, 200)
		_, _ = w.Write([]byte(data))
		next.ServeHTTP(w, r)
	})
}

func (rh *RequestHandler) Update(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/projects/")
		// TODO: extract this data from the request instead of stubbing

		newProjectStub := domain.GameProject{
			Name:            "Tokyo job hunter - adventure game",
			BriefSynopsis:   "This is an adventure game about moving trying to move to Tokyo",
			FullDescription: "Enter the exciting world of applying for jobs as a software developer in Tokyo with the amzingly fun adventure simulator game!",
			Contributors:    []string{"user1", "user2", "user3"},
			State:           "LIVE",
		}

		res, err := rh.service.Update(id, newProjectStub)

		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}

		buildResponse(w, 200)
		_, _ = w.Write([]byte(data))
		next.ServeHTTP(w, r)
	})
}

// FindAll returns a HTTPHandler function that carries out the logic of this request
func (rh *RequestHandler) FindAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := rh.service.FindAll()

		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}
		fmt.Println("found all")
		buildResponse(w, 200)
		_, _ = w.Write([]byte(data))
		next.ServeHTTP(w, r)
	})
}

func (rh *RequestHandler) Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/projects/")

		err := rh.service.Delete(id)

		if err != nil {
			panic(err)
		}

		buildResponse(w, 200)
		_, _ = w.Write([]byte(""))
		next.ServeHTTP(w, r)
	})
}

func (rh *RequestHandler) Cancel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func buildResponse(resp http.ResponseWriter, status int) http.ResponseWriter {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(status)
	return resp
}

func errResponse(resp http.ResponseWriter, code string, message string) {
	err := ErrResponse{
		Code:    code,
		Message: message,
	}
	data, e := json.Marshal(err)

	if e != nil {
		panic(err)
	}

	d, e := strconv.Atoi(code)

	if e != nil {
		panic(err)
	}

	resp.WriteHeader(d)
	resp.Write([]byte(data))
}

// New builds and returns a RequestHandler
func New(s s.Service) *RequestHandler {
	return &RequestHandler{
		service: s,
	}
}
