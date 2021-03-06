package user

// UserRequest model
type UserRequest struct {
	UserName  string `json:"userName" bson:"userName,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
}
