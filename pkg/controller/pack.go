package controller

import (
	"fmt"
	"net/http"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CreatePack creates a new icon pack
func CreatePack(w http.ResponseWriter, r *http.Request) {

	dev := chi.URLParam(r, "developer")
	var pack model.Pack
	err := render.DecodeJSON(r.Body, &pack)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	if pack.Name == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		return
	}

	if pack.DevUsername == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing developer username value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing developer username value")))
		return
	}

	// URL Dev != request body dev
	if pack.DevUsername != dev {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Developer username mismatch")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Developer username mismatch")))
		return
	}

	if pack.URL == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing URL value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing URL value")))
		return
	}

	if pack.BillingStatus == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing billing status value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing billing status value")))
		return
	}

	packName, err := packService.CreatePack(pack)

	if err != nil {
		logger.Errorf("Failed to create icon pack, error: %s", err)
		render.Render(w, r, errors.ErrInternalServer(err))
		return
	}

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Created icon pack named " + packName,
	})
}

// GetPacks gets all the icon packs
func GetPacks(w http.ResponseWriter, r *http.Request) {
	packs, err := packService.GetPacks()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	if len(packs) >= 1 {
		render.JSON(w, r, &response{
			Status: "success",
			Packs:  packs,
		})
	} else {
		render.JSON(w, r, &response{
			Status:  "success",
			Message: "No icon packs found",
		})
	}
}

// GetPacksByDev renders all the packs by the given dev
func GetPacksByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	packs, err := packService.GetPacksByDev(dev)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	if len(packs) >= 1 {
		render.JSON(w, r, &response{
			Status: "success",
			Packs:  packs,
		})
	} else {
		render.JSON(w, r, &response{
			Status:  "success",
			Message: "No icon packs found for developer " + dev,
		})
	}
}

// GetPackCountByDev responds with the number of icon requests
func GetPackCountByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	count, err := packService.GetPackCountByDev(dev)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}
