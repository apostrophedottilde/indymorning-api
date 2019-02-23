package jwt

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

// Sign a JWT token based on the details of the authenticated user.
var Sign = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: fetch user from database if creds match then sigh jwt based on id of fetched user
		byteHash := []byte("$2a$04$pb8RFF0v6KjRe72M6Na0jufpzbKQ4kLZdU9opYbX6NEhncApoA8BC")
		passwordOk := comparePasswords(string(byteHash), []byte("users-password"))

		if passwordOk == false {
			panic("password not matching bcrypted password")
		}
		fmt.Println("passwords must have matched")

		tk := &Token{UserId: 4356786}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(MySigningKey))
		_, _ = w.Write([]byte(tokenString))
	})
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
