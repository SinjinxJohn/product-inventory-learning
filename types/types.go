package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *RegisterUser) error
	GetUserByID(email int) (*User, error)
}
type User struct {
	ID        int       `json:"ID"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}
type RegisterUser struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// type ProfilePayload struct {
// 	FirstName string    `json:"firstname"`
// 	LastName  string    `json:"lastname"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

type Categories struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateCategoryPayload struct {
	Name string `json:"name" validate:"required"`
}

type CategoryStore interface {
	GetAllCategories() ([]*Categories, error)
	CreateCategory(*CreateCategoryPayload) error
	GetCategoryByName(categoryName string) (*Categories, error)
}
