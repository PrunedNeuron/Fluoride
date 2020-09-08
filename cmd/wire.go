// +build wireinject

package cmd

import (
	"icon-requests/api"
	"icon-requests/database/postgres"
	"icon-requests/server"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Create a new server
func NewServer() (*server.Server, error) {
	wire.Build(server.New)
	return &server.Server{}, nil
}

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
