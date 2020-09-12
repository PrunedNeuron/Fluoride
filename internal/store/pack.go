package store

import (
	"fluoride/internal/model"
	"fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// PackStore is the repository for the Icon Pack model (pack).
type PackStore interface {

	// Gets all the icon packs by the developer
	GetAllPacks(string) ([]model.Pack, error)
}

// NewPackStore creates and returns a new icon store instance
func NewPackStore() PackStore {
	var packStore PackStore
	var err error

	switch viper.GetString("storage.type") {
	case "postgres":
		packStore, err = database.New()
	}

	if err != nil {
		zap.S().Fatalw("Database error", "error", err)
	}
	return packStore
}