package store

import (
	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/PrunedNeuron/Fluoride/pkg/database"
	"go.uber.org/zap"
)

// PackStore is the repository for the Icon Pack model (pack).
type PackStore interface {

	// Check whether givne pack exists
	PackExists(string, string) (bool, error)

	// Create new icon pack
	CreatePack(model.Pack) (string, error)

	// Gets all the icon packs in the database
	//!!!UNIMPLEMENTED
	GetPacks() ([]model.Pack, error)

	// Gets all the icon packs by the given dev
	GetPacksByDev(string) ([]model.Pack, error)

	// Gets the number of icon packs by the developer
	GetPackCountByDev(string) (int, error)
}

// NewPackStore creates and returns a new icon store instance
func NewPackStore() PackStore {
	var packStore PackStore
	var err error

	switch config.GetConfig().Database.Type {
	case "postgres":
		packStore, err = database.New()
	}

	if err != nil {
		zap.S().Errorw("Database error", "error", err)
	}
	return packStore
}
