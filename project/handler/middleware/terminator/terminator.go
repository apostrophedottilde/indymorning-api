package terminator

import (
	"fmt"
	"net/http"
)

// End middleare for http reqest
func End() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Returning response and terminating")
	})
}
