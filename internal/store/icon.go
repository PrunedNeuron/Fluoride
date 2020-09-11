package store

import (
	"fluoride/internal/model"
	"fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// IconStore is the repository for the Icon model.
type IconStore interface {
	GetAllIcons() ([]model.Icon, error)
	GetIconByComponent(string) (*model.Icon, error)
	SaveIcon(*model.Icon) (int, error)
	SaveIcons([]*model.Icon) (int, error)
	GetCount() (int, error)
	UpdateStatus(string, string) (string, error)
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
