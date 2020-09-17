package store

import (
	"github.com/PrunedNeuron/Fluoride/internal/model"
	"github.com/PrunedNeuron/Fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// IconStore is the repository for the Icon model.
type IconStore interface {

	// Gets all the icons in the database
	GetIcons() ([]model.Icon, error)

	// Gets all icons by the developer
	GetIconsByDev(string) ([]model.Icon, error)

	// Gets all the icons by the developer
	GetIconsByPackByDev(string, string) ([]model.Icon, error)

	// Gets icons with status = pending
	GetPendingIconsByPackByDev(string, string) ([]model.Icon, error)

	// Gets icons with status = done
	GetDoneIconsByPackByDev(string, string) ([]model.Icon, error)

	// Gets icons with component = given string
	GetIconByComponentByPackByDev(string, string, string) (*model.Icon, error)

	// Saves a single icon
	SaveIcon(string, *model.Icon) (int, error)

	// Saves an array of icons
	SaveIcons(string, []*model.Icon) (int, error)

	// Gets the number of icons
	GetIconCountByDev(string) (int, error)

	// Gets the number of pending icons
	GetPendingIconCountByDev(string) (int, error)

	// Gets the number of done icons
	GetDoneIconCountByDev(string) (int, error)

	// Updates the status of the icon with the given component with the given status
	UpdateIconStatus(string, string, string, string) (string, error)

	GetIconPackIDFromName(string, string) (int, error)
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
