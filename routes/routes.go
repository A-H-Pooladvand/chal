package routes

import (
	"github.com/gin-gonic/gin"
	"theList/internal/handlers"
)

func RegisterWebRoutes(e *gin.Engine, w *handlers.Handlers) {
	v1 := e.Group("api/v1")

	usersGroup := v1.Group("users")
	{
		usersGroup.POST("", w.User.Create)
	}
}
