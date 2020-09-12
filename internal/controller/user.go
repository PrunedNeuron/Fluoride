package controller

import (
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := new(model.User)

	err := render.DecodeJSON(r.Body, &user)

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(err))
		return
	}

	if user.Role == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing role value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing role value")))
		return
	}

	if user.Role != "developer" && user.Role != "admin" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Invalid role value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Invalid role value")))
		return
	}

	if user.Name == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing name value")))
		return
	}

	if user.Username == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing username value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing username value")))
		return
	}

	if user.Email == "" {
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("Missing email value")))
		logger.Errorw("Error: ", errors.ErrInvalidRequest(fmt.Errorf("Missing email value")))
		return
	}

	username, role, err := userService.CreateUser(user)

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Created " + role + " with username " + username,
	})
}
