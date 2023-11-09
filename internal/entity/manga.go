package entity

import (
	"errors"
	"time"

	"github.com/janapc/manga-api/pkg/entity"
)

type Manga struct {
	ID          entity.ID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Finished    bool      `json:"finished"`
	InitialDate string    `json:"initial_date"`
	FinalDate   string    `json:"final_date"`
}

type NewMangaInput struct {
	Title, Description string
	Finished           bool
	InitialDate        string
	FinalDate          string
}

func NewManga(input NewMangaInput) (*Manga, error) {
	err := ValidateInput(input)
	if err != nil {
		return nil, err
	}
	return &Manga{
		ID:          entity.NewID(),
		Title:       input.Title,
		Description: input.Description,
		Finished:    input.Finished,
		InitialDate: input.InitialDate,
		FinalDate:   input.FinalDate,
	}, nil
}

func ValidateInput(input NewMangaInput) error {
	if input.Description == "" || input.Title == "" || input.InitialDate == "" {
		return errors.New("the fields description, title and Initial_date is mandatory")
	}
	err := validateDate(input.InitialDate)
	if err != nil {
		return err
	}
	if input.Finished && input.FinalDate == "" {
		return errors.New("the fields finished and final_date is mandatory")
	}

	if input.FinalDate != "" && !input.Finished {
		return errors.New("the fields finished and final_date is mandatory")
	}
	if input.FinalDate != "" {
		err := validateDate(input.FinalDate)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateDate(date string) error {
	layout := "02/01/2006"
	_, err := time.Parse(layout, date)
	if err != nil {
		return errors.New("date invalid")
	}
	return nil
}
