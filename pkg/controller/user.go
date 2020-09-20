package controller

import (
	"fmt"
	"net/http"

	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/go-chi/render"
)

// GetUsers renders all the users in the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetUsers()

	if err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, &response{
		Status: "success",
		Users:  users,
	})
}

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

	if err != nil {
		logger.Errorw("Failed to create new user, error: %s", err)
		render.Render(w, r, errors.ErrInvalidRequest(fmt.Errorf("user may already exist")))
		return
	}

	render.JSON(w, r, &response{
		Status:  "success",
		Message: "Created " + role + " with username " + username,
	})
}
