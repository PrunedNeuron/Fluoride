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

// GetIconsByDev responds with a list of all the icons
func GetIconsByDev(w http.ResponseWriter, r *http.Request) {

	// Get dev from url
	dev := chi.URLParam(r, "developer")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	list, err := iconService.GetIconsByDev(dev)
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

// GetIconsByPackByDev responds with a list of all the icons
func GetIconsByPackByDev(w http.ResponseWriter, r *http.Request) {

	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	list, err := iconService.GetIconsByPackByDev(dev, pack)
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

// GetPendingIconsByPackByDev responds with a list of all the icons
func GetPendingIconsByPackByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	list, err := iconService.GetPendingIconsByPackByDev(dev, pack)
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

// GetDoneIconsByPackByDev responds with a list of all the icons
func GetDoneIconsByPackByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	list, err := iconService.GetDoneIconsByPackByDev(dev, pack)
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

// GetIconByComponentByPackByDev responds with the matching icon
func GetIconByComponentByPackByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	// Get component from url
	component := chi.URLParam(r, "component")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	if component == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid component")))
		return
	}

	icon, err := iconService.GetIconByComponentByPackByDev(dev, pack, component)

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
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

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

	id, err := iconService.SaveIcon(dev, icon)
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
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	// Get pack from url
	pack := chi.URLParam(r, "pack")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

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

	count, err := iconService.SaveIcons(dev, icons)
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

// GetIconCountByDev responds with the number of icon requests
func GetIconCountByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "developer")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	count, err := iconService.GetIconCountByDev(dev)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetPendingIconCountByDev responds with the number of icon requests
func GetPendingIconCountByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "dev")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	count, err := iconService.GetPendingIconCountByDev(dev)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// GetDoneIconCountByDev responds with the number of icon requests
func GetDoneIconCountByDev(w http.ResponseWriter, r *http.Request) {
	// Get dev from url
	dev := chi.URLParam(r, "dev")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	count, err := iconService.GetDoneIconCountByDev(dev)
	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Count:  count,
	})
}

// UpdateIconStatus takes the new status and updates the database
func UpdateIconStatus(w http.ResponseWriter, r *http.Request) {

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

	// Get dev from url
	dev := chi.URLParam(r, "dev")

	if dev == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid dev")))
		return
	}

	if pack == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("invalid pack")))
		return
	}

	status, err := iconService.UpdateIconStatus(dev, pack, req.Component, req.Status)

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Updated status to " + status + " for icon request with component " + req.Component,
	})
}
