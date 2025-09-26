package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/souther1407/blog/handlers"
	"github.com/souther1407/blog/interfaces"
	"github.com/souther1407/blog/internal/database"
	"github.com/souther1407/blog/middlewares"
)

var router *chi.Mux

var apiConfig interfaces.ApiConfig

func configRouter() {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Get("/health", handlers.HandlerHealth)
	router.Post("/users", middlewares.InjectDB(handlers.CreateUser, apiConfig))
	router.Post("/login", middlewares.InjectDB(handlers.LoginHandler, apiConfig))
	router.Post("/posts", middlewares.GetAuth(handlers.CreatePostHandler, apiConfig))
	router.Put("/posts/{post_id}", middlewares.GetAuth(handlers.UpdatePostHandler, apiConfig))
	router.Get("/posts/lasts", middlewares.InjectDB(handlers.GetLastPostsHandler, apiConfig))
	router.Delete("/posts/{post_id}", middlewares.GetAuth(handlers.DeletePostHandler, apiConfig))
}

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT env variable is required")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatalln("DB_URL env variable is required")
	}
	sessionKey := os.Getenv("SESSION_KEY")

	if sessionKey == "" {
		log.Fatalln("SESSION_KEY env variable is required")
	}

	router = chi.NewRouter()
	router.Use(middleware.Logger)

	log.Println("Connecting with database...")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln("DB connection failed", err)
	}
	sessionKeyDecoded, _ := hex.DecodeString(sessionKey)
	apiConfig = interfaces.ApiConfig{
		DB:          database.New(dbConn),
		CookieStore: sessions.NewCookieStore(sessionKeyDecoded),
	}

	log.Println("Connection with DB sucess")
	configRouter()
	log.Println("\033[92mListening at", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
