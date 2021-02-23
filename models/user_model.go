package models

import "time"

// User struct
type User struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserName     string    `json:"username"`
	Email        string    `json:"email" validate:"email,required"`
	Password     string    `json:"password" validate:"required,min=6"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
