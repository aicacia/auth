package repository

import (
	"time"

	"github.com/xlzd/gotp"
)

type TOTPRowST struct {
	Id        int32     `db:"id"`
	TenentId  int32     `db:"tenent_id"`
	UserId    int32     `db:"user_id"`
	Enabled   bool      `db:"enabled"`
	Secret    string    `db:"secret"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func GetTOTPsByUserId(userId int32) ([]TOTPRowST, error) {
	return All[TOTPRowST](
		`SELECT utt.*, (um.id IS NOT NULL) as enabled
		FROM totps utt
		LEFT JOIN user_mfas um ON um.type = 'totp' AND um.id = utt.id
		WHERE utt.user_id = $1;`,
		userId)
}

func GetTOTPsByUserIdAndTenentId(userId, tenentId int32) (*TOTPRowST, error) {
	return GetOptional[TOTPRowST](`SELECT utt.*, (um.id IS NOT NULL) as enabled
		FROM totps utt
		LEFT JOIN user_mfas um ON um.type = 'totp' AND um.id = utt.id
		WHERE utt.user_id = $1 AND utt.tenent_id = $2
		LIMIT 1;`,
		userId, tenentId)
}

func CreateTOTP(userId, tenentId int32) (TOTPRowST, error) {
	return Get[TOTPRowST](`INSERT INTO totps (tenent_id, user_id, secret)
		VALUES ($1, $2, $3)
		RETURNING *;`,
		userId, tenentId, gotp.RandomSecret(16))
}

func DeleteTOTP(userId, tenentId int32) (bool, error) {
	return Execute(`DELETE FROM totps WHERE user_id=$1 AND tenent_id=$2;`,
		userId, tenentId)
}
