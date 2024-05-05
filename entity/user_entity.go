package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UpdateUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
