package usecase

import (
	"net/http"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldRegisterANewManga(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDb := database.NewMangaDatabase(db)
	usecase := NewRegisterManga(mangaDb)
	input := webserver.RegisterMangaInputDTO{
		Title:       "banana",
		Description: "test test",
		InitialDate: "02/10/2003",
	}
	outputError := usecase.ExecuteRegisterManga(input)
	assert.Empty(t, outputError)

	m, err := mangaDb.FindMangaByTitle(input.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, m)
	assert.NotEmpty(t, m.ID.String())
	assert.False(t, m.Finished)
	assert.Empty(t, m.FinalDate)
	assert.Equal(t, input.Description, m.Description)
}

func TestShouldNotRegisterMangaIfMangaAlreadyExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})
	m := entity.NewMangaInput{
		Title:       "banana",
		Description: "test test",
		InitialDate: "02/10/2003",
	}
	manga, err := entity.NewManga(m)
	assert.Nil(t, err)
	db.Create(&manga)
	mangaDb := database.NewMangaDatabase(db)
	usecase := NewRegisterManga(mangaDb)
	input := webserver.RegisterMangaInputDTO{
		Title:       "banana",
		Description: "test 2",
		InitialDate: "12/10/2003",
	}
	outputError := usecase.ExecuteRegisterManga(input)
	assert.NotEmpty(t, outputError)
	assert.Equal(t, outputError.StatusCode, http.StatusConflict)
}

func TestShouldNotRegisterMangaIfTitleIsEmpty(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})
	mangaDb := database.NewMangaDatabase(db)
	usecase := NewRegisterManga(mangaDb)
	input := webserver.RegisterMangaInputDTO{
		Title:       "",
		Description: "test 2",
		InitialDate: "12/10/2003",
	}
	outputError := usecase.ExecuteRegisterManga(input)
	assert.NotEmpty(t, outputError)
	assert.Equal(t, outputError.StatusCode, http.StatusBadRequest)
}
