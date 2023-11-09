package usecase

import (
	"fmt"
	"testing"

	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldGetAllMangas(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Manga{})
	for i := 1; i <= 5; i++ {
		manga, err := entity.NewManga(entity.NewMangaInput{Title: fmt.Sprintf("Test %d", i), Description: fmt.Sprintf("Description  %d", i), InitialDate: fmt.Sprintf("21/12/200%d", i)})
		assert.NoError(t, err)
		db.Create(manga)
	}

	mangaDB := database.NewMangaDatabase(db)
	usecase := NewGetAllMangas(mangaDB)
	mangas, errorOutput := usecase.ExecuteGetAllMangas("", "", "")
	assert.Nil(t, errorOutput)
	assert.Len(t, mangas, 5)

	mangas, errorOutput = usecase.ExecuteGetAllMangas("3", "1", "ac")
	assert.Nil(t, errorOutput)
	assert.Len(t, mangas, 3)
	assert.Equal(t, mangas[0].Title, "Test 1")

	mangas, errorOutput = usecase.ExecuteGetAllMangas("2", "1", "desc")
	assert.Nil(t, errorOutput)
	assert.Len(t, mangas, 2)
	assert.Equal(t, mangas[0].Title, "Test 5")

}
