package usecase

import (
	"errors"
	"net/http"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	pkgId "github.com/janapc/manga-api/pkg/entity"
)

type UpdateManga struct {
	mangaDB database.MangaInterface
}

func NewUpdateManga(db database.MangaInterface) *UpdateManga {
	return &UpdateManga{mangaDB: db}
}

func (u *UpdateManga) ExecuteUpdateManga(input *entity.Manga, id string) *ErrorOutputDTO {
	var errorOutput ErrorOutputDTO
	idValid, err := pkgId.ParseID(id)
	if err != nil {
		errorOutput.StatusCode = http.StatusBadRequest
		errorOutput.Message = errors.New("id is invalid").Error()
		return &errorOutput
	}
	input.ID = idValid
	_, err = u.mangaDB.FindMangaById(input.ID.String())
	if err != nil {
		errorOutput.StatusCode = http.StatusNotFound
		errorOutput.Message = errors.New("manga is not found").Error()
		return &errorOutput
	}
	err = u.mangaDB.UpdateManga(input)
	if err != nil {
		errorOutput.StatusCode = http.StatusInternalServerError
		errorOutput.Message = errors.New("internal server error").Error()
		return &errorOutput
	}
	return nil
}
