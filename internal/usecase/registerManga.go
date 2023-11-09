package usecase

import (
	"errors"
	"net/http"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
)

type RegisterManga struct {
	mangaDB database.MangaInterface
}

func NewRegisterManga(db database.MangaInterface) *RegisterManga {
	return &RegisterManga{mangaDB: db}
}

func (r *RegisterManga) ExecuteRegisterManga(input webserver.RegisterMangaInputDTO) *ErrorOutputDTO {
	var errorOuput ErrorOutputDTO
	newManga := entity.NewMangaInput{
		Title:       input.Title,
		Description: input.Description,
		FinalDate:   input.FinalDate,
		Finished:    input.Finished,
		InitialDate: input.InitialDate,
	}
	manga, err := entity.NewManga(newManga)
	if err != nil {
		errorOuput.StatusCode = http.StatusBadRequest
		errorOuput.Message = err.Error()
		return &errorOuput
	}

	_, err = r.mangaDB.FindMangaByTitle(manga.Title)
	if err == nil {
		errorOuput.StatusCode = http.StatusConflict
		errorOuput.Message = errors.New("manga already exists").Error()
		return &errorOuput
	}

	err = r.mangaDB.SaveManga(manga)
	if err != nil {
		errorOuput.StatusCode = http.StatusInternalServerError
		errorOuput.Message = errors.New("internal server error").Error()
		return &errorOuput
	}
	return nil
}
