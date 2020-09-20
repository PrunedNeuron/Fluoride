package store

import (
	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/PrunedNeuron/Fluoride/pkg/model"
	"github.com/PrunedNeuron/Fluoride/pkg/database"
	"go.uber.org/zap"
)

// PlanStore is the repository for the plan model
type PlanStore interface {

	// Check whether the plans table exists
	PlansExists() (bool, error)

	// Create new plan
	CreatePlan(*model.Plan) (string, error)

	// Gets all the icon packs by the developer
	GetPlans() ([]model.Plan, error)
}

// NewPlanStore creates and returns a new plan store instance
func NewPlanStore() PlanStore {
	var planStore PlanStore
	var err error

	switch config.GetConfig().Database.Type {
	case "postgres":
		planStore, err = database.New()
	}

	if err != nil {
		zap.S().Errorw("Database error", "error", err)
	}
	return planStore
}
