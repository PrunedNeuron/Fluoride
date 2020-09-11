package model

import (
	"encoding/json"
	"fluoride/pkg/errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// !DEPRECATED

// Admin is the admin type
type Admin struct {
	ID        int       `json:"id" boil:"id" db:"id"`                         // Admin ID
	Name      string    `json:"name" boil:"name" db:"name"`                   // Admin name
	Username  string    `json:"username" boil:"username" db:"username"`       // Admin username
	Email     string    `json:"email" boil:"email" db:"email"`                // Admin email
	URL       string    `json:"url" boil:"url" db:"url"`                      // Admin website
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"` // Date when the admin was added
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"` // Date when the admin was updated
}

func (admin *Admin) String() string {
	json, err := json.Marshal(admin)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
