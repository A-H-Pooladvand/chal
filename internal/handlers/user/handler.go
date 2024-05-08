package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"theList/internal/entities"
	"theList/internal/response"
	"theList/internal/services"
)

type Handler struct {
	Service *services.User
}

type HandlerParams struct {
	fx.In
	Service *services.User
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		Service: params.Service,
	}
}

func (h Handler) Create(c *gin.Context) {
	var request Request
	res := response.New(c)
	// Todo:: Add tracer

	if err := c.Bind(&request); err != nil {
		res.UnprocessableEntity(map[string]any{
			"error": err.Error(),
		})
		return
	}

	user := entities.NewUser(
		request.Name,
		request.Surname,
	)

	r, err := h.Service.Create(c, user)

	if err != nil {
		res.ServerError()
		return
	}

	res.SetMessage("user created successfully").Created(map[string]any{
		"user": r,
	})
}
