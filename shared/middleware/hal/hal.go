package hal

import (
	"fmt"
	"net/http"
)

// HAL middleare to convert response to json/hal hateoas format
func SerializeResource(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Converted resource to hal")
		next.ServeHTTP(w, r)
	})
}

// HAL middleare to convert response to json/hal hateoas format
func SerializeCollection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Converted collection to hal")
		next.ServeHTTP(w, r)
	})
}
