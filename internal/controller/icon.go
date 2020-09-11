package controller

import (
	"fluoride/internal/model"
	"fluoride/internal/service"
	"fluoride/internal/store"
	"fluoride/pkg/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type response struct {
	Status  string       `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	Count   int          `json:"count,omitempty"`
	Icons   []model.Icon `json:"icons,omitempty"`
}

var (
	logger      = zap.S().With("package", "controller.icon")
	iconStore   = store.NewIconStore()
	iconService = service.NewIconService(iconStore)
)

// GetIcons responds with a list of all the icons
func GetIcons(w http.ResponseWriter, r *http.Request) {
	list, err := iconService.GetAllIcons()
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status:  "success",
		Message: "retrieved " + strconv.Itoa(len(list)) + " icons",
		Icons:   list,
	})
}

// GetIconByComponent responds with the matching icon
func GetIconByComponent(w http.ResponseWriter, r *http.Request) {
	// Get component
	component := chi.URLParam(r, "component")

	if component == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid component")))
		return
	}

	icon, err := iconService.GetIconByComponent(component)

	if err == errors.ErrDatabaseNotFound {
		render.Render(w, r, errors.ErrNotFound)
	} else if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
	} else {
		render.JSON(w, r, icon)
	}

}

// SaveIcon saves the icon to the database
func SaveIcon(w http.ResponseWriter, r *http.Request) {

	var icon = new(model.Icon)
	if err := render.DecodeJSON(r.Body, &icon); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	id, err := iconService.SaveIcon(icon)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Message: "Inserted icon with id = " + strconv.Itoa(id),
	})
}

// SaveIcons saves the list of icons to the database
func SaveIcons(w http.ResponseWriter, r *http.Request) {

	var icons []*model.Icon

	if err := render.DecodeJSON(r.Body, &icons); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	count, err := iconService.SaveIcons(icons)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Inserted " + strconv.Itoa(count) + " icons successfully.",
	})
}

// GetCount responds with the number of icon requests
func GetCount(w http.ResponseWriter, r *http.Request) {
	count, err := iconService.GetCount()
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// UpdateStatus takes the new status and updates the database
func UpdateStatus(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Component string `json:"component"`
		Status    string `json:"status"`
	}

	req := new(request)

	err := render.DecodeJSON(r.Body, &req)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	if req.Component == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing component value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing component value")))
		return
	}

	if req.Status == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing status value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing status value")))
		return
	}

	if req.Status != "pending" && req.Status != "complete" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Invalid status value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Invalid status value")))
		return
	}

	status, err := iconService.UpdateStatus(req.Component, req.Status)

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Updated status to " + status + " for icon request with component " + req.Component,
	})
}
