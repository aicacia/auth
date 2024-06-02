package repository

import (
	"time"

	"github.com/google/uuid"
)

type TenentRowST struct {
	Id                            int32     `db:"id"`
	ApplicationId                 int32     `db:"application_id"`
	Description                   string    `db:"description"`
	URI                           string    `db:"uri"`
	AuthorizationWebsite          string    `db:"authorization_website"`
	RegistrationWebsite           *string   `db:"registration_website"`
	EmailEndpoint                 *string   `db:"email_endpoint"`
	PhoneNumberEndpoint           *string   `db:"phone_number_endpoint"`
	Algorithm                     string    `db:"algorithm"`
	ClientId                      uuid.UUID `db:"client_id"`
	ClientSecret                  string    `db:"client_secret"`
	PublicKey                     *string   `db:"public_key"`
	PrivateKey                    string    `db:"private_key"`
	ExpiresInSeconds              int64     `db:"expires_in_seconds"`
	RefreshExpiresInSeconds       int64     `db:"refresh_expires_in_seconds"`
	PasswordResetExpiresInSeconds int64     `db:"password_reset_expires_in_seconds"`
	UpdatedAt                     time.Time `db:"updated_at"`
	CreatedAt                     time.Time `db:"created_at"`
}

func GetTenents(applicationId int32, limit, offset *int) ([]TenentRowST, error) {
	if limit == nil && offset == nil {
		return All[TenentRowST](`SELECT at.* 
			FROM tenents at 
			WHERE at.application_id = $1 
			ORDER BY at.updated_at DESC ;`, applicationId)
	}
	if limit == nil {
		limit = new(int)
		*limit = 10
	}
	if offset == nil {
		offset = new(int)
		*offset = 0
	}
	return All[TenentRowST](`SELECT at.* 
		FROM tenents at 
		WHERE at.application_id = $1 
		ORDER BY at.updated_at DESC 
		LIMIT $2 OFFSET $3;`, applicationId, limit, offset)
}

func GetTenentById(id int32) (*TenentRowST, error) {
	return GetOptional[TenentRowST](`SELECT at.* 
		FROM tenents at 
		WHERE at.id = $1
		ORDER BY at.updated_at DESC 
		LIMIT 1;`, id)
}

func GetTenentByClientId(clientId uuid.UUID) (*TenentRowST, error) {
	return GetOptional[TenentRowST](`SELECT at.* 
		FROM tenents at 
		WHERE at.client_id = $1
		ORDER BY at.updated_at DESC 
		LIMIT 1;`, clientId)
}

type CreateTenentST struct {
	Description                   string     `json:"description" validate:"required"`
	URI                           string     `json:"uri" validate:"required"`
	AuthorizationWebsite          string     `json:"authorization_website"`
	RegistrationWebsite           *string    `json:"registration_website"`
	EmailEndpoint                 *string    `json:"email_endpoint"`
	PhoneNumberEndpoint           *string    `json:"phone_number_endpoint"`
	ClientId                      *uuid.UUID `json:"client_id"`
	Algorithm                     *string    `json:"algorithm"`
	PublicKey                     *string    `json:"public_key"`
	PrivateKey                    *string    `json:"private_key"`
	ExpiresInSeconds              *int64     `json:"expires_in_seconds"`
	RefreshExpiresInSeconds       *int64     `json:"refresh_expires_in_seconds"`
	PasswordResetExpiresInSeconds *int64     `json:"password_reset_expires_in_seconds"`
}

func CreateTenent(applicationId int32, create CreateTenentST) (TenentRowST, error) {
	tenent, err := Get[TenentRowST](`INSERT INTO tenents 
		(application_id, description, uri, authorization_website)
		VALUES ($1, $2, $3, $5) 
		RETURNING *;`,
		applicationId, create.Description, create.URI, create.AuthorizationWebsite,
	)
	if err != nil {
		return tenent, err
	}
	updatedTenentApplication, err := UpdateTenent(tenent.Id, UpdateTenentST{
		RegistrationWebsite:           create.RegistrationWebsite,
		EmailEndpoint:                 create.EmailEndpoint,
		PhoneNumberEndpoint:           create.PhoneNumberEndpoint,
		ClientId:                      create.ClientId,
		Algorithm:                     create.Algorithm,
		PublicKey:                     create.PublicKey,
		PrivateKey:                    create.PrivateKey,
		ExpiresInSeconds:              create.ExpiresInSeconds,
		RefreshExpiresInSeconds:       create.RefreshExpiresInSeconds,
		PasswordResetExpiresInSeconds: create.PasswordResetExpiresInSeconds,
	})
	if updatedTenentApplication == nil {
		return tenent, err
	}
	return *updatedTenentApplication, err
}

type UpdateTenentST struct {
	AuthorizationWebsite          *string    `json:"authorization_website"`
	RegistrationWebsite           *string    `json:"registration_website"`
	EmailEndpoint                 *string    `json:"email_endpoint"`
	PhoneNumberEndpoint           *string    `json:"phone_number_endpoint"`
	Description                   *string    `json:"description" validate:"required"`
	URI                           *string    `json:"uri" validate:"required"`
	ClientId                      *uuid.UUID `json:"client_id"`
	Algorithm                     *string    `json:"algorithm"`
	PublicKey                     *string    `json:"public_key"`
	PrivateKey                    *string    `json:"private_key"`
	ExpiresInSeconds              *int64     `json:"expires_in_seconds"`
	RefreshExpiresInSeconds       *int64     `json:"refresh_expires_in_seconds"`
	PasswordResetExpiresInSeconds *int64     `json:"password_reset_expires_in_seconds"`
}

func UpdateTenent(id int32, update UpdateTenentST) (*TenentRowST, error) {
	return GetOptional[TenentRowST](`UPDATE tenents SET
		description = COALESCE($2, description),
		uri = COALESCE($3, uri),
		authorization_website = COALESCE($4, authorization_website),
		registration_website = COALESCE($5, registration_website),
		email_endpoint = COALESCE($6, email_endpoint),
		phone_number_endpoint = COALESCE($7, phone_number_endpoint),
		client_id = COALESCE($8, client_id),
		algorithm = COALESCE($9, algorithm),
		public_key = COALESCE($10, public_key),
		private_key = COALESCE($11, private_key),
		expires_in_seconds = COALESCE($12, expires_in_seconds),
		refresh_expires_in_seconds = COALESCE($13, refresh_expires_in_seconds),
		password_reset_expires_in_seconds = COALESCE($14, password_reset_expires_in_seconds)
		WHERE id = $1
		RETURNING *;`,
		id, update.Description, update.URI, update.AuthorizationWebsite, update.RegistrationWebsite, update.EmailEndpoint, update.PhoneNumberEndpoint, update.ClientId, update.Algorithm, update.PublicKey, update.PrivateKey, update.ExpiresInSeconds, update.RefreshExpiresInSeconds, update.PasswordResetExpiresInSeconds,
	)
}

func RegenerateTenentSecret(id int32) (*TenentRowST, error) {
	return GetOptional[TenentRowST](`UPDATE tenents 
		SET client_secret = encode(gen_random_bytes(32), 'hex')
		WHERE id = $1 
		RETURNING *;`,
		id,
	)
}

func DeleteTenent(id int32) (bool, error) {
	return Execute(`DELETE FROM tenents WHERE id = $1;`, id)
}
