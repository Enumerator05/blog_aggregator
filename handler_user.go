package main

import (
    "net/http"
    "blot_aggregator/internal/database"
    "github.com/google/uuid"
    "time"
    "encoding/json"
    "log"
    "blot_aggregator/internal/auth"
)

func (cfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type params struct {
        Name  string `json:"name"`
    }

    decoder := json.NewDecoder(r.Body)
    param := params{}
    err := decoder.Decode(&param)

    if err != nil {
        log.Print(err)
        respondWithError(w, http.StatusInternalServerError, "Could not decode body")
        return
    }

    user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
        ID: uuid.New(),
        CreatedAt:  time.Now().UTC(),
        UpdatedAt:  time.Now().UTC(),
        Name:       param.Name,
    })

    if err != nil {
        log.Print(err)
        respondWithError(w, http.StatusInternalServerError, "Could not Create user")
        return
    }

    respondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {

    apiKey, err := auth.GetApiKey(r.Header)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Request")
        return
    }

    user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "user not found")
        return
    }

    respondWithJSON(w, http.StatusOK, user)
}
