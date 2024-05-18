package repository

import (
	"time"
)

type EmailRowST struct {
	Id                int32     `db:"id"`
	ApplicationId     int32     `db:"application_id"`
	UserId            int32     `db:"user_id"`
	Email             string    `db:"email"`
	Confirmed         bool      `db:"confirmed"`
	ConfirmationToken *string   `db:"confirmation_token"`
	UpdatedAt         time.Time `db:"updated_at"`
	CreatedAt         time.Time `db:"created_at"`
}

func GetEmailsByUserId(userId int32) ([]EmailRowST, error) {
	return All[EmailRowST](`SELECT e.*
		FROM emails e
		WHERE e.user_id = $1;`,
		userId)
}

func GetUserPrimaryEmail(userId int32) (*EmailRowST, error) {
	return GetOptional[EmailRowST](`SELECT e.*
		FROM users u
		JOIN emails e ON e.id = u.email_id
		WHERE u.id = $1
		LIMIT 1;`,
		userId)
}

func CreateEmail(applicationId, userId int32, email, confirmationToken string) (EmailRowST, error) {
	return Get[EmailRowST](`INSERT INTO emails (application_id, user_id, email, confirmation_token)
		VALUES ($1, $2, $3, $4)
		RETURNING *;`,
		applicationId, userId, email, confirmationToken)
}

func SetEmailConfirmation(userId, id int32, confirmationToken string) (EmailRowST, error) {
	return Get[EmailRowST](`UPDATE emails SET
		confirmation_token=COALESCE($3, confirmation_token)
		WHERE user_id=$1 
			AND id=$2
		RETURNING *;`,
		userId, id, confirmationToken)
}

func ConfirmEmail(userId, id int32, confirmationToken string) (EmailRowST, error) {
	return Get[EmailRowST](`UPDATE emails SET
		confirmed=true
		WHERE user_id=$1 
			AND id=$2
			AND confirmation_token=$3
		RETURNING *;`,
		userId, id, confirmationToken)
}

func SetPrimaryEmail(userId, id int32) (bool, error) {
	return Execute(`UPDATE users SET email_id=$2 WHERE id=$1;`, userId, id)
}

func DeleteEmail(userId, id int32) (bool, error) {
	return Execute(`DELETE FROM emails WHERE user_id=$1 AND id=$2;`,
		userId, id)
}
