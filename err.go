package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}