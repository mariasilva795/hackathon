package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mariasilva795/go-api-rest/handlers"
	"github.com/mariasilva795/go-api-rest/middleware"
	"github.com/mariasilva795/go-api-rest/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file%v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	api.HandleFunc("/log", handlers.InsertEmotionalDailyLog(s)).Methods(http.MethodPost)

}
