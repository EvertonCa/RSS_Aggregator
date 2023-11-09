package main

import "net/http"

// must have this signature for http handler with go std library
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
