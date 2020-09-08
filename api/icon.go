package api

import (
	"context"
	"fmt"
)

// IconStore is the repo for the Icon model.
type IconStore interface {
	GetIcons(context.Context) ([]Icon, error)
	GetIconByComponent(context.Context, string) (*Icon, error)
}

type Icon struct {
	ID         string `json:"id" boil:"id" db:"id"`
	Name       string `json:"name" boil:"name" db:"name"`
	Component  string `json:"component" boil:"component" db:"component"`
	URL        string `json:"url" boil:"url" db:"url"`
	Requesters string `json:"requesters" boil:"requesters" db:"requesters"`
	Status     string `json:"status" boil:"status" db:"status"`
}

// Show displays the icon request
func (i *Icon) Show() {
	fmt.Printf("%s, %s\n", i.Name, i.Component)
}

// ToString returns a string representation
func (i *Icon) ToString() string {
	return fmt.Sprintf("Name: %s\nComponent: %s", i.Name, i.Component)
}
