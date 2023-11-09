package entity

import (
	"errors"

	"github.com/janapc/manga-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(email, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password is mandatory")
	}
	hash, err := generatePassword(password)
	if err != nil {
		return nil, errors.New("password is wrong")
	}
	return &User{
		ID:       entity.NewID(),
		Email:    email,
		Password: string(hash),
	}, nil
}

func generatePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
