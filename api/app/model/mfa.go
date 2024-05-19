package model

type ValidateMFAST struct {
	Code string `json:"code" validate:"required"`
}
