package model

type RequestPasswordResetST struct {
	Email       string `json:"email" validate:"required,omitempty"`
	PhoneNumber string `json:"phoneNumber" validate:"required,omitempty"`
}

type PasswordResetST struct {
	Token                string `json:"token" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
}
