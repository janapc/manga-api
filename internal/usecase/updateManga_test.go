package usecase

import (
	"net/http"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldUpdateManga(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})
	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "description", InitialDate: "21/12/2001"})
	assert.NoError(t, err)
	db.Create(manga)

	mangaDB := database.NewMangaDatabase(db)
	usecase := NewUpdateManga(mangaDB)
	manga.Title = "Testing"
	errorOutput := usecase.ExecuteUpdateManga(manga, manga.ID.String())
	assert.Nil(t, errorOutput)

	m, err := mangaDB.FindMangaById(manga.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, "Testing", m.Title)
	assert.Equal(t, "description", m.Description)
}

func TestShouldNotUpdateMangaIfIdIsInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)
	usecase := NewUpdateManga(mangaDB)
	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "description", InitialDate: "21/12/2001"})
	assert.NoError(t, err)
	errorOutput := usecase.ExecuteUpdateManga(manga, "asd")
	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusBadRequest)
}

func TestShouldNotUpdateMangaIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})
	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "description", InitialDate: "21/12/2001"})
	assert.NoError(t, err)

	mangaDB := database.NewMangaDatabase(db)
	usecase := NewUpdateManga(mangaDB)
	manga.Title = "Testing"
	errorOutput := usecase.ExecuteUpdateManga(manga, manga.ID.String())
	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusNotFound)
}
