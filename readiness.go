package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	respondWithJSON(w, http.StatusOK, struct{
		Status string `json:"status"`
	}{
		Status: http.StatusText(http.StatusOK),
	})
}