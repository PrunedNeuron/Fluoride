package cmd

import (
	"icon-requests/api/controller"
	"icon-requests/config"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Start the API",
		Long:  "Prepare and setup the API and launch it",
		Run: func(cmd *cobra.Command, args []string) { // Initialize the database
			// Create the server (uses wire DI)
			server, err := NewServer()
			if err != nil {
				logger.Fatalw("Could not create server", "error", err)
			}

			// Setup and register the icon request server
			iconStore := NewIconStore()
			controller.Setup(server.Router(), iconStore)

			err = server.ListenAndServe()
			if err != nil {
				logger.Fatalw("Could not start server", "error", err)
			}

			<-config.Stop.Chan() // wait until StopChan
			config.Stop.Wait()   // wait until everyone cleans up
			zap.L().Sync()       // Flush the logger

		},
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)
}
