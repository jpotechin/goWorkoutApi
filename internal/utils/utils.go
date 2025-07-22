package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Envelope is a shorthand for a generic JSON object response.
// It allows sending JSON like {"key": value} in a structured way.
type Envelope map[string]any

// WriteJSON marshals the data into indented JSON and writes it to the HTTP response
// with the provided status code and proper headers.
func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	// Marshal the data into human-readable (indented) JSON
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		// Return the marshaling error if it fails
		return err
	}

	// Add a newline for better output formatting in terminals/logs
	js = append(js, '\n')

	// Set content type so the client interprets it correctly
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code
	w.WriteHeader(status)

	// Write the JSON response body
	w.Write(js)

	return nil
}

func ReadIdParam(r *http.Request) (int64, error) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		return 0, errors.New("Invalid Id parameter")
	}

	id, err := strconv.ParseInt(paramsWorkoutID, 10, 64)

	if err != nil {

		return 0, errors.New("Invalid Id parameter type")
	}

	return id, nil
}
