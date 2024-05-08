package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"theList/internal/logger"
)

var cmd = &cobra.Command{
	Use:   "app",
	Short: "webserver",
	Long:  `Initializing...`,
}

func Execute() {
	cmd.AddCommand(serveCmd)
	cmd.AddCommand(migrateCmd)

	if err := cmd.Execute(); err != nil {
		logger.Panic(err.Error(), zap.Error(err))
	}
}
