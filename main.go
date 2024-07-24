package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ManojGnanapalam/feedAggregator/handler"
	"github.com/ManojGnanapalam/feedAggregator/internal/database"
	"github.com/ManojGnanapalam/feedAggregator/internal/middleware"
	"github.com/ManojGnanapalam/feedAggregator/internal/mytype"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	portNum := os.Getenv("PORT")
	if portNum == "" {
		log.Panic("Port is not found in the environment")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Panic("Data Base URL is not found in the environment")
	}
	conn, error := sql.Open("postgres", dbURL)
	if error != nil {
		log.Fatal("Database not connected ", error)
	}

	apiCfg := &handler.LocalApiconfig{
		ApiConfig: mytype.ApiConfig{
			DB: database.New(conn),
		},
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
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", middleware.MiddlewareAuth((*middleware.LocalApiconfig)(apiCfg), apiCfg.HandlerGetUser))
	v1Router.Post("/rssfeed", middleware.MiddlewareAuth((*middleware.LocalApiconfig)(apiCfg), apiCfg.HandlerCreateFeed))
	v1Router.Get("/rssfeeds", apiCfg.HandlerGetFeeds)
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
