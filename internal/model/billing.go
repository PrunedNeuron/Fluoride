package model

import (
	"encoding/json"
	"fluoride/pkg/errors"
	"fmt"

	"go.uber.org/zap"
)

// Billing is the junction type of User and Plan
type Billing struct {
	DevID  int `json:"dev_id" boil:"dev_id" db:"dev_id"`    // ID of the associated developer
	PlanID int `json:"plan_id" boil:"plan_id" db:"plan_id"` // ID of the plan being billed
}

func (billing *Billing) String() string {
	json, err := json.Marshal(billing)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
