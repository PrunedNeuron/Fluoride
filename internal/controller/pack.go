package controller

import (
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// CreatePack creates a new icon pack
func CreatePack(w http.ResponseWriter, r *http.Request) {

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
		render.Render(w, r, errors.ErrInternalServer(fmt.Errorf("Failed to create icon pack")))
	}

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Created icon pack named " + packName,
	})
}
