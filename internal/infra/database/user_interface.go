package database

import "github.com/janapc/manga-api/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
