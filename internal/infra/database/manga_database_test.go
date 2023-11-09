package database

import (
	"fmt"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldCreateAnMangaInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDb := NewMangaDatabase(db)
	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	err = mangaDb.SaveManga(manga)
	assert.Nil(t, err)
	m, err := mangaDb.FindMangaByTitle(manga.Title)
	assert.Nil(t, err)
	assert.Equal(t, m.ID, manga.ID)
}

func TestShouldFindMangaByTitleInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	db.Create(manga)

	mangaDb := NewMangaDatabase(db)
	result, err := mangaDb.FindMangaByTitle("test")
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result.ID, manga.ID)
	assert.Equal(t, result.Title, manga.Title)
	assert.Equal(t, result.Description, manga.Description)
	assert.Equal(t, result.InitialDate, manga.InitialDate)
	assert.Equal(t, result.FinalDate, manga.FinalDate)
	assert.Equal(t, result.Finished, manga.Finished)
}

func TestShouldNotFindMangaByTitleInDBIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDb := NewMangaDatabase(db)
	result, err := mangaDb.FindMangaByTitle("test")
	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestShouldFindMangaByIdInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	db.Create(manga)

	mangaDb := NewMangaDatabase(db)
	result, err := mangaDb.FindMangaById(manga.ID.String())
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result.ID, manga.ID)
	assert.Equal(t, result.Title, manga.Title)
	assert.Equal(t, result.Description, manga.Description)
	assert.Equal(t, result.InitialDate, manga.InitialDate)
	assert.Equal(t, result.FinalDate, manga.FinalDate)
	assert.Equal(t, result.Finished, manga.Finished)
}

func TestShouldNotFindMangaByIdInDBIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDb := NewMangaDatabase(db)
	result, err := mangaDb.FindMangaById("asd-123")
	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestShouldUpdateMangaInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	db.Create(manga)

	mangaDb := NewMangaDatabase(db)
	manga.Title = "Testing update"
	err = mangaDb.UpdateManga(manga)
	assert.Nil(t, err)

	result, err := mangaDb.FindMangaById(manga.ID.String())
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, result.ID, manga.ID)
	assert.Equal(t, result.Title, "Testing update")
	assert.Equal(t, result.Description, manga.Description)
	assert.Equal(t, result.InitialDate, manga.InitialDate)
	assert.Equal(t, result.FinalDate, manga.FinalDate)
	assert.Equal(t, result.Finished, manga.Finished)
}

func TestShouldNotUpdateMangaInDBIfMangaDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDb := NewMangaDatabase(db)
	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	err = mangaDb.UpdateManga(manga)
	assert.NotNil(t, err)
}

func TestShouldRemoveMangaInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	manga, err := entity.NewManga(entity.NewMangaInput{Title: "test", Description: "testing", InitialDate: "21/12/2009"})
	assert.Nil(t, err)
	db.Create(manga)

	mangaDb := NewMangaDatabase(db)
	err = mangaDb.RemoveManga(manga)
	assert.Nil(t, err)

	result, err := mangaDb.FindMangaById(manga.ID.String())
	assert.NotEmpty(t, err)
	assert.Nil(t, result)
}

func TestShouldFindAllMangasInDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	for i := 1; i <= 10; i++ {
		manga, err := entity.NewManga(entity.NewMangaInput{Title: fmt.Sprintf("Test %d", i), Description: fmt.Sprintf("Description  %d", i), InitialDate: "21/12/2009"})
		assert.NoError(t, err)
		db.Create(manga)
	}

	mangaDB := NewMangaDatabase(db)
	mangas, err := mangaDB.FindAllMangas(1, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, mangas, 5)
	assert.Equal(t, "Test 1", mangas[0].Title)
	assert.Equal(t, "Test 5", mangas[4].Title)

	mangas, err = mangaDB.FindAllMangas(2, 5, "asc")
	assert.NoError(t, err)
	assert.Len(t, mangas, 5)
	assert.Equal(t, "Test 6", mangas[0].Title)
	assert.Equal(t, "Test 10", mangas[4].Title)
}

func TestShouldNotFindAllMangasInDBIfDoesNotExistsResgiter(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})

	mangaDB := NewMangaDatabase(db)
	mangas, err := mangaDB.FindAllMangas(1, 5, "asc")
	assert.Nil(t, err)
	assert.Len(t, mangas, 0)
}
