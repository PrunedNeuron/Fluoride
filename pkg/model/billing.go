package model

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// Billing is the junction type of User and Plan
// swagger:model
type Billing struct {
	// ID of the associated developer
	// example: 3
	DevID int `json:"dev_id" boil:"dev_id" db:"dev_id"`

	// ID of the plan being billed
	// example: 1
	PlanID int `json:"plan_id" boil:"plan_id" db:"plan_id"`
}

func (billing *Billing) String() string {
	json, err := json.Marshal(billing)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
