package service

import (
	"fluoride/internal/model"
	"fluoride/internal/store"

	"go.uber.org/zap"
)

// IconService provides access to services related to icon requests
type IconService struct {
	logger    *zap.SugaredLogger
	iconStore store.IconStore
}

// NewIconService creates and returns a new icon service instance
func NewIconService(iconStore store.IconStore) *IconService {
	service := &IconService{
		logger:    zap.S().With("package", "service"),
		iconStore: iconStore,
	}

	return service

}

// GetAllIcons retrieves a list of all the icon requests in the database
func (service *IconService) GetAllIcons(pack string) ([]model.Icon, error) {
	return service.iconStore.GetAllIcons(pack)
}

// GetPendingIcons retrieves a list of all the icon requests in the database
func (service *IconService) GetPendingIcons(pack string) ([]model.Icon, error) {
	return service.iconStore.GetPendingIcons(pack)
}

// GetDoneIcons retrieves a list of all the icon requests in the database
func (service *IconService) GetDoneIcons(pack string) ([]model.Icon, error) {
	return service.iconStore.GetDoneIcons(pack)
}

// GetIconByComponent returns the matching icon
func (service *IconService) GetIconByComponent(pack, component string) (*model.Icon, error) {
	return service.iconStore.GetIconByComponent(pack, component)
}

// SaveIcon upserts an icon
func (service *IconService) SaveIcon(icon *model.Icon) (int, error) {
	return service.iconStore.SaveIcon(icon)
}

// SaveIcons upserts a list of icons
func (service *IconService) SaveIcons(icons []*model.Icon) (int, error) {
	return service.iconStore.SaveIcons(icons)
}

// GetIconCount retrieves the number of icons in the database
func (service *IconService) GetIconCount(pack string) (int, error) {
	return service.iconStore.GetIconCount(pack)
}

// GetPendingIconCount retrieves the number of icons in the database
func (service *IconService) GetPendingIconCount(pack string) (int, error) {
	return service.iconStore.GetPendingIconCount(pack)
}

// GetDoneIconCount retrieves the number of icons in the database
func (service *IconService) GetDoneIconCount(pack string) (int, error) {
	return service.iconStore.GetDoneIconCount(pack)
}

// UpdateStatus sets the new status of the icon request
func (service *IconService) UpdateStatus(pack, component, status string) (string, error) {
	return service.iconStore.UpdateStatus(pack, component, status)
}
