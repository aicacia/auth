package repository

import "time"

const (
	MFATypeTOTP = "totp"
)

type MFARowST struct {
	UserId    int32     `db:"user_id"`
	Id        int32     `db:"id"`
	Type      string    `db:"type"`
	Enabled   bool      `db:"enabled"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func GetMFA(userId int32) (*MFARowST, error) {
	return GetOptional[MFARowST](`SELECT m.*, (t.id IS NOT NULL) AS enabled
		FROM user_mfas m
		LEFT JOIN totps t ON m.type = 'totp' AND m.id = t.id
		WHERE m.user_id = $1;`,
		userId)
}

func UpsertMFA(userId, id int32, typ string) (bool, error) {
	return Execute(`INSERT INTO user_mfas (user_id, id, type)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO UPDATE
		SET id = $2, type = $3;`,
		userId, id, typ)
}

func DeleteMFA(userId int32) (bool, error) {
	return Execute(`DELETE FROM user_mfas WHERE user_id = $1;`, userId)
}
