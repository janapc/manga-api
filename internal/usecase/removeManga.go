package usecase

import (
	"errors"
	"net/http"

	"github.com/janapc/manga-api/internal/infra/database"
	pkgId "github.com/janapc/manga-api/pkg/entity"
)

type RemoveManga struct {
	mangaDB database.MangaInterface
}

func NewRemoveManga(db database.MangaInterface) *RemoveManga {
	return &RemoveManga{
		mangaDB: db,
	}
}

func (r *RemoveManga) ExecuteRemoveManga(id string) *ErrorOutputDTO {
	var errorOutput ErrorOutputDTO
	_, err := pkgId.ParseID(id)
	if err != nil {
		errorOutput.StatusCode = http.StatusBadRequest
		errorOutput.Message = errors.New("id is invalid").Error()
		return &errorOutput
	}
	manga, err := r.mangaDB.FindMangaById(id)
	if err != nil {
		errorOutput.StatusCode = http.StatusNotFound
		errorOutput.Message = errors.New("manga is not found").Error()
		return &errorOutput
	}

	err = r.mangaDB.RemoveManga(manga)
	if err != nil {
		errorOutput.StatusCode = http.StatusInternalServerError
		errorOutput.Message = errors.New("internal server error").Error()
		return &errorOutput
	}

	return nil
}
