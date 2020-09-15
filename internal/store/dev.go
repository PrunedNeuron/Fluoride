package store

import (
	"fluoride/internal/model"
	"fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// DevStore is the repository for the generic user model.
type DevStore interface {

	// Check whether the developer exists
	DevExists(string) (bool, error)

	// Get all the developers in the database
	GetDevs() ([]model.User, error)

	// Get packs by the developer
	GetPacksByDev(string) ([]model.Pack, error)

	// Get icon requests belonging to the developer
	GetIconsByDev(string) ([]model.Icon, error)

	// Get the number of developers in the database
	GetDevCount() (int, error)

	// Get the developer with the given username
	GetDevByUsername(string) (model.User, error)
}

// NewDevStore creates and returns a new dev store instance
func NewDevStore() DevStore {
	var devStore DevStore
	var err error

	switch viper.GetString("storage.type") {
	case "postgres":
		devStore, err = database.New()
	}

	if err != nil {
		zap.S().Fatalw("Database error", "error", err)
	}
	return devStore
}
