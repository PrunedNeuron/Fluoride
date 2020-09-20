package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"

	"go.uber.org/zap"
)

// User is the generic user type
type User struct {
	ID        int       `json:"id" boil:"id" db:"id"`                         // User ID
	Role      string    `json:"role" boil:"role" db:"role"`                   // User role (admin | developer)
	Name      string    `json:"name" boil:"name" db:"name"`                   // User name
	Username  string    `json:"username" boil:"username" db:"username"`       // User username
	Email     string    `json:"email" boil:"email" db:"email"`                // User email
	URL       string    `json:"url" boil:"url" db:"url"`                      // User website
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"` // Date when the User was added
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"` // Date when the User was updated
}

func (user *User) String() string {
	json, err := json.Marshal(user)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
