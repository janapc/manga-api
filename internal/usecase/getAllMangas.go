package usecase

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
)

type GetAllMangas struct {
	mangaDB database.MangaInterface
}

func NewGetAllMangas(db database.MangaInterface) *GetAllMangas {
	return &GetAllMangas{
		mangaDB: db,
	}
}

func (g *GetAllMangas) ExecuteGetAllMangas(limit, page, sort string) ([]entity.Manga, *ErrorOutputDTO) {
	var errorOutput ErrorOutputDTO
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	mangas, err := g.mangaDB.FindAllMangas(pageInt, limitInt, sort)
	if err != nil {
		errorOutput.StatusCode = http.StatusInternalServerError
		errorOutput.Message = errors.New("internal server error").Error()
		return nil, &errorOutput
	}

	return mangas, nil
}
