package usecase

import (
	"net/http"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldRegisterANewUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	usecase := NewRegisterUser(userDb)
	input := webserver.CreateUserInputDTO{Email: "test@test.com", Password: "test123"}
	outputError := usecase.ExecuteRegisterUser(input)
	assert.Empty(t, outputError)

	u, err := userDb.FindUserByEmail(input.Email)
	assert.Nil(t, err)
	assert.NotEmpty(t, u)
	assert.NotEqual(t, input.Password, u.Password)
	assert.Equal(t, input.Email, u.Email)
}

func TestShouldNotRegisterANewUserIfUserExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	user, err := entity.NewUser("test@test.com", "test123")
	assert.Nil(t, err)
	userDb.CreateUser(user)

	usecase := NewRegisterUser(userDb)
	input := webserver.CreateUserInputDTO{Email: "test@test.com", Password: "test12s3"}
	outputError := usecase.ExecuteRegisterUser(input)
	assert.Equal(t, outputError.StatusCode, http.StatusConflict)
	assert.Equal(t, outputError.Message, "user already exists")
}

func TestShouldNotRegisterANewUserIfEmailIsEmpty(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	usecase := NewRegisterUser(userDb)
	input := webserver.CreateUserInputDTO{Email: "", Password: "test123"}
	outputError := usecase.ExecuteRegisterUser(input)
	assert.NotEmpty(t, outputError)
	assert.Equal(t, outputError.StatusCode, http.StatusBadRequest)
}
