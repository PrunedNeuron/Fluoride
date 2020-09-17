package service

import (
	"github.com/PrunedNeuron/Fluoride/internal/model"
	"github.com/PrunedNeuron/Fluoride/internal/store"

	"go.uber.org/zap"
)

// PlanService provides access to services related to users
type PlanService struct {
	logger    *zap.SugaredLogger
	planStore store.PlanStore
}

// NewPlanService creates and returns a new user service instance
func NewPlanService(planStore store.PlanStore) *PlanService {
	service := &PlanService{
		logger:    zap.S().With("package", "service"),
		planStore: planStore,
	}

	return service

}

// PlansExists returns whether the plans table exists
func (service *PlanService) PlansExists() (bool, error) {
	return service.planStore.PlansExists()
}

// CreatePlan creates a new user
func (service *PlanService) CreatePlan(plan *model.Plan) (string, error) {
	return service.planStore.CreatePlan(plan)
}

// GetPlans gets the list of all plans
func (service *PlanService) GetPlans() ([]model.Plan, error) {
	return service.planStore.GetPlans()
}
