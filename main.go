package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

	// configure the server
	router := chi.NewRouter()
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
