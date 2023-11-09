package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
	"github.com/janapc/manga-api/internal/usecase"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @tags users
// @Accept json
// @Produce json
// @Param request body webserver.CreateUserInputDTO true "user request"
// @Success 201
// @Failure 400 {object} webserver.UserErrorOutputDTO
// @Failure 404 {object} webserver.UserErrorOutputDTO
// @Failure 500 {object} webserver.UserErrorOutputDTO
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user webserver.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webserver.UserErrorOutputDTO{Message: "data invalid"})
		return
	}
	registerUser := usecase.NewRegisterUser(h.UserDB)
	errorOutput := registerUser.ExecuteRegisterUser(user)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.UserErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetUserToken godoc
// @Summary get a token
// @Description get a token
// @tags users
// @Accept json
// @Produce json
// @Param request body webserver.GetUserTokenInputDTO true "user request"
// @Success 200 {object} usecase.GenerateTokenUserOutputDTO
// @Failure 400 {object} webserver.UserErrorOutputDTO
// @Failure 401 {object} webserver.UserErrorOutputDTO
// @Failure 500 {object} webserver.UserErrorOutputDTO
// @Router /users/generate_token [post]
func (h *UserHandler) GetUserToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user webserver.GetUserTokenInputDTO
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webserver.UserErrorOutputDTO{Message: "data invalid"})
		return
	}
	generateTokenUser := usecase.NewGenerateTokenUser(h.UserDB, jwtExpiresIn, jwt)
	token, errorOutput := generateTokenUser.ExecuteGenerateTokenUser(user)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.UserErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
