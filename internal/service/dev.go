package service

import (
	"github.com/PrunedNeuron/Fluoride/internal/model"
	"github.com/PrunedNeuron/Fluoride/internal/store"

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

// DevExists returns whether the given dev exists
func (service *DevService) DevExists(dev string) (bool, error) {
	return service.devStore.DevExists(dev)
}

// GetDevs gets all the users that are developers
func (service *DevService) GetDevs() ([]model.User, error) {
	return service.devStore.GetDevs()
}

// GetDevCount gets the number of developers in the database
func (service *DevService) GetDevCount() (int, error) {
	return service.devStore.GetDevCount()
}

// GetDevByUsername gets the dev with the matching username
func (service *DevService) GetDevByUsername(dev string) (model.User, error) {
	return service.devStore.GetDevByUsername(dev)
}
