package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomError struct {
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func ApiHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// Handle the error
			log.Printf("Error: %+v\n", err)
			JSON(w, http.StatusBadRequest, CustomError{Error: err.Error()})
		}
	}
}

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
