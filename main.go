package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env Loading failed")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port Not Found in ENV")
	}
	fmt.Printf("Port running on : %v", portString)

	//CHI router
	router := chi.NewRouter()

	//Setup Server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	srvErr := server.ListenAndServe()
	if srvErr != nil {
		log.Fatal("Something went wrong", srvErr)
	}
}
