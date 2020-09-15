package controller

import (
	"fluoride/config"
	"fluoride/internal/model"
	"fluoride/internal/service"
	"fluoride/internal/store"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type response struct {
	Status  string       `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	Count   int          `json:"count,omitempty"`
	Icons   []model.Icon `json:"icons,omitempty"`
	Packs   []model.Pack `json:"packs,omitempty"`
	Devs    []model.User `json:"developers,omitempty"`
	Admins  []model.User `json:"admins,omitempty"`
	Users   []model.User `json:"users,omitempty"`
	Plans   []model.Plan `json:"plans,omitempty"`
}

var (
	logger      = zap.S().With("package", "controller.icon")
	iconStore   = store.NewIconStore()
	packStore   = store.NewPackStore()
	devStore    = store.NewDevStore()
	userStore   = store.NewUserStore()
	planStore   = store.NewPlanStore()
	iconService = service.NewIconService(iconStore)
	packService = service.NewPackService(packStore)
	devService  = service.NewDevService(devStore)
	userService = service.NewUserService(userStore)
	planService = service.NewPlanService(planStore)
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
