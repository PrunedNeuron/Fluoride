package service

import (
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/PrunedNeuron/Fluoride/pkg/store"
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

// GetIcons retrieves a list of all the icon requests in the DB
func (service *IconService) GetIcons() ([]model.Icon, error) {
	return service.iconStore.GetIcons()
}

// GetIconsByDev retrieves a list of all the icon requests belonging to the dev
func (service *IconService) GetIconsByDev(dev string) ([]model.Icon, error) {
	return service.iconStore.GetIconsByDev(dev)
}

// GetIconsByPackByDev retrieves a list of all the icon requests in the database
func (service *IconService) GetIconsByPackByDev(dev, pack string) ([]model.Icon, error) {
	return service.iconStore.GetIconsByPackByDev(dev, pack)
}

// GetPendingIconsByPackByDev retrieves a list of all the icon requests in the database
func (service *IconService) GetPendingIconsByPackByDev(dev, pack string) ([]model.Icon, error) {
	return service.iconStore.GetPendingIconsByPackByDev(dev, pack)
}

// GetDoneIconsByPackByDev retrieves a list of all the icon requests in the database
func (service *IconService) GetDoneIconsByPackByDev(dev, pack string) ([]model.Icon, error) {
	return service.iconStore.GetDoneIconsByPackByDev(dev, pack)
}

// GetIconByComponentByPackByDev returns the matching icon
func (service *IconService) GetIconByComponentByPackByDev(dev, pack, component string) (*model.Icon, error) {
	return service.iconStore.GetIconByComponentByPackByDev(dev, pack, component)
}

// SaveIcon upserts an icon
func (service *IconService) SaveIcon(dev string, icon *model.Icon) (int, error) {
	return service.iconStore.SaveIcon(dev, icon)
}

// SaveIcons upserts a list of icons
func (service *IconService) SaveIcons(dev string, icons []*model.Icon) (int, error) {
	return service.iconStore.SaveIcons(dev, icons)
}

// GetIconCountByDev retrieves the number of icons in the database
func (service *IconService) GetIconCountByDev(dev string) (int, error) {
	return service.iconStore.GetIconCountByDev(dev)
}

// GetPendingIconCountByDev retrieves the number of icons in the database
func (service *IconService) GetPendingIconCountByDev(dev string) (int, error) {
	return service.iconStore.GetPendingIconCountByDev(dev)
}

// GetDoneIconCountByDev retrieves the number of icons in the database
func (service *IconService) GetDoneIconCountByDev(dev string) (int, error) {
	return service.iconStore.GetDoneIconCountByDev(dev)
}

// UpdateIconStatus sets the new status of the icon request
func (service *IconService) UpdateIconStatus(dev, pack, component, status string) (string, error) {
	return service.iconStore.UpdateIconStatus(dev, pack, component, status)
}

// GetIconPackIDFromName sets the new status of the icon request
func (service *IconService) GetIconPackIDFromName(dev, pack string) (int, error) {
	return service.iconStore.GetIconPackIDFromName(dev, pack)
}
