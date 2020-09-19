package store

import (
	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/PrunedNeuron/Fluoride/pkg/database"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"go.uber.org/zap"
)

// UserStore is the repository for the generic user model.
type UserStore interface {

	// Gets the list of all users
	GetUsers() ([]model.User, error)
	// Creates a new user
	CreateUser(*model.User) (string, string, error)
}

// NewUserStore creates and returns a new user store instance
func NewUserStore() UserStore {
	var userStore UserStore
	var err error

	switch config.GetConfig().Database.Type {
	case "postgres":
		userStore, err = database.New()
	}

	if err != nil {
		zap.S().Errorw("Database error", "error", err)
	}
	return userStore
}
