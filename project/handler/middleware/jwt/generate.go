package jwt

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var Generate = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tk := &Token{UserId: 4356786}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(MySigningKey))
		fmt.Println(tokenString)
	})
}
