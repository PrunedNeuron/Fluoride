package model

import (
	"encoding/json"
	"fluoride/pkg/errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Plan is the plan type
type Plan struct {
	ID           int       `json:"id" boil:"id" db:"id"`                                  // Plan ID
	Name         string    `json:"name" boil:"name" db:"name"`                            // Plan name
	Description  string    `json:"description" boil:"description" db:"description"`       // Plan description
	Intro        string    `json:"intro" boil:"intro" db:"intro"`                         // Plan introduction
	Price        string    `json:"price" boil:"price" db:"price"`                         // Plan price per month
	BillingCycle int       `json:"billing_cycle" boil:"billing_cycle" db:"billing_cycle"` // Billing cycle in days
	CreatedAt    time.Time `json:"created_at" boil:"created_at" db:"created_at"`          // Date when the plan was added
	UpdatedAt    time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`          // Date when the plan was updated
}

func (plan *Plan) String() string {
	json, err := json.Marshal(plan)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
