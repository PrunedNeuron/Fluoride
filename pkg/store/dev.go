package store

import (
	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/PrunedNeuron/Fluoride/pkg/database"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"go.uber.org/zap"
)

// DevStore is the repository for the generic user model.
type DevStore interface {

	// Check whether the developer exists
	DevExists(string) (bool, error)

	// Get all the developers in the database
	GetDevs() ([]model.User, error)

	// Get the number of developers in the database
	GetDevCount() (int, error)

	// Get the developer with the given username
	GetDevByUsername(string) (model.User, error)
}

// NewDevStore creates and returns a new dev store instance
func NewDevStore() DevStore {
	var devStore DevStore
	var err error

	switch config.GetConfig().Database.Type {
	case "postgres":
		devStore, err = database.New()
	}

	if err != nil {
		zap.S().Errorw("Database error", "error", err)
	}
	return devStore
}
