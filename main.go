package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ManojGnanapalam/feedAggregator/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portNum := os.Getenv("PORT")

	if portNum == "" {
		log.Panic("Port is not found in the environment")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: false,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portNum,
	}
	log.Printf("server listen on Port %v", portNum)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
