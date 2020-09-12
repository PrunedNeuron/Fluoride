package controller

import (
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetAllIcons responds with a list of all the icons
func GetAllIcons(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}
	list, err := iconService.GetAllIcons(pack)
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

// GetPendingIcons responds with a list of all the icons
func GetPendingIcons(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}
	list, err := iconService.GetPendingIcons(pack)
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

// GetDoneIcons responds with a list of all the icons
func GetDoneIcons(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}
	list, err := iconService.GetDoneIcons(pack)
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
	// Get pack from url
	pack := chi.URLParam(r, "pack")
	// Get component
	component := chi.URLParam(r, "component")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}
	if component == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid component")))
		return
	}

	icon, err := iconService.GetIconByComponent(pack, component)

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
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	var icon = new(model.Icon)

	if err := render.DecodeJSON(r.Body, &icon); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	if icon.Pack != pack {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("icon pack mismatch")))
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
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	var icons []*model.Icon

	if err := render.DecodeJSON(r.Body, &icons); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	for _, icon := range icons {
		if icon.Pack != pack {
			render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("icon pack mismatch")))
			return
		}
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

// GetIconCount responds with the number of icon requests
func GetIconCount(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	count, err := iconService.GetIconCount(pack)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetPendingIconCount responds with the number of icon requests
func GetPendingIconCount(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	count, err := iconService.GetPendingIconCount(pack)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetDoneIconCount responds with the number of icon requests
func GetDoneIconCount(w http.ResponseWriter, r *http.Request) {
	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	count, err := iconService.GetDoneIconCount(pack)
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

	if req.Status != "pending" && req.Status != "done" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Invalid status value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Invalid status value")))
		return
	}

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	status, err := iconService.UpdateStatus(pack, req.Component, req.Status)

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Updated status to " + status + " for icon request with component " + req.Component,
	})
}
