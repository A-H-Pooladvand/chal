package entities

import "time"

type UserRequest struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

func NewUser(name string, surname string) UserRequest {
	return UserRequest{Name: name, Surname: surname}
}

type UserResponse struct {
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserResponse(
	name string,
	surname string,
	createdAt time.Time,
	updatedAt time.Time,
) UserResponse {
	return UserResponse{
		Name:      name,
		Surname:   surname,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
