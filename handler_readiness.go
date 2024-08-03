package main

import (
    "net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
    resonse := struct{
        Status  string `json:"status"`
    }{
        Status: "ok",
    }

    respondWithJSON(w, http.StatusOK, resonse)
}
