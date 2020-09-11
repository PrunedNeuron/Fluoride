package store

import (
	"fluoride/internal/model"
	"fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// IconStore is the repository for the Icon model.
type IconStore interface {

	// Gets all the icons in the DB
	GetAllIcons(string) ([]model.Icon, error)

	// Gets icons with status = pending
	GetPendingIcons(string) ([]model.Icon, error)

	// Gets icons with status = done
	GetDoneIcons(string) ([]model.Icon, error)

	// Gets icons with component = given string
	GetIconByComponent(string, string) (*model.Icon, error)

	// Saves a single icon
	SaveIcon(*model.Icon) (int, error)

	// Saves an array of icons
	SaveIcons([]*model.Icon) (int, error)

	// Gets the number of icons
	GetIconCount(string) (int, error)

	// Gets the number of pending icons
	GetPendingIconCount(string) (int, error)

	// Gets the number of done icons
	GetDoneIconCount(string) (int, error)

	// Updates the status of the icon with the given component with the given status
	UpdateStatus(string, string, string) (string, error)
}

// NewIconStore creates and returns a new icon store instance
func NewIconStore() IconStore {
	var iconStore IconStore
	var err error

	switch viper.GetString("storage.type") {
	case "postgres":
		iconStore, err = database.New()
	}

	if err != nil {
		zap.S().Fatalw("Database error", "error", err)
	}
	return iconStore
}
