package user

import (
	"encoding/json"
	"fmt"
	"github.com/apostrophedottilde/go-forum-api/shared/jwt"
	"net/http"
	"strconv"
	"strings"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Controller struct {
	service service
}

func (rh *Controller) Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := jwt.Sign()
		buildResponse(w, 200)
		_, _ = w.Write([]byte(token))
	})
}

func (rh *Controller) Register(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request
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

func (rh *Controller) FindOne(next http.Handler) http.Handler {
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

func (rh *Controller) Update(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := userId(r)

		id := strings.TrimPrefix(r.URL.Path, "/projects/")
		// TODO: extract this data from the request instead of stubbing

		var req Request
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
func (rh *Controller) FindAll(next http.Handler) http.Handler {
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

func (rh *Controller) Delete(next http.Handler) http.Handler {
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

func (rh *Controller) Cancel(next http.Handler) http.Handler {
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

// NewController builds and returns a Controller
func NewController(s service) *Controller {
	return &Controller{
		service: s,
	}
}
