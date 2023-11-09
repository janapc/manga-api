package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/janapc/manga-api/internal/entity"
	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver"
	"github.com/janapc/manga-api/internal/usecase"
)

type MangaHandler struct {
	MangaDB database.MangaInterface
}

func NewMangaHandler(mangaDB database.MangaInterface) *MangaHandler {
	return &MangaHandler{MangaDB: mangaDB}
}

// RegisterManga godoc
// @Summary register a new manga
// @Description register a new manga
// @tags mangas
// @Accept json
// @Produce json
// @Param request body webserver.RegisterMangaInputDTO true "manga request"
// @Success 201
// @Failure 400 {object} webserver.MangaErrorOutputDTO
// @Failure 404 {object} webserver.MangaErrorOutputDTO
// @Failure 500 {object} webserver.MangaErrorOutputDTO
// @Router /mangas [post]
// @Security BearerAuth
func (m *MangaHandler) RegisterManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var manga webserver.RegisterMangaInputDTO
	err := json.NewDecoder(r.Body).Decode(&manga)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: "data invalid"})
		return
	}
	registerManga := usecase.NewRegisterManga(m.MangaDB)
	errorOutput := registerManga.ExecuteRegisterManga(manga)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetMangaById godoc
// @Summary get a manga by id
// @Description get a manga by id
// @tags mangas
// @Accept json
// @Produce json
// @Param id path string true "manga id" Format(uuid)
// @Success 200 {object} entity.Manga
// @Failure 400 {object} webserver.MangaErrorOutputDTO
// @Failure 404 {object} webserver.MangaErrorOutputDTO
// @Failure 500 {object} webserver.MangaErrorOutputDTO
// @Router /mangas/{id} [get]
// @Security BearerAuth
func (m *MangaHandler) GetMangaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: "the field id is mandatory"})
		return
	}
	getMangaById := usecase.NewGetMangaById(m.MangaDB)
	manga, errorOutput := getMangaById.ExecuteGetMangaById(id)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(manga)
}

// GetMangas godoc
// @Summary get all mangas
// @Description get all mangas
// @tags mangas
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "limit"
// @Param sort query string false "sort"
// @Success 200 {array} entity.Manga
// @Failure 404 {object} webserver.MangaErrorOutputDTO
// @Failure 500 {object} webserver.MangaErrorOutputDTO
// @Router /mangas [get]
// @Security BearerAuth
func (m *MangaHandler) GetMangas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")
	getAllMangas := usecase.NewGetAllMangas(m.MangaDB)
	mangas, errorOutput := getAllMangas.ExecuteGetAllMangas(limit, page, sort)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mangas)
}

// UpdateManga godoc
// @Summary update a manga
// @Description update a manga
// @tags mangas
// @Accept json
// @Produce json
// @Param request body entity.Manga true "manga request"
// @Param id path string true "manga id" Format(uuid)
// @Success 204
// @Failure 400 {object} webserver.MangaErrorOutputDTO
// @Failure 404 {object} webserver.MangaErrorOutputDTO
// @Failure 500 {object} webserver.MangaErrorOutputDTO
// @Router /mangas/{id} [patch]
// @Security BearerAuth
func (m *MangaHandler) UpdateManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var manga entity.Manga
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&manga)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: "data invalid"})
		return
	}
	updateManga := usecase.NewUpdateManga(m.MangaDB)
	errorOutput := updateManga.ExecuteUpdateManga(&manga, id)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// RemoveManga godoc
// @Summary remove a manga
// @Description remove a manga
// @tags mangas
// @Accept json
// @Produce json
// @Param id path string true "manga id" Format(uuid)
// @Success 204
// @Failure 400 {object} webserver.MangaErrorOutputDTO
// @Failure 404 {object} webserver.MangaErrorOutputDTO
// @Failure 500 {object} webserver.MangaErrorOutputDTO
// @Router /mangas/{id} [delete]
// @Security BearerAuth
func (m *MangaHandler) RemoveManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	removeManga := usecase.NewRemoveManga(m.MangaDB)
	errorOutput := removeManga.ExecuteRemoveManga(id)
	if errorOutput != nil {
		w.WriteHeader(errorOutput.StatusCode)
		json.NewEncoder(w).Encode(webserver.MangaErrorOutputDTO{Message: errorOutput.Message})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
