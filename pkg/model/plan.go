package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// Plan is the plan type
// swagger:model
type Plan struct {
	// Plan ID
	// example: 5
	ID int `json:"id" boil:"id" db:"id"`

	// Plan name
	// example: John Doe
	Name string `json:"name" boil:"name" db:"name"`

	// Plan description
	// example: Best suited for experienced icon pack devs.
	Description string `json:"description" boil:"description" db:"description"`

	// Plan introduction
	// example: Pro
	Intro string `json:"intro" boil:"intro" db:"intro"`

	// Plan price per month
	// example: $1.99
	Price string `json:"price" boil:"price" db:"price"`

	// Billing cycle in days
	// example: 30
	BillingCycle int `json:"billing_cycle" boil:"billing_cycle" db:"billing_cycle"`

	// Date when the plan was added
	// example: 2020-09-17T03:07:13.418204+05:30
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"`

	// Date when the plan was updated
	// example: 2020-09-17T03:07:13.418204+05:30
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`
}

func (plan *Plan) String() string {
	json, err := json.Marshal(plan)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
