package models

import "time"

import "github.com/ayimdomnic/go-web-server/types"

type User struct {
	ID          uint            `json:"id" gorm:"primary_key"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Avatar      string          `json:"avatar"`
	UserRole    types.UserRoles `json:"user_role" sql:"type:user_role"`
	Password    string          `json:"password"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   time.Time       `json:"deleted_at"`
	Ticket []Ticket
}
