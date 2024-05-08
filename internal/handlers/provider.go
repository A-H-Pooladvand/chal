package handlers

import "theList/internal/handlers/user"

type Handlers struct {
	User *user.Handler
}

func NewHandlers(user *user.Handler) *Handlers {
	return &Handlers{User: user}
}
