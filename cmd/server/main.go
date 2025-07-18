package main

import (
	"log"
	"net/http"

	"github.com/Higor-Edgar/booktime.git/config"
	"github.com/Higor-Edgar/booktime.git/internal/handler"
	"github.com/Higor-Edgar/booktime.git/internal/repository"
	"github.com/Higor-Edgar/booktime.git/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db := config.InitDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	authHandler := handler.NewAuthHandler(userService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", authHandler.SignUp)
	})

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
