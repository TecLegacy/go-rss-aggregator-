package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	//* load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env Loading failed")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port Not Found in ENV")
	}
	fmt.Printf("Port running on : %v", portString)

	//*CHI router
	router := chi.NewRouter()

	//*CORS middleware policy
	router.Use(cors.Handler(cors.Options{

		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, //? Maximum value not ignored by any of major browsers for preflight checkup
	}))

	// * HookUp Home handler
	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerHealthz)

	v1Router.Get("/home", handlerHome)
	v1Router.Get("/err", handlerERR)

	// * Mount for API versioning
	router.Mount("/v1", v1Router)

	//*Setup Server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	srvErr := server.ListenAndServe()
	if srvErr != nil {
		log.Fatal("Something went wrong", srvErr)
	}
}
