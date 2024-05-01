package repository

import (
	"time"

	"github.com/aicacia/auth/api/app/util"
	"github.com/google/uuid"
)

type ServiceAccountRowST struct {
	Id              int32     `db:"id"`
	Name            string    `db:"name"`
	Key             uuid.UUID `db:"key"`
	EncryptedSecret string    `db:"encrypted_secret"`
	UpdatedAt       time.Time `db:"updated_at"`
	CreatedAt       time.Time `db:"created_at"`
}

func GetServiceAccounts(limit, offset int) ([]ServiceAccountRowST, error) {
	return All[ServiceAccountRowST](`SELECT sa.* 
		FROM service_accounts sa 
		ORDER BY sa.updated_at DESC 
		LIMIT $1 OFFSET $2;`, limit, offset)
}

func GetServiceAccountById(serviceAccountId int32) (*ServiceAccountRowST, error) {
	return GetOptional[ServiceAccountRowST](`SELECT sa.*
		FROM service_accounts sa
		WHERE sa.id = $1
		LIMIT 1;`,
		serviceAccountId)
}

func GetServiceAccountByKey(key uuid.UUID) (*ServiceAccountRowST, error) {
	return GetOptional[ServiceAccountRowST](`SELECT sa.*
		FROM service_accounts sa
		WHERE sa.key = $1
		LIMIT 1;`,
		key)
}

func CreateServiceAccount(name string) (ServiceAccountRowST, string, error) {
	secret, err := util.GenerateRandomHex(32)
	if err != nil {
		return ServiceAccountRowST{}, "", err
	}
	encryptedSecret, err := util.EncryptPassword(secret)
	if err != nil {
		return ServiceAccountRowST{}, "", err
	}
	serviceAccount, err := Get[ServiceAccountRowST](`INSERT INTO service_accounts (name, encrypted_secret)
		VALUES ($1, $2)
		RETURNING *;`,
		name, encryptedSecret)
	if err != nil {
		return ServiceAccountRowST{}, "", err
	}
	return serviceAccount, secret, nil
}

func ResetServiceAccountKey(id int32) (*ServiceAccountRowST, error) {
	return GetOptional[ServiceAccountRowST](`UPDATE service_accounts
		SET key = gen_random_uuid()
		WHERE id = $1
		RETURNING *;`,
		id)
}

func ResetServiceAccountSecret(id int32) (*ServiceAccountRowST, string, error) {
	secret, err := util.GenerateRandomHex(32)
	if err != nil {
		return nil, "", err
	}
	encryptedSecret, err := util.EncryptPassword(secret)
	if err != nil {
		return nil, "", err
	}
	serviceAccount, err := GetOptional[ServiceAccountRowST](`UPDATE service_accounts
		SET encrypted_password = $2
		WHERE id = $1
		RETURNING *;`,
		id, encryptedSecret)
	if err != nil {
		return nil, "", err
	}
	return serviceAccount, secret, nil
}
