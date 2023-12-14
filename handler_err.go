package main

import "net/http"

func handlerERR(w http.ResponseWriter, r *http.Request) {
	respondWithERR(w, 500, "Server Maintenance")
	// respondWithERR(w, 100, "client error")
}
