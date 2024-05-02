package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
	"github.com/google/uuid"
)

type TenentST struct {
	Id                            int32     `json:"id" validate:"required"`
	ApplicationId                 int32     `json:"application_id" validate:"required"`
	Description                   string    `json:"description" validate:"required"`
	Uri                           string    `json:"uri" validate:"required"`
	AuthorizationWebsite          string    `json:"authorization_website" validate:"required"`
	RegistrationWebsite           *string   `json:"registration_website"`
	ClientId                      uuid.UUID `json:"client_id" validate:"required"`
	Algorithm                     string    `json:"algorithm" validate:"required"`
	PublicKey                     *string   `json:"public_key"`
	ExpiresInSeconds              int64     `json:"expires_in_seconds" validate:"required"`
	RefreshExpiresInSeconds       int64     `json:"refresh_expires_in_seconds" validate:"required"`
	PasswordResetExpiresInSeconds int64     `json:"password_reset_expires_in_seconds" validate:"required"`
	UpdatedAt                     time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt                     time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name Tenent

type TenentWithSecretsST struct {
	TenentST
	ClientSecret string `json:"client_secret"`
	PrivateKey   string `json:"private_key" validate:"required"`
} // @name TenentWithSecrets

func TenentFromRow(row repository.TenentRowST) TenentST {
	return TenentST{
		Id:                            row.Id,
		ApplicationId:                 row.ApplicationId,
		Description:                   row.Description,
		Uri:                           row.Uri,
		AuthorizationWebsite:          row.AuthorizationWebsite,
		RegistrationWebsite:           row.RegistrationWebsite,
		Algorithm:                     row.Algorithm,
		ClientId:                      row.ClientId,
		PublicKey:                     row.PublicKey,
		ExpiresInSeconds:              row.ExpiresInSeconds,
		RefreshExpiresInSeconds:       row.RefreshExpiresInSeconds,
		PasswordResetExpiresInSeconds: row.PasswordResetExpiresInSeconds,
		UpdatedAt:                     row.UpdatedAt,
		CreatedAt:                     row.CreatedAt,
	}
}

func TenentWithSecretsFromRow(row repository.TenentRowST) TenentWithSecretsST {
	return TenentWithSecretsST{
		TenentST:     TenentFromRow(row),
		ClientSecret: row.ClientSecret,
		PrivateKey:   row.PrivateKey,
	}
}

type CreateTenentST struct {
	repository.CreateTenentST
} // @name CreateTenent

type UpdateTenentST struct {
	repository.UpdateTenentST
} // @name UpdateTenent
