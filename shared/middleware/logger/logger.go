package logger

import (
	"fmt"
	"net/http"
)

// Log middleare for http reqest
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("USED LOGGING MIDDLEWARE")
		next.ServeHTTP(w, r)
	})
}
