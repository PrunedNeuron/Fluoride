package server

import (
	"icon-requests/config"
	"net/http"

	"github.com/go-chi/render"
)

// GetVersion returns version
func GetVersion() http.HandlerFunc {

	// Simple version struct
	type version struct {
		Version string `json:"version"`
	}
	var v = &version{Version: config.GitVersion}

	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, v)
	}
}
