package jwt

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// Sign a JWT token based on the details of the authenticated user.
var Sign = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: fetch user from database if creds match then sigh jwt based on id of fetched user
		tk := &Token{UserId: 4356786}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(MySigningKey))
		_, _ = w.Write([]byte(tokenString))
	})
}
