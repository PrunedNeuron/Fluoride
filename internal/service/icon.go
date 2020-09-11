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
func (service *IconService) GetAllIcons() ([]model.Icon, error) {
	return service.iconStore.GetAllIcons()
}

// GetIconByComponent returns the matching icon
func (service *IconService) GetIconByComponent(component string) (*model.Icon, error) {
	return service.iconStore.GetIconByComponent(component)
}

// SaveIcon upserts an icon
func (service *IconService) SaveIcon(icon *model.Icon) (int, error) {

	return service.iconStore.SaveIcon(icon)

}

// SaveIcons upserts a list of icons
func (service *IconService) SaveIcons(icons []*model.Icon) (int, error) {
	return service.iconStore.SaveIcons(icons)
}

// GetCount retrieves the number of icons in the database
func (service *IconService) GetCount() (int, error) {
	return service.iconStore.GetCount()
}
