package database

import (
	"github.com/janapc/manga-api/internal/entity"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserDatabase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{DB: db}
}

func (u *UserDatabase) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDatabase) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
