package service

import (
	"fluoride/internal/model"
	"fluoride/internal/store"

	"go.uber.org/zap"
)

// UserService provides access to services related to users
type UserService struct {
	logger    *zap.SugaredLogger
	userStore store.UserStore
}

// NewUserService creates and returns a new user service instance
func NewUserService(userStore store.UserStore) *UserService {
	service := &UserService{
		logger:    zap.S().With("package", "service"),
		userStore: userStore,
	}

	return service

}

// GetUsers gets the list of all users in the database
func (service *UserService) GetUsers() ([]model.User, error) {
	return service.userStore.GetUsers()
}

// CreateUser creates a new user
func (service *UserService) CreateUser(user *model.User) (string, string, error) {
	return service.userStore.CreateUser(user)
}
