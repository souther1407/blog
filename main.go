package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/souther1407/blog/handlers"
)

var router *chi.Mux

func configRouter() {
	router.Get("/health", handlers.HandlerHealth)
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT env variable is required")
	}
	router = chi.NewRouter()
	router.Use(middleware.Logger)
	configRouter()
	log.Println("\033[92mListening at", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
