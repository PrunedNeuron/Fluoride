package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// Pack is the icon pack type
// swagger:model
type Pack struct {
	// Icon pack ID
	// example: 2
	ID int `json:"id" boil:"id" db:"id"`

	// Icon pack name
	// example: Valacons
	Name string `json:"name" boil:"name" db:"name"`

	// Icon pack developer username
	// example: jdoe
	DevUsername string `json:"developer_username" boil:"developer_username" db:"developer_username"`

	// Icon pack url (play store)
	// example: https://play.google.com
	URL string `json:"url" boil:"url" db:"url"`

	// Billing status
	// example: billed
	BillingStatus string `json:"billing_status" boil:"billing_status" db:"billing_status"`

	// Date when the pack was added
	// example: 2020-09-17T03:07:13.418204+05:30
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"`

	// Date when the pack was updated
	// example: 2020-09-17T03:07:13.418204+05:30
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`
}

func (pack *Pack) String() string {
	json, err := json.Marshal(pack)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
