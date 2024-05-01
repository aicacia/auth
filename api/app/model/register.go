package model

type RegistrationRequestST struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
} // @name RegistrationRequest
