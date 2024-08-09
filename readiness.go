package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, http.StatusOK, struct{
		Status string `json:"status"`
	}{
		Status: http.StatusText(http.StatusOK),
	})
}

func handlerErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}