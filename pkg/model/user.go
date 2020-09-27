package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// User is the generic user type
// swagger:model
type User struct {
	// User ID
	// example: 24
	ID int `json:"id" boil:"id" db:"id"`

	// User role (admin | developer)
	// example: developer
	Role string `json:"role" boil:"role" db:"role"`

	// User name
	// example: John Doe
	Name string `json:"name" boil:"name" db:"name"`

	// User username
	// example: jdoe
	Username string `json:"username" boil:"username" db:"username"`

	// User email
	// example: jdoe@gmail.com
	Email string `json:"email" boil:"email" db:"email"`

	// User website
	// example: https://jdoe.co
	URL string `json:"url" boil:"url" db:"url"`

	// Date when the User was added
	// example: 2020-09-17T03:07:13.418204+05:30
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"`

	// Date when the User was updated
	// example: 2020-09-17T03:07:13.418204+05:30
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`
}

func (user *User) String() string {
	json, err := json.Marshal(user)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
