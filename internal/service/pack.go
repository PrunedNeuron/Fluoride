package service

import (
	"fluoride/internal/model"
	"fluoride/internal/store"

	"go.uber.org/zap"
)

// PackService provides access to services related to icon packs
type PackService struct {
	logger    *zap.SugaredLogger
	packStore store.PackStore
}

// NewPackService creates and returns a new pack service instance
func NewPackService(packStore store.PackStore) *PackService {
	service := &PackService{
		logger:    zap.S().With("package", "service"),
		packStore: packStore,
	}

	return service

}

// CreatePack creates a new icon pack
func (service *PackService) CreatePack(pack model.Pack) (string, error) {
	return service.packStore.CreatePack(pack)
}

// GetPacksByDev gets the list of all icon packs by the given dev
func (service *PackService) GetPacksByDev(dev string) ([]model.Pack, error) {
	return service.packStore.GetPacksByDev(dev)
}

// GetPackCountByDev gets the number of icon packs by the dev
func (service *PackService) GetPackCountByDev(dev string) (int, error) {
	return service.packStore.GetPackCountByDev(dev)
}
