package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	w.WriteHeader(code)
	dat, _ := json.Marshal(payload)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJSON(w, code, struct{
		Error string `json:"error"`
	}{
		Error: msg,
	})
}