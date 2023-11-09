package database

import (
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldCreateAnUserInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := NewUserDatabase(db)
	user, err := entity.NewUser("test@test.com", "Test123")
	assert.Nil(t, err)
	err = userDb.CreateUser(user)
	assert.Nil(t, err)
	u, err := userDb.FindUserByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.NotEqual(t, u.Password, "Test123")
}

func TestShouldNotFindUserByEmailInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	userDb := NewUserDatabase(db)
	u, err := userDb.FindUserByEmail("test@test.com")
	assert.Nil(t, u)
	assert.Error(t, err)
}

func TestShouldFindUserByEmailInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	userDb := NewUserDatabase(db)
	user, err := entity.NewUser("test@test.com", "Test123")
	assert.Nil(t, err)
	err = userDb.CreateUser(user)
	assert.Nil(t, err)
	u, err := userDb.FindUserByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Email, user.Email)
	assert.Equal(t, u.Password, user.Password)
}
