package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eduvedras/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	params := struct {
		FeedId uuid.UUID `json:"feed_id"`
	}{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode the parameters")
		return
	}

	databaseFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create the feed follow")
		fmt.Println(err)
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(databaseFeedFollow))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := r.PathValue("feedFollowID")

	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handlerGetFeedFollowsOfUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, offset, err := getLimitAndOffset(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	databaseFeedFollows, err := cfg.DB.GetFeedFollowsOfUser(r.Context(), database.GetFeedFollowsOfUserParams{
		Limit:  int64(limit),
		Offset: int64(offset),
		UserID: user.ID})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(databaseFeedFollows))
}
