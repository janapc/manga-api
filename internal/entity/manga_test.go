package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewManga(t *testing.T) {
	manga, err := NewManga(NewMangaInput{Title: "Banana Banana", Description: "test test", Finished: false, InitialDate: "12/01/2001"})
	assert.Nil(t, err)
	assert.NotEmpty(t, manga.ID)
	assert.NotEmpty(t, manga.Description)
	assert.False(t, manga.Finished)
	assert.NotEmpty(t, manga.InitialDate)
	assert.Equal(t, manga.FinalDate, "")
}

func TestShouldValidateInput(t *testing.T) {
	input := NewMangaInput{Title: "", Description: "test"}
	assert.Error(t, ValidateInput(input))
	manga := NewMangaInput{Title: "Banana Banana", Description: "test test", Finished: false, InitialDate: "12/01/2001"}
	assert.Nil(t, ValidateInput(manga))
	manga.InitialDate = "12/13/1990"
	assert.Error(t, ValidateInput(manga))
	manga.InitialDate = "12/11/1990"
	manga.Finished = true
	assert.Error(t, ValidateInput(manga))
	manga.FinalDate = "19/10/223"
	assert.Error(t, ValidateInput(manga))
	manga.FinalDate = "19/10/2023"
	assert.Nil(t, ValidateInput(manga))
	manga.Finished = false
	assert.Error(t, ValidateInput(manga))
}

func TestShouldNotCreateANewMangaIfTitleIsEmpty(t *testing.T) {
	manga, err := NewManga(NewMangaInput{Description: "test test", Finished: false, InitialDate: "12/01/2001"})
	assert.NotNil(t, err)
	assert.Empty(t, manga)
}
