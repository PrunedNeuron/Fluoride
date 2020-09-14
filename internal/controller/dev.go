package controller

import (
	"fluoride/pkg/errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetPacksByDev creates a new user
func GetPacksByDev(w http.ResponseWriter, r *http.Request) {

	// Get dev from url
	dev := chi.URLParam(r, "dev")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	packs, err := devService.GetPacksByDev(dev)

	if err == errors.ErrDatabaseNotFound {
		render.Render(w, r, errors.ErrNotFound)
	} else if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
	} else {
		render.JSON(w, r, packs)
	}
}
