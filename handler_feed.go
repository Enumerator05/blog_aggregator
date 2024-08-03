package main

import (
    "net/http"
    "blot_aggregator/internal/database"
    "encoding/json"
    "github.com/google/uuid"
    "time"
    "log"
)

func (cfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
        Name    string  `json:"name"`
        Url     string  `json:"url"`
    }

    type response struct {
        FeedFollow  database.FeedFollow `json:"feed_follow"`
        Feed        database.Feed       `json:"feed"`
    }

    params := parameters{}
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not decode parameters")
        return
    }

    feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
        ID:         uuid.New(),
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
        Url:        params.Url,
        UserID:     user.ID,
        Name:       params.Name,
    })

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not create feed")
        return
    }

    feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
        ID:         uuid.New(),
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
        FeedID:     feed.ID,
        UserID:     user.ID,
    })
    if err != nil {
        log.Print(err)
        respondWithError(w, http.StatusInternalServerError, "could not create feed follow")
        return
    }


    respondWithJSON(w, http.StatusOK, response{
        Feed:       feed,
        FeedFollow: feedFollow,
    })
}

func (cfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
    feeds, err := cfg.DB.GetAllFeeds(r.Context())
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not get feeds")
        return
    }

    respondWithJSON(w, http.StatusOK, feeds)
}
