package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// Icon is the icon request type
// swagger:model
type Icon struct {
	// Icon request ID
	// example: 51
	ID int `json:"id" boil:"id" db:"id"`

	// App name
	// example: Canva
	Name string `json:"name" boil:"name" db:"name"`

	// App component
	// example: com.canva.editor/com.canva.app.editor.splash.SplashActivity
	Component string `json:"component" boil:"component" db:"component"`

	// Play Store URL (may be autogenerated)
	// example: https://play.google.com/store/apps/details?id=com.canva.editor
	URL string `json:"url" boil:"url" db:"url"`

	// Number of requesters
	// example: 28
	Requesters int `json:"requesters" boil:"requesters" db:"requesters"`

	// Status of request (pending / complete)
	// example: pending
	Status string `json:"status" boil:"status" db:"status"`

	// Name of the  Icon pack it belongs to
	// example: Valacons
	Pack string `json:"pack" boil:"pack" db:"icon_pack_name"`

	// Date created at
	// example: 2020-09-17T03:07:13.418204+05:30
	CreatedAt time.Time `json:"created_at" boil:"created_at" db:"created_at"`

	// Date updated at
	// example: 2020-09-17T03:07:13.418204+05:30
	UpdatedAt time.Time `json:"updated_at" boil:"updated_at" db:"updated_at"`
}

func (icon *Icon) String() string {
	json, err := json.Marshal(icon)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
