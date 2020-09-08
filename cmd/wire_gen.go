// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"icon-requests/api"
	"icon-requests/database/postgres"
	"icon-requests/server"

	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewServer() (*server.Server, error) {
	serverServer, err := server.New()
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}

// wire.go:

// Create a new thing store
func NewIconStore() api.IconStore {
	var iconStore api.IconStore
	var err error

	switch viper.GetString("storage.type") {
	case "postgres":
		iconStore, err = postgres.New()
	}

	if err != nil {
		logger.Fatalw("Database error", "error", err)
	}
	return iconStore
}
