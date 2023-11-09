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

func TestShouldRemoveManga(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	manga, err := entity.NewManga(entity.NewMangaInput{
		Title: "Test", Description: "Testing", InitialDate: "23/10/1990",
	})
	assert.Nil(t, err)
	mangaDB.SaveManga(manga)

	usecase := NewRemoveManga(mangaDB)
	errorOutput := usecase.ExecuteRemoveManga(manga.ID.String())
	assert.Nil(t, errorOutput)

	_, err = mangaDB.FindMangaById(manga.ID.String())
	assert.NotNil(t, err)
}

func TestShouldNotRemoveMangaIfIdIsInvalid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	manga, err := entity.NewManga(entity.NewMangaInput{
		Title: "Test", Description: "Testing", InitialDate: "23/10/1990",
	})
	assert.Nil(t, err)
	mangaDB.SaveManga(manga)

	usecase := NewRemoveManga(mangaDB)
	errorOutput := usecase.ExecuteRemoveManga("123")
	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusBadRequest)
}

func TestShouldNotRemoveMangaIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := database.NewMangaDatabase(db)

	manga, err := entity.NewManga(entity.NewMangaInput{
		Title: "Test", Description: "Testing", InitialDate: "23/10/1990",
	})
	assert.Nil(t, err)

	usecase := NewRemoveManga(mangaDB)
	errorOutput := usecase.ExecuteRemoveManga(manga.ID.String())
	assert.NotNil(t, errorOutput)
	assert.Equal(t, errorOutput.StatusCode, http.StatusNotFound)
}
