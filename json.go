package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// errors on the server side
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON response: %v", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err2 := w.Write(dat)
	if err2 != nil {
		log.Printf("Failed to write JSON: %v", err2)
	}
}
