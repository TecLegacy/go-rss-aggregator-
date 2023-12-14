package main

import "net/http"

// * Sending Response to client
func handlerHome(w http.ResponseWriter, r *http.Request) {

	// payload := map[string]string{
	// 	"message": "hie",
	// }

	type Response struct {
		Message string `json:"message"`
	}

	response := Response{
		Message: "hie",
	}
	respondWithJSON(w, http.StatusOK, response)
	// respondWithJSON(w, http.StatusOK, struct{}{})
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	type serverStatus struct {
		UpStatus bool `json:"server-status"`
	}

	respondWithJSON(w, 200, serverStatus{
		UpStatus: true,
	})
}
