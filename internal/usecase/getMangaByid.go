package usecase

import (
	"errors"
	"net/http"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	pkgId "github.com/janapc/manga-api/pkg/entity"
)

type GetMangaById struct {
	mangaDB database.MangaInterface
}

func NewGetMangaById(db database.MangaInterface) *GetMangaById {
	return &GetMangaById{
		mangaDB: db,
	}
}

func (g *GetMangaById) ExecuteGetMangaById(id string) (*entity.Manga, *ErrorOutputDTO) {
	var errorOutput ErrorOutputDTO
	if _, err := pkgId.ParseID(id); err != nil {
		errorOutput.StatusCode = http.StatusBadRequest
		errorOutput.Message = errors.New("id is invalid").Error()
		return nil, &errorOutput
	}
	manga, err := g.mangaDB.FindMangaById(id)
	if err != nil {
		errorOutput.StatusCode = http.StatusNotFound
		errorOutput.Message = errors.New("manga is not found").Error()
		return nil, &errorOutput
	}
	return manga, nil
}
