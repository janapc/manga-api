package database

import "github.com/janapc/manga-api/internal/entity"

type MangaInterface interface {
	SaveManga(manga *entity.Manga) error
	FindMangaByTitle(title string) (*entity.Manga, error)
	FindAllMangas(page, limit int, sort string) ([]entity.Manga, error)
	FindMangaById(id string) (*entity.Manga, error)
	UpdateManga(manga *entity.Manga) error
	RemoveManga(manga *entity.Manga) error
}
