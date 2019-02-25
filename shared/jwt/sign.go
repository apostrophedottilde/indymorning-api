package jwt

import (
	"fmt"
	"log"

	jot "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Sign a JWT token based on the details of the authenticated user.
var Sign = func() string {
	// TODO: fetch user from database if creds match then sigh jwt based on id of fetched user
	byteHash := []byte("$2a$04$pb8RFF0v6KjRe72M6Na0jufpzbKQ4kLZdU9opYbX6NEhncApoA8BC")
	passwordOk := comparePasswords(string(byteHash), []byte("users-password"))

	if passwordOk == false {
		panic("password not matching bcrypted password")
	}
	fmt.Println("passwords must have matched")

	tk := &Token{UserId: 4356786}
	token := jot.NewWithClaims(jot.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(MySigningKey))
	return tokenString
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
