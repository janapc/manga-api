package usecase

type ErrorOutputDTO struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GenerateTokenUserOutputDTO struct {
	AccessToken string `json:"access_token"`
}
