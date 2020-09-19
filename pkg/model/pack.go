package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"

	"go.uber.org/zap"
)

// Pack is the icon pack type
type Pack struct {
	ID            int       `json:"id" boil:"id" db:"id"`                                                 // Icon pack ID
	Name          string    `json:"name" boil:"name" db:"name"`                                           // Icon pack name
	DevUsername   string    `json:"developer_username" boil:"developer_username" db:"developer_username"` // Icon pack developer username
	URL           string    `json:"url" boil:"url" db:"url"`                                              // Icon pack url (play store)
	BillingStatus string    `json:"billing_status" boil:"billing_status" db:"billing_status"`             // Billing status
	CreatedAt     time.Time `json:"created_at" boil:"created_at" db:"created_at"`                         // Date when the pack was added
	UpdatedAt     time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`                         // Date when the pack was updated
}

func (pack *Pack) String() string {
	json, err := json.Marshal(pack)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
