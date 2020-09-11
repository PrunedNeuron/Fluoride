package model

import (
	"encoding/json"
	"fluoride/pkg/errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// !DEPRECATED

// Dev is the icon pack developer type
type Dev struct {
	ID        int       `json:"id" boil:"id" db:"id"`                         // Developer ID
	Name      string    `json:"name" boil:"name" db:"name"`                   // Developer name
	Username  string    `json:"username" boil:"username" db:"username"`       // Developer username
	Email     string    `json:"email" boil:"email" db:"email"`                // Developer email
	URL       string    `json:"url" boil:"url" db:"url"`                      // Developer website
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"` // Date when the developer was added
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"` // Date when the developer was updated
}

func (dev *Dev) String() string {
	json, err := json.Marshal(dev)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
