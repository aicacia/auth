package repository

import (
	"time"

	"github.com/xlzd/gotp"
)

type TOTPRowST struct {
	TenentId  int32     `db:"tenent_id"`
	UserId    int32     `db:"user_id"`
	Secret    string    `db:"secret"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func GetTOTPsByUserId(userId int32) ([]TOTPRowST, error) {
	return All[TOTPRowST](
		`SELECT utt.*
		FROM user_tenent_totps utt
		WHERE utt.user_id = $1;`,
		userId)
}

func GetTOTPsByUserIdAndTenentId(userId, tenentId int32) (*TOTPRowST, error) {
	return GetOptional[TOTPRowST](`SELECT utt.*
		FROM user_tenent_totps utt
		WHERE utt.user_id = $1 AND utt.tenent_id = $2
		LIMIT 1;`,
		userId, tenentId)
}

func CreateTOTP(userId, tenentId int32) (TOTPRowST, error) {
	return Get[TOTPRowST](`INSERT INTO user_tenent_totps (tenent_id, user_id, secret)
		VALUES ($1, $2, $3)
		RETURNING *;`,
		userId, tenentId, gotp.RandomSecret(16))
}
