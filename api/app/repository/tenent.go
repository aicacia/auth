package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TenentRowST struct {
	Id                      int32          `db:"id"`
	ApplicationId           int32          `db:"application_id"`
	Description             string         `db:"description"`
	Uri                     string         `db:"uri"`
	PublicUri               sql.NullString `db:"public_uri"`
	Algorithm               string         `db:"algorithm"`
	ClientId                uuid.UUID      `db:"client_id"`
	ClientSecret            string         `db:"client_secret"`
	PublicKey               sql.NullString `db:"public_key"`
	PrivateKey              string         `db:"private_key"`
	ExpiresInSeconds        int64          `db:"expires_in_seconds"`
	RefreshExpiresInSeconds int64          `db:"refresh_expires_in_seconds"`
	ResetExpiresInSeconds   int64          `db:"reset_expires_in_seconds"`
	UpdatedAt               time.Time      `db:"updated_at"`
	CreatedAt               time.Time      `db:"created_at"`
}

func GetTenents(applicationId int32, limit, offset int) ([]TenentRowST, error) {
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
	Description             string     `json:"description" validate:"required"`
	Uri                     string     `json:"uri" validate:"required"`
	ClientId                *uuid.UUID `json:"client_id"`
	Algorithm               *string    `json:"algorithm"`
	PublicKey               *string    `json:"public_key"`
	PrivateKey              *string    `json:"private_key"`
	ExpiresInSeconds        *int64     `json:"expires_in_seconds"`
	RefreshExpiresInSeconds *int64     `json:"refresh_expires_in_seconds"`
	ResetExpiresInSeconds   *int64     `json:"reset_expires_in_seconds"`
}

func CreateTenent(applicationId int32, create CreateTenentST) (TenentRowST, error) {
	applicationTenent, err := Get[TenentRowST](`INSERT INTO tenents 
		(application_id, description, uri)
		VALUES ($1, $2, $3) 
		RETURNING *;`,
		applicationId, create.Description, create.Uri,
	)
	if err != nil {
		return applicationTenent, err
	}
	updatedTenentApplication, err := UpdateTenent(applicationTenent.Id, UpdateTenentST{
		ClientId:                create.ClientId,
		Algorithm:               create.Algorithm,
		PublicKey:               create.PublicKey,
		PrivateKey:              create.PrivateKey,
		ExpiresInSeconds:        create.ExpiresInSeconds,
		RefreshExpiresInSeconds: create.RefreshExpiresInSeconds,
		ResetExpiresInSeconds:   create.ResetExpiresInSeconds,
	})
	if updatedTenentApplication == nil {
		return applicationTenent, err
	}
	return *updatedTenentApplication, err
}

type UpdateTenentST struct {
	Description             *string    `json:"description" validate:"required"`
	Uri                     *string    `json:"uri" validate:"required"`
	ClientId                *uuid.UUID `json:"client_id"`
	Algorithm               *string    `json:"algorithm"`
	PublicKey               *string    `json:"public_key"`
	PrivateKey              *string    `json:"private_key"`
	ExpiresInSeconds        *int64     `json:"expires_in_seconds"`
	RefreshExpiresInSeconds *int64     `json:"refresh_expires_in_seconds"`
	ResetExpiresInSeconds   *int64     `json:"reset_expires_in_seconds"`
}

func UpdateTenent(id int32, update UpdateTenentST) (*TenentRowST, error) {
	return GetOptional[TenentRowST](`UPDATE tenents SET
		description = COALESCE($2, description),
		uri = COALESCE($3, uri),
		client_id = COALESCE($4, client_id),
		algorithm = COALESCE($5, algorithm),
		public_key = COALESCE($6, public_key),
		private_key = COALESCE($7, private_key),
		expires_in_seconds = COALESCE($8, expires_in_seconds),
		refresh_expires_in_seconds = COALESCE($9, refresh_expires_in_seconds),
		reset_expires_in_seconds = COALESCE($10, reset_expires_in_seconds)
		WHERE id = $1
		RETURNING *;`,
		id, update.Description, update.Uri, update.ClientId, update.Algorithm, update.PublicKey, update.PrivateKey, update.ExpiresInSeconds, update.RefreshExpiresInSeconds,
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
