package main

import (
	"FeedAggregator/internal/database"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the env files")
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening a connection to database")
	}

	config := config{DB: database.New(db)}

	mainRouter := chi.NewRouter()
	addCors(mainRouter)

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", config.handlerPostUser)
	v1Router.Get("/users", config.handlerGetUserByApiKey)

	v1Router.Post("/feeds", config.middlewareAuth(config.handlerPostFeed))
	v1Router.Get("/feeds", config.handlerGetFeed)

	v1Router.Get("/posts", config.middlewareAuth(config.handlerGetPost))

	v1Router.Post("/feed_follows", config.middlewareAuth(config.handlerPostFeedFollow))
	v1Router.Get("/feed_follows", config.middlewareAuth(config.handlerGetFeedFollow))
	v1Router.Delete("/feed_follows/{FeedFollowId}", config.middlewareAuth(config.handlerDeleteFeedFollow))

	mainRouter.Mount("/v1", v1Router)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(config.DB, collectionConcurrency, collectionInterval)

	fmt.Println("Running the server on port: " + port)
	log.Fatal(srv.ListenAndServe())
}

func addCors(mainRouter *chi.Mux) {
	mainRouter.Use(cors.Handler(cors.Options{
		//AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
}
