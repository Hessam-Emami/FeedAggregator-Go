package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the env files")
	}

	port := os.Getenv("PORT")

	mainRouter := chi.NewRouter()
	addCors(mainRouter)

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handlerReadiness)
	v1Router.Get("/err", handlerError)

	mainRouter.Mount("/v1", v1Router)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	fmt.Println("Running the server on port: " + port)
	log.Fatal(srv.ListenAndServe())
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
