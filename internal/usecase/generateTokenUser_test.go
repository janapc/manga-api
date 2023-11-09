package usecase

import (
	"net/http"
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwt = jwtauth.New("HS256", []byte("test"), nil)

func TestShouldCreateANewAccessToken(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	user, _ := entity.NewUser("test@test.com", "test123")
	userDb.CreateUser(user)

	usecase := NewGenerateTokenUser(userDb, 3000, jwt)
	input := webserver.GetUserTokenInputDTO{Email: "test@test.com", Password: "test123"}
	result, outputError := usecase.ExecuteGenerateTokenUser(input)
	assert.Empty(t, outputError)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.AccessToken)
}

func TestShouldNotCreateANewAccessTokenIfUserNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	usecase := NewGenerateTokenUser(userDb, 3000, jwt)
	input := webserver.GetUserTokenInputDTO{Email: "test@test.com", Password: "test123"}
	result, outputError := usecase.ExecuteGenerateTokenUser(input)
	assert.Empty(t, result)
	assert.NotEmpty(t, outputError)
	assert.Equal(t, outputError.StatusCode, http.StatusNotFound)
}

func TestShouldNotCreateANewAccessTokenIfPasswordIsInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})

	userDb := database.NewUserDatabase(db)
	user, _ := entity.NewUser("test@test.com", "test123")
	userDb.CreateUser(user)

	usecase := NewGenerateTokenUser(userDb, 3000, jwt)
	input := webserver.GetUserTokenInputDTO{Email: "test@test.com", Password: "test1234"}
	result, outputError := usecase.ExecuteGenerateTokenUser(input)
	assert.Empty(t, result)
	assert.NotEmpty(t, outputError)
	assert.Equal(t, outputError.StatusCode, http.StatusUnauthorized)
}
