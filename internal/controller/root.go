package controller

import (
	"fluoride/config"
	"net/http"

	"github.com/go-chi/render"
)

// GetIndex returns the index response
func GetIndex(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "Fluoride, a robust icon pack management tool.")
}

// GetVersion returns the current version of the application
func GetVersion(w http.ResponseWriter, r *http.Request) {

	// Version struct to marshal into json
	type version struct {
		Version string `json:"version"`
	}
	var v = &version{Version: config.Version}

	render.JSON(w, r, v)
}

// NotFound handles the case where the url has no mapping
func NotFound(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "The page you're looking for does not exist. 404 error.")
}
