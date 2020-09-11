package cmd

import (
	"fluoride/config"
	"fluoride/pkg/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	Long:    "Start the server and respond to requests",
	Aliases: []string{"server", "api"},
	Run: func(cmd *cobra.Command, args []string) {
		// Create server
		server, err := server.New()
		if err != nil {
			logger.Errorf("Failed to start server, error: ", err.Error())
		}
		err = server.Serve()

		if err != nil {
			logger.Fatalw("Could not start the server", "error", err)
		}

		<-config.Stop.Chan() // wait until StopChan
		config.Stop.Wait()   // wait until everyone cleans up
		zap.L().Sync()       // Flush the logger
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
