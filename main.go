package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// load keyvalue pairs to env path
	godotenv.Load()

	// use the loaded env variable to get the port
	portString := os.Getenv("PORT")

	// kill the server/program if port wasn't found
	if portString == "" {
		log.Fatal("PORT not found in environment")
	}

	// configure the router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// configure the server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// start the server
	fmt.Println("server starting on port:", portString)
	err := server.ListenAndServe()

	// if the server spits out an error, kill the program/server and log it
	if err != nil {
		log.Fatal(err)
	}
}
