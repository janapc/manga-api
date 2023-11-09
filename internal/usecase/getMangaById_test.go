package usecase

import (
	"net/http"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	pkgId "github.com/janapc/manga-api/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldGetMangaById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	manga, err := entity.NewManga(entity.NewMangaInput{
		Title: "Test", Description: "Testing", InitialDate: "23/10/1990",
	})
	assert.Nil(t, err)
	mangaDB.SaveManga(manga)

	usecase := NewGetMangaById(mangaDB)
	result, errorOutput := usecase.ExecuteGetMangaById(manga.ID.String())

	assert.Nil(t, errorOutput)
	assert.NotEmpty(t, result)
	assert.Equal(t, result.Description, manga.Description)
	assert.Equal(t, result.FinalDate, manga.FinalDate)
	assert.Equal(t, result.Title, manga.Title)
	assert.Equal(t, result.InitialDate, manga.InitialDate)
	assert.Equal(t, result.ID, manga.ID)
	assert.Equal(t, result.Finished, manga.Finished)
}

func TestShouldNotGetMangaByIdIfIdIsInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	usecase := NewGetMangaById(mangaDB)
	result, errorOutput := usecase.ExecuteGetMangaById("123")

	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusBadRequest)
	assert.Empty(t, result)
}

func TestShouldNotGetMangaByIdIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	usecase := NewGetMangaById(mangaDB)
	result, errorOutput := usecase.ExecuteGetMangaById(pkgId.NewID().String())

	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusNotFound)
	assert.Empty(t, result)
}
