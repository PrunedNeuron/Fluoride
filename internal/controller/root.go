package controller

import (
	"fluoride/config"
	"net/http"

	"github.com/go-chi/render"
)

// GetIndex returns the index response
func GetIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fluoride, a robust icon pack management tool."))
	}
}

// GetVersion returns the current version of the application
func GetVersion() http.HandlerFunc {

	// Version struct to marshal into json
	type version struct {
		Version string `json:"version"`
	}
	var v = &version{Version: config.Version}

	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, v)
	}
}
