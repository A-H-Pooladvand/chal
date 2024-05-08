package services

import (
	"context"
	"go.uber.org/fx"
	"theList/internal/entities"
	"theList/internal/repositories"
)

type User struct {
	repository *repositories.User
}

type UserParams struct {
	fx.In
	Repository *repositories.User
}

func NewUser(params UserParams) *User {
	return &User{
		repository: params.Repository,
	}
}

func (u User) Create(ctx context.Context, user entities.UserRequest) (entities.UserResponse, error) {
	// Todo:: Add tracer

	return u.repository.Create(ctx, user)
}
