package service

import (
	"fluoride/internal/model"
	"fluoride/internal/store"
	"fluoride/pkg/errors"

	"go.uber.org/zap"
)

// DevService provides access to services related to users
type DevService struct {
	logger   *zap.SugaredLogger
	devStore store.DevStore
}

// NewDevService creates and returns a new user service instance
func NewDevService(devStore store.DevStore) *DevService {
	service := &DevService{
		logger:   zap.S().With("package", "service"),
		devStore: devStore,
	}

	return service

}

// GetPacksByDev creates a new user
func (service *DevService) GetPacksByDev(dev string) ([]model.Pack, error) {
	devExists, err := service.devStore.DevExists(dev)

	if err != nil {
		return nil, err
	}

	if devExists {
		return service.devStore.GetPacksByDev(dev)
	}

	return nil, errors.ErrDatabaseRelationNotFound
}
