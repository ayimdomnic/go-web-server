package models

import (
	"html"
	"strings"
	"time"

	"github.com/ayimdomnic/go-web-server/helpers"
	"github.com/ayimdomnic/go-web-server/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint            `json:"id" gorm:"primary_key"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Email       string          `json:"email" gorm:"unique"`
	PhoneNumber string          `json:"phone_number" gorm:"unique"`
	Avatar      string          `json:"avatar"`
	UserRole    types.UserRoles `json:"user_role" sql:"type:user_role"`
	Password    string          `json:"password"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   time.Time       `json:"deleted_at"`
	Ticket      []Ticket
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}

func (u *User) SaveUser() (*User, error) {

	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	u := User{}
	err := DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := helpers.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
