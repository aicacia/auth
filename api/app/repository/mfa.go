package repository

import "time"

const (
	MFATypeTOTP = "totp"
)

type MFARowST struct {
	UserId    int32     `db:"user_id"`
	Id        int32     `db:"id"`
	Type      string    `db:"type"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func GetMFA(userId int32) (*MFARowST, error) {
	return GetOptional[MFARowST](`SELECT m.*
		FROM user_mfas m
		WHERE m.user_id = $1;`,
		userId)
}
