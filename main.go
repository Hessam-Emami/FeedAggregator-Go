package main

import (
	"FeedAggregator/internal/database"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)
import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	mainRouter.Mount("/v1", v1Router)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	fmt.Println("Running the server on port: " + port)
	log.Fatal(srv.ListenAndServe())
}

func (c config) handlerPostUser(writer http.ResponseWriter, request *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}
	type responseBody struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
	}

	body := requestBody{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Couldn't decode parameters")
		fmt.Println("Error decoding parameter: " + err.Error())
		return
	}
	if len(body.Name) == 0 {
		respondWithError(writer, http.StatusBadRequest, "Invalid request body")
		return
	}

	uUID, err := uuid.NewUUID()
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating uuid: " + err.Error())
		return
	}

	usr, err := c.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      body.Name,
	})
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, "Internal server error")
		fmt.Println("Error creating user: " + err.Error())
		return
	}

	rBody := responseBody{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
	}
	respondWithJSON(writer, http.StatusOK, rBody)
}

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, 200, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func handlerError(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, 500, "Interval Server Error")
}

func addCors(mainRouter *chi.Mux) {
	mainRouter.Use(cors.Handler(cors.Options{
		//AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
}
