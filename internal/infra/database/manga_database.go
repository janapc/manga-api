package database

import (
	"errors"

	"github.com/janapc/manga-api/internal/entity"
	"gorm.io/gorm"
)

type MangaDatabase struct {
	DB *gorm.DB
}

func NewMangaDatabase(db *gorm.DB) *MangaDatabase {
	return &MangaDatabase{DB: db}
}

func (m *MangaDatabase) SaveManga(manga *entity.Manga) error {
	return m.DB.Create(manga).Error
}

func (m *MangaDatabase) FindMangaByTitle(title string) (*entity.Manga, error) {
	var manga entity.Manga
	if err := m.DB.First(&manga, "title = ?", title).Error; err != nil {
		return nil, err
	}
	return &manga, nil
}

func (m *MangaDatabase) FindAllMangas(page, limit int, sort string) ([]entity.Manga, error) {
	var mangas []entity.Manga
	var err error
	if page != 0 && limit != 0 {
		err = m.DB.Limit(limit).Offset((page - 1) * limit).Order("initial_date " + sort).Find(&mangas).Error
	} else {
		err = m.DB.Order("initial_date " + sort).Find(&mangas).Error
	}
	return mangas, err
}

func (m *MangaDatabase) FindMangaById(id string) (*entity.Manga, error) {
	var manga entity.Manga
	if err := m.DB.First(&manga, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &manga, nil
}

func (m *MangaDatabase) UpdateManga(manga *entity.Manga) error {
	result := m.DB.Model(entity.Manga{}).Where("id = ?", manga.ID).Updates(manga)
	if result.RowsAffected == 0 {
		return errors.New("internal server error")
	}
	return nil
}

func (m *MangaDatabase) RemoveManga(manga *entity.Manga) error {
	return m.DB.Delete(manga).Error
}
