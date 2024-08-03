package main

import (
    "net/http"
    "blot_aggregator/internal/database"
    "github.com/google/uuid"
    "encoding/json"
    "time"
)

func (cfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
        FeedId     uuid.UUID `json:"feed_id"`
    }

    params := parameters{}
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "could not decode parameters")
        return
    }

    _, err = cfg.DB.GetFeedById(r.Context(), params.FeedId)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "could not find feed")
        return
    }

    feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
        ID:         uuid.New(),
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
        FeedID:     params.FeedId,
        UserID:     user.ID,
    })

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not create feed follow")
        return
    }

    respondWithJSON(w, http.StatusOK, feedFollow)
}

func (cfg *apiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollowIdStr := r.PathValue("feed_follow_id")

    feedFollowId, err := uuid.Parse(feedFollowIdStr)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not parse uuid")
        return
    }

    feedFollow, err := cfg.DB.DeleteFeedFollowById(r.Context(), database.DeleteFeedFollowByIdParams{
        ID:     feedFollowId,
        UserID: user.ID,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not delete feed follow")
        return
    }

    respondWithJSON(w, http.StatusOK, feedFollow)
}

func (cfg *apiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollows, err := cfg.DB.GetAllFeedFollowsByUserId(r.Context(), user.ID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not get feed follows")
        return
    }

    respondWithJSON(w, http.StatusOK, feedFollows)
}
