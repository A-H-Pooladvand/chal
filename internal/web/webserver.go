package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"theList/configs"
	"theList/internal/handlers"
	"theList/internal/logger"
	"theList/routes"
	"time"
)

func Invoke(lc fx.Lifecycle, c *configs.App, w *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	routes.RegisterWebRoutes(router, w)
	RegisterMiddlewares(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.AppPort),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Printf("server up and running on http://127.0.0.1:%s\n", c.AppPort)
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Fatal(err.Error(), zap.Error(err))
				}

				quit := make(chan os.Signal)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
			}()

			return nil
		},

		// graceful shutdown
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				return err
			}

			return nil
		},
	})

	return router
}
