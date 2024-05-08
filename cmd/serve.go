package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"theList/configs"
	"theList/internal/app"
	"theList/internal/db"
	"theList/internal/handlers"
	"theList/internal/handlers/user"
	"theList/internal/logger"
	"theList/internal/web"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the application",
	// Provide the necessary protocols such as gRPC, HTTP, etc...
	Run: runApplication,
}

func runApplication(cmd *cobra.Command, args []string) {
	app.LoadEnvironmentVariablesInLocalEnv()

	fx.New(
		user.Module,
		fx.Provide(
			// Loading configs
			configs.NewApp,
			configs.NewPostgres,
			db.New,
			// Loading services
			handlers.NewHandlers,
		),

		fx.Invoke(
			logger.Invoke,
			web.Invoke,
		),
	).Run()
}
