package webserver

type CreateUserInputDTO struct {
	Email    string
	Password string
}

type GetUserTokenInputDTO struct {
	Email    string
	Password string
}

type UserErrorOutputDTO struct {
	Message string `json:"message"`
}

type RegisterMangaInputDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
}

type MangaErrorOutputDTO struct {
	Message string `json:"message"`
}

type UpdateMangaInputDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
}
