package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eduvedras/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	params := struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode the parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	databaseFeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't get users")
		return
	}

	respondWithJSON(w, http.StatusOK,databaseFeedsToFeeds(databaseFeeds))
}