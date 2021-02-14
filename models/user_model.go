package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
