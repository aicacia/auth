package repository

import (
	"database/sql"
	"time"
)

type UserInfoRowST struct {
	UserId        int32          `db:"user_id"`
	ApplicationId int32          `db:"application_id"`
	Name          sql.NullString `db:"name"`
	GivenName     sql.NullString `db:"given_name"`
	FamilyName    sql.NullString `db:"family_name"`
	MiddleName    sql.NullString `db:"middle_name"`
	Nickname      sql.NullString `db:"nickname"`
	Profile       sql.NullString `db:"profile"`
	Picture       sql.NullString `db:"picture"`
	Website       sql.NullString `db:"website"`
	Gender        sql.NullString `db:"gender"`
	Birthdate     sql.NullTime   `db:"birthdate"`
	Zoneinfo      sql.NullString `db:"zoneinfo"`
	Locale        sql.NullString `db:"locale"`
	StreetAddress sql.NullString `db:"street_address"`
	Locality      sql.NullString `db:"locality"`
	Region        sql.NullString `db:"region"`
	PostalCode    sql.NullString `db:"postal_code"`
	Country       sql.NullString `db:"country"`
	UpdatedAt     time.Time      `db:"updated_at"`
	CreatedAt     time.Time      `db:"created_at"`
}

func GetUserInfoByUserId(userId int32) (*UserInfoRowST, error) {
	return GetOptional[UserInfoRowST](`SELECT ui.*
		FROM user_infos ui
		WHERE ui.user_id = $1
		LIMIT 1;`,
		userId)
}

type UpdateUserInfoST struct {
	Name          *string    `json:"name"`
	GivenName     *string    `json:"given_name"`
	FamilyName    *string    `json:"family_name"`
	MiddleName    *string    `json:"middle_name"`
	Nickname      *string    `json:"nickname"`
	Profile       *string    `json:"profile"`
	Picture       *string    `json:"picture"`
	Website       *string    `json:"website"`
	Gender        *string    `json:"gender"`
	Birthdate     *time.Time `json:"birthdate"`
	Zoneinfo      *string    `json:"zoneinfo"`
	Locale        *string    `json:"locale"`
	StreetAddress *string    `json:"street_address"`
	Locality      *string    `json:"locality"`
	Region        *string    `json:"region"`
	PostalCode    *string    `json:"postal_code"`
	Country       *string    `json:"country"`
}

func UpdateUserInfoByUserId(userId int32, updates UpdateUserInfoST) (*UserInfoRowST, error) {
	return GetOptional[UserInfoRowST](`UPDATE user_infos
		SET name = COALESCE($2, name),
		  given_name = COALESCE($3, given_name),
		  family_name = COALESCE($4, family_name),
		  middle_name = COALESCE($5, middle_name),
		  nickname = COALESCE($6, nickname),
		  profile = COALESCE($7, profile),
		  picture = COALESCE($8, picture),
		  website = COALESCE($9, website),
		  gender = COALESCE($10, gender),
		  birthdate = COALESCE($11, birthdate),
		  zoneinfo = COALESCE($12, zoneinfo),
		  locale = COALESCE($13, locale),
		  street_address = COALESCE($14, street_address),
		  locality = COALESCE($15, locality),
		  region = COALESCE($16, region),
		  postal_code = COALESCE($17, postal_code),
		  country = COALESCE($18, country)
		WHERE user_id = $1
		RETURNING *;`,
		userId,
		updates.Name,
		updates.GivenName,
		updates.FamilyName,
		updates.MiddleName,
		updates.Nickname,
		updates.Profile,
		updates.Picture,
		updates.Website,
		updates.Gender,
		updates.Birthdate,
		updates.Zoneinfo,
		updates.Locale,
		updates.StreetAddress,
		updates.Locality,
		updates.Region,
		updates.PostalCode,
		updates.Country,
	)
}
