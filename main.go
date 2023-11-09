package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// load the local env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// get the PORT value from env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in the environment")
	}

	// creates a new router
	router := chi.NewRouter()

	// server configurations
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	// connect router to http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// starts the server
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
