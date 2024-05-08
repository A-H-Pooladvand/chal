package user

type Request struct {
	Name           string `json:"name" binding:"required"`
	Surname        string `json:"surname" binding:"required"`
	SomeExtraField string `json:"someExtraField"`
}
