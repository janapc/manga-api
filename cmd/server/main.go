package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/janapc/manga-api/configs"

	_ "github.com/janapc/manga-api/docs"
	"github.com/janapc/manga-api/internal/entity"

	"github.com/janapc/manga-api/internal/infra/database"
	"github.com/janapc/manga-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Manga API
// @version 1.0
// @description Manager mangas

// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dsn := configs.DBUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Manga{})

	userDB := database.NewUserDatabase(db)
	userHandler := handlers.NewUserHandler(userDB)

	mangaDB := database.NewMangaDatabase(db)
	mangaHandler := handlers.NewMangaHandler(mangaDB)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.JwtAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))
	baseUrl := fmt.Sprintf("%s/swagger/doc.json", configs.BaseUrlV1)
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userRouter(userHandler))
		r.Mount("/mangas", mangaRouter(mangaHandler, configs.JwtAuth))
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(baseUrl)))
	})
	http.ListenAndServe(":3000", r)
}

func userRouter(userHandler *handlers.UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/", userHandler.CreateUser)
	r.Post("/generate_token", userHandler.GetUserToken)
	return r
}

func mangaRouter(mangaHandler *handlers.MangaHandler, jwtAuth *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(jwtAuth))
	r.Use(jwtauth.Authenticator)
	r.Post("/", mangaHandler.RegisterManga)
	r.Get("/{id}", mangaHandler.GetMangaById)
	r.Get("/", mangaHandler.GetMangas)
	r.Patch("/{id}", mangaHandler.UpdateManga)
	r.Delete("/{id}", mangaHandler.RemoveManga)
	return r
}
