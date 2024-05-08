package repositories

import (
	"context"
	"go.uber.org/fx"
	"theList/internal/entities"
	"theList/internal/models"
	"theList/pkg/postgres"
)

type User struct {
	db *postgres.Client
}

type UserParams struct {
	fx.In
	DB *postgres.Client
}

func NewUser(params UserParams) *User {
	return &User{
		db: params.DB,
	}
}

func (u User) Create(ctx context.Context, user entities.UserRequest) (response entities.UserResponse, err error) {
	// Todo:: Add tracer

	model := models.User{
		FirstName: user.Name,
		LastName:  user.Surname,
	}

	err = u.db.WithContext(ctx).Create(&model).InsertError()

	response = entities.NewUserResponse(
		model.FirstName,
		model.LastName,
		model.CreatedAt,
		model.UpdatedAt,
	)

	return response, err
}
