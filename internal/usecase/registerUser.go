package usecase

import (
	"errors"
	"net/http"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
)

type RegisterUser struct {
	UserDB database.UserInterface
}

func NewRegisterUser(db database.UserInterface) *RegisterUser {
	return &RegisterUser{UserDB: db}
}

func (c *RegisterUser) ExecuteRegisterUser(input webserver.CreateUserInputDTO) *ErrorOutputDTO {
	var errorOutput ErrorOutputDTO
	user, err := entity.NewUser(input.Email, input.Password)
	if err != nil {
		errorOutput.StatusCode = http.StatusBadRequest
		errorOutput.Message = err.Error()
		return &errorOutput
	}
	u, _ := c.UserDB.FindUserByEmail(user.Email)
	if u != nil {
		errorOutput.StatusCode = http.StatusConflict
		errorOutput.Message = errors.New("user already exists").Error()
		return &errorOutput
	}
	err = c.UserDB.CreateUser(user)
	if err != nil {
		errorOutput.StatusCode = http.StatusInternalServerError
		errorOutput.Message = errors.New("internal server error").Error()
		return &errorOutput
	}
	return nil
}
