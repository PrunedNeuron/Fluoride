package cmd

import (
	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/PrunedNeuron/Fluoride/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	Long:    "Start the server and respond to requests",
	Aliases: []string{"server", "api"},
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Infow("Application version " + config.GetConfig().Application.Version)
		// Create server
		logger.Info("Creating server")
		server, err := server.New()
		if err != nil {
			logger.Errorf("Failed to start server, error: ", err.Error())
		}

		if err = server.Serve(); err != nil {
			logger.Fatalw("Could not start the server", "error", err)
		}

		<-config.Stop.Chan() // wait until stop channel
		config.Stop.Wait()   // wait until everything else has gracefully exited
		zap.L().Sync()       // Flush the logger
	},
}
