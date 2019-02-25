package user

import (
	"encoding/json"
	"fmt"
	"github.com/apostrohedottilde/indymorning/api/shared"
	"net/http"
	"strconv"
	"strings"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UserController struct {
	service Service
}

func (rh *UserController) Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := shared.Sign()
		buildResponse(w, 200)
		_, _ = w.Write([]byte(token))
	})
}

func (rh *UserController) Register(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req UserRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			panic("could not convert post data to request struct")
		}

		newUserStub := User{
			UserName:  req.UserName,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Password:  req.Password,
		}

		res, err := rh.service.Register(newUserStub)

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

func (rh *UserController) FindOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := userId(r)

		id := strings.TrimPrefix(r.URL.Path, "/projects/")
		res, err := rh.service.FindOne(userID, id)

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

func (rh *UserController) Update(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := userId(r)

		id := strings.TrimPrefix(r.URL.Path, "/projects/")
		// TODO: extract this data from the request instead of stubbing

		var req UserRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)

		if err != nil {
			panic("could not convert post data to request struct")
		}

		newUserStub := User{
			UserName:  req.UserName,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Password:  req.Password,
		}

		res, err := rh.service.Update(userID, id, newUserStub)

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
func (rh *UserController) FindAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := userId(r)

		res, err := rh.service.FindAll(userID)

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

func (rh *UserController) Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := userId(r)

		id := strings.TrimPrefix(r.URL.Path, "/projects/")

		err := rh.service.Delete(userID, id)

		if err != nil {
			panic(err)
		}

		buildResponse(w, 200)
		_, _ = w.Write([]byte(""))
		next.ServeHTTP(w, r)
	})
}

func (rh *UserController) Cancel(next http.Handler) http.Handler {
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

// userId of current logged in user
func userId(r *http.Request) string {
	return fmt.Sprintf("%v", r.Context().Value("user"))
}

// NewController builds and returns a UserController
func NewController(s Service) *UserController {
	return &UserController{
		service: s,
	}
}
