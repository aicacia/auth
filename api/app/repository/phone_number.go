package repository

import (
	"database/sql"
	"time"

	"github.com/aicacia/auth/api/app/util"
)

type PhoneNumberRowST struct {
	Id                int32          `db:"id"`
	ApplicationId     int32          `db:"application_id"`
	UserId            int32          `db:"user_id"`
	PhoneNumber       string         `db:"phone_number"`
	Confirmed         bool           `db:"confirmed"`
	ConfirmationToken sql.NullString `db:"confirmation_token"`
	UpdatedAt         time.Time      `db:"updated_at"`
	CreatedAt         time.Time      `db:"created_at"`
}

func GetPhoneNumbersByUserId(userId int32) ([]PhoneNumberRowST, error) {
	return All[PhoneNumberRowST](
		`SELECT p.*
		FROM phone_numbers p
		WHERE p.user_id = $1;`,
		userId)
}

func GetUserPrimaryPhoneNumber(userId int32) (*PhoneNumberRowST, error) {
	return GetOptional[PhoneNumberRowST](`SELECT p.*
		FROM users u
		JOIN phone_numbers p ON p.id = u.phone_number_id
		WHERE u.id = $1
		LIMIT 1;`,
		userId)
}

func CreatePhoneNumber(applicationId, userId int32, phoneNumber string) (PhoneNumberRowST, error) {
	confirmationToken, err := util.GenerateRandomHex(8)
	if err != nil {
		var empty PhoneNumberRowST
		return empty, err
	}
	return Get[PhoneNumberRowST](`INSERT INTO phone_numbers (application_id, user_id, phone_number, confirmation_token)
		VALUES ($1, $2, $3, $4)
		RETURNING *;`,
		applicationId, userId, phoneNumber, confirmationToken)
}

func SetPhoneNumberConfirmation(userId, id int32, token string) (PhoneNumberRowST, error) {
	return Get[PhoneNumberRowST](`UPDATE phone_numbers SET
		confirmation_token=COALESCE($3, confirmation_token)
		WHERE user_id=$1 
			AND id=$2
		RETURNING *;`,
		userId, id, token)
}

func ConfirmPhoneNumber(userId, id int32, token string) (PhoneNumberRowST, error) {
	return Get[PhoneNumberRowST](`UPDATE phone_numbers SET
		confirmed=true
		WHERE user_id=$1 
			AND id=$2
			AND confirmation_token=$3
		RETURNING *;`,
		userId, id, token)
}

func SetPrimaryPhoneNumber(userId, id int32) (bool, error) {
	return Execute(`UPDATE users SET phone_number_id=$2 WHERE id=$1;`, userId, id)
}

func DeletePhoneNumber(userId, id int32) (bool, error) {
	return Execute(`DELETE FROM phone_numbers WHERE user_id=$1 AND id=$2;`,
		userId, id)
}
