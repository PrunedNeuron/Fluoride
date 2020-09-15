package store

import (
	"fluoride/internal/model"
	"fluoride/pkg/database"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// PlanStore is the repository for the plan model
type PlanStore interface {

	// Create new plan
	CreatePlan(*model.Plan) (string, error)

	// Gets all the icon packs by the developer
	GetPlans() ([]model.Plan, error)
}

// NewPlanStore creates and returns a new plan store instance
func NewPlanStore() PlanStore {
	var planStore PlanStore
	var err error

	switch viper.GetString("storage.type") {
	case "postgres":
		planStore, err = database.New()
	}

	if err != nil {
		zap.S().Fatalw("Database error", "error", err)
	}
	return planStore
}
