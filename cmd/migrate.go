package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"theList/configs"
	"theList/internal/app"
	"theList/internal/db"
	"theList/internal/logger"
	"theList/internal/models"
	"theList/pkg/postgres"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs the migrations",
	Run:   runMigrations,
}

func runMigrations(cmd *cobra.Command, args []string) {
	app.LoadEnvironmentVariablesInLocalEnv()

	application := fx.New(
		fx.Provide(
			// Loading configs
			configs.NewApp,
			configs.NewPostgres,
			// Loading services
			db.New,
		),

		fx.Invoke(
			logger.Invoke,
			//apm.Invoke,
			func(db *postgres.Client) {
				err := db.AutoMigrate(
					migrations()...,
				)

				if err != nil {
					zap.L().Fatal("failed to run the migrations", zap.Error(err))
				}
			},
		),
	)

	if err := application.Start(context.Background()); err != nil {
		zap.L().Fatal("failed to start the application", zap.Error(err))

		return
	}
	fmt.Println("All migrations completed successfully")
}

func migrations() []any {
	return []any{
		models.User{},
	}
}
