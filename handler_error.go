package main

import "net/http"

// must have this signature for http handler with go std library
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
