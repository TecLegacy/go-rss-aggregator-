package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// * Helper function for JSON response
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	//* Marshal payload
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error occurred during serialization", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

// * Helper Function for serialized consistent Error
func respondWithERR(w http.ResponseWriter, statusCode int, msg string) {

	//* Check error are server caused
	if statusCode > 499 {
		log.Printf("Error Caused by: %v", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	payload := errResponse{
		Error: msg,
	}

	respondWithJSON(w, statusCode, payload)
}
