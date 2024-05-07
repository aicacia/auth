package repository

import (
	"database/sql"
	"strings"
	"time"

	"github.com/aicacia/auth/api/app/util"
	"github.com/jmoiron/sqlx"
)

type UserRowST struct {
	Id                int32         `db:"id"`
	ApplicationId     int32         `db:"application_id"`
	EmailId           sql.NullInt32 `db:"email_id"`
	PhoneNumberId     sql.NullInt32 `db:"phone_number_id"`
	Username          string        `db:"username"`
	EncryptedPassword string        `db:"encrypted_password"`
	UpdatedAt         time.Time     `db:"updated_at"`
	CreatedAt         time.Time     `db:"created_at"`
}

func GetUsers(applicationId int32, limit, offset int) ([]UserRowST, error) {
	return All[UserRowST](`SELECT u.* 
		FROM users u 
		WHERE u.application_id = $1
		ORDER BY u.updated_at DESC 
		LIMIT $2 OFFSET $3;`, applicationId, limit, offset)
}

func GetUsersEmails(applicationId int32, limit, offset int) ([]EmailRowST, error) {
	return All[EmailRowST](`SELECT e.* 
		FROM emails e
		WHERE e.user_id in (
			SELECT u.id 
			FROM users u
			WHERE u.application_id = $1
			ORDER BY u.updated_at DESC 
			LIMIT $2 OFFSET $3
		);`, applicationId, limit, offset)
}

func GetUsersPhoneNumbers(applicationId int32, limit, offset int) ([]PhoneNumberRowST, error) {
	return All[PhoneNumberRowST](`SELECT pn.* 
			FROM phone_numbers pn
			WHERE pn.user_id in (
				SELECT u.id 
				FROM users u
				WHERE u.application_id = $1
				ORDER BY u.updated_at DESC 
				LIMIT $2 OFFSET $3
			);`, applicationId, limit, offset)
}

func GetUserById(applicationId, userId int32) (*UserRowST, error) {
	return GetOptional[UserRowST](`SELECT u.*
		FROM users u
		WHERE u.application_id = $1 AND u.id = $2
		LIMIT 1;`,
		applicationId, userId)
}

func GetUserByUsernameOrEmail(applicationId int32, usernameOrEmail string) (*UserRowST, error) {
	return GetOptional[UserRowST](`SELECT u.*
		FROM users u
		LEFT JOIN emails e ON e.id = u.email_id
		WHERE u.application_id = $1 AND u.username = $2 OR e.email = $2
		LIMIT 1;`,
		applicationId, usernameOrEmail)
}

func GetUserByEmail(applicationId int32, email string) (*UserRowST, error) {
	return GetOptional[UserRowST](`SELECT u.*
		FROM users u
		LEFT JOIN emails e ON e.id = u.email_id
		WHERE u.application_id = $1 AND e.email = $2
		LIMIT 1;`,
		applicationId, email)
}

func GetUserByPhoneNumber(applicationId int32, phoneNumber string) (*UserRowST, error) {
	return GetOptional[UserRowST](`SELECT u.*
		FROM users u
		LEFT JOIN phone_numbers e ON e.id = u.phone_number_id
		WHERE u.application_id = $1 AND e.phone_number = $2
		LIMIT 1;`,
		applicationId, phoneNumber)
}

func GetUserByUsername(applicationId int32, username string) (*UserRowST, error) {
	return GetOptional[UserRowST](`SELECT u.*
		FROM users u
		WHERE u.application_id = $1 AND u.username = $2
		LIMIT 1;`,
		applicationId, username)
}

type UserAndUserInfoST struct {
	User     UserRowST
	UserInfo UserInfoRowST
}

func CreateUserWithPassword(applicationId int32, username, password string) (UserAndUserInfoST, error) {
	return Transaction(func(tx *sqlx.Tx) (UserAndUserInfoST, error) {
		var result UserAndUserInfoST
		encryptedPassword, err := util.EncryptPassword(password)
		if err != nil {
			return result, err
		}
		userRow := tx.QueryRowx(`INSERT INTO users (application_id, username, encrypted_password)
			VALUES ($1, $2, $3)
			RETURNING *;`,
			applicationId, username, encryptedPassword)
		if userRow.Err() != nil {
			return result, userRow.Err()
		}
		err = userRow.StructScan(&result.User)
		if err != nil {
			return result, err
		}
		userInfoRow := tx.QueryRowx(`INSERT INTO user_infos (application_id, user_id) VALUES ($1, $2) RETURNING *;`, applicationId, result.User.Id)
		err = userInfoRow.StructScan(&result.UserInfo)
		if err != nil {
			return result, err
		}
		return result, nil
	})
}

func CreateUserFromUsername(applicationId int32, username string) (UserAndUserInfoST, error) {
	return Transaction(func(tx *sqlx.Tx) (UserAndUserInfoST, error) {
		var result UserAndUserInfoST
		password, err := util.GenerateRandomHex(32)
		if err != nil {
			return result, err
		}
		encryptedPassword, err := util.EncryptPassword(password)
		if err != nil {
			return result, err
		}
		userRow := tx.QueryRowx(`INSERT INTO users (application_id, username, encrypted_password)
			VALUES ($1, $2, $3)
			RETURNING *;`,
			applicationId, username, encryptedPassword)
		if userRow.Err() != nil {
			return result, userRow.Err()
		}
		err = userRow.StructScan(&result.User)
		if err != nil {
			return result, err
		}
		userInfoRow := tx.QueryRowx(`INSERT INTO user_infos (application_id, user_id) VALUES ($1, $2) RETURNING *;`, applicationId, result.User.Id)
		err = userInfoRow.StructScan(&result.UserInfo)
		if err != nil {
			return result, err
		}
		return result, nil
	})
}

func CreateUserFromEmail(applicationId int32, email string) (UserAndUserInfoST, error) {
	return Transaction(func(tx *sqlx.Tx) (UserAndUserInfoST, error) {
		var result UserAndUserInfoST
		user, err := GetUserByEmail(applicationId, email)
		if err != nil {
			return result, err
		}
		if user != nil {
			userInfo, err := GetUserInfoByUserId(user.Id)
			if err != nil {
				return result, err
			}
			result.User = *user
			result.UserInfo = *userInfo
			return result, nil
		}
		username := strings.Split(email, "@")[0]
		for {
			user, err := GetUserByUsername(applicationId, username)
			if err != nil {
				return result, err
			}
			if user == nil {
				break
			}
			hex, err := util.GenerateRandomHex(2)
			if err != nil {
				return result, err
			}
			username += hex
		}
		password, err := util.GenerateRandomHex(32)
		if err != nil {
			return result, err
		}
		encryptedPassword, err := util.EncryptPassword(password)
		if err != nil {
			return result, err
		}
		userRow := tx.QueryRowx(`INSERT INTO users (application_id, username, encrypted_password)
			VALUES ($1, $2, $3)
			RETURNING *;`,
			applicationId, username, encryptedPassword)
		if userRow.Err() != nil {
			return result, userRow.Err()
		}
		err = userRow.StructScan(&result.User)
		if err != nil {
			return result, err
		}
		userInfoRow := tx.QueryRowx(`INSERT INTO user_infos (application_id, user_id) VALUES ($1, $2) RETURNING *;`, applicationId, result.User.Id)
		err = userInfoRow.StructScan(&result.UserInfo)
		if err != nil {
			return result, err
		}
		return result, nil
	})
}

func UpdateUserPassword(applicationId, id int32, password string) (*UserRowST, error) {
	encryptedPassword, err := util.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	return GetOptional[UserRowST](`UPDATE users
		SET encrypted_password = $3
		WHERE application_id=$1 AND id=$2
		RETURNING *;`,
		applicationId, id, encryptedPassword)
}

func UpdateUsername(applicationId, id int32, username string) (*UserRowST, error) {
	return GetOptional[UserRowST](`UPDATE users
		SET username = $3
		WHERE application_id=$1 AND id=$2
		RETURNING *;`,
		applicationId, id, username)
}

func DeleteUserById(applicationId, id int32) (bool, error) {
	return Execute(`DELETE FROM users WHERE application_id=$1 AND id=$2;`, applicationId, id)
}
