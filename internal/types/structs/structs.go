package structs

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"user_id"    db:"user_id"    validate:"-"`
	FirstName string    `json:"first_name" db:"first_name" validate:"omitempty,string,min=5"`
	LastName  string    `json:"last_name"  db:"last_name"  validate:"omitempty,string,min=5"`
	Email     string    `json:"email"      db:"email"      validate:"omitempty,email"`
	Phone     string    `json:"phone"      db:"phone"      validate:"omitempty,phone"`
	Username  string    `json:"username"   db:"username"   validate:"required,min=5"`
	Birthdate string    `json:"birthdate"  db:"birthdate"  validate:"omitempty,string,date"`
}

type UserList struct {
	Items      []User     `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type UserEditingFields struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Username  string
	Birthdate string
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Sort struct {
	Field  string `json:"field"`
	IsDesc bool   `json:"is_desc"`
}
