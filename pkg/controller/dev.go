package controller

import (
	"fmt"
	"net/http"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetDevs renders all the devs in the database
func GetDevs(w http.ResponseWriter, r *http.Request) {
	devs, err := devService.GetDevs()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Devs:   devs,
	})
}

// GetDevCount renders the number of devs in the database
func GetDevCount(w http.ResponseWriter, r *http.Request) {

	count, err := devService.GetDevCount()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetDevByUsername renders the dev with the given username
func GetDevByUsername(w http.ResponseWriter, r *http.Request) {

	// Get dev from url
	username := chi.URLParam(r, "developer")

	if username == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	dev, err := devService.GetDevByUsername(username)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	var devs []model.User
	devs = append(devs, dev)

	render.JSON(w, r, &response{
		Status: "success",
		Devs:   devs,
	})
}
