package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
)

type GenerateTokenUser struct {
	UserDB       database.UserInterface
	JwtExpiresIn int
	Jwt          *jwtauth.JWTAuth
}

func NewGenerateTokenUser(db database.UserInterface, jwtExpiresIn int,
	jwt *jwtauth.JWTAuth) *GenerateTokenUser {
	return &GenerateTokenUser{UserDB: db, JwtExpiresIn: jwtExpiresIn, Jwt: jwt}
}

func (g *GenerateTokenUser) ExecuteGenerateTokenUser(user webserver.GetUserTokenInputDTO) (*GenerateTokenUserOutputDTO, *ErrorOutputDTO) {
	var errorOuput ErrorOutputDTO
	userFound, err := g.UserDB.FindUserByEmail(user.Email)
	if err != nil {
		errorOuput.StatusCode = http.StatusNotFound
		errorOuput.Message = errors.New("user not found").Error()
		return &GenerateTokenUserOutputDTO{}, &errorOuput
	}
	if !userFound.ValidatePassword(user.Password) {
		errorOuput.StatusCode = http.StatusUnauthorized
		errorOuput.Message = errors.New("user unauthorized").Error()
		return &GenerateTokenUserOutputDTO{}, &errorOuput
	}
	_, tokenString, _ := g.Jwt.Encode(map[string]interface{}{
		"sub": userFound.ID.String(),
		"exp": time.Now().Add(time.Minute * time.Duration(g.JwtExpiresIn)).Unix(),
	})
	accessToken := GenerateTokenUserOutputDTO{AccessToken: tokenString}
	return &accessToken, nil
}
