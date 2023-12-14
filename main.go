package main

import (
	"fmt"
	"log"
	"os"

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
}
