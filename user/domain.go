package user

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type user interface{}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName  string             `json:"userName" bson:"userName,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName" bson:"lastName,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
}
