package repository

import (
	"time"
)

type ApplicationRowST struct {
	Id          int32     `db:"id"`
	Description string    `db:"description"`
	Uri         string    `db:"uri"`
	IsAdmin     bool      `db:"is_admin"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}

func GetApplications(limit, offset int) ([]ApplicationRowST, error) {
	return All[ApplicationRowST](`SELECT a.* 
		FROM applications a 
		ORDER BY a.updated_at DESC 
		LIMIT $1 OFFSET $2;`, limit, offset)
}

func GetApplicationById(id int32) (*ApplicationRowST, error) {
	return GetOptional[ApplicationRowST]("SELECT a.* FROM applications a WHERE a.id=$1 LIMIT 1;", id)
}

func GetApplicationUsers(id int32) ([]UserRowST, error) {
	return All[UserRowST](`SELECT u.* 
		FROM application_users au 
		JOIN users u ON u.id=au.user_id
		WHERE au.application_id=$1;`,
		id)
}

func GetUserApplications(userId int32, limit, offset int) ([]ApplicationRowST, error) {
	return All[ApplicationRowST](`SELECT a.* 
		FROM applications a 
		JOIN application_users au ON au.application_id=a.id
		JOIN users u ON u.id=au.user_id
		WHERE u.id=$1
		LIMIT $2 OFFSET $3;`, userId, limit, offset)
}

type CreateApplicationST struct {
	Description string `json:"description" validate:"required"`
	Uri         string `json:"uri" validate:"required"`
}

func CreateApplication(create CreateApplicationST) (ApplicationRowST, error) {
	return Get[ApplicationRowST](`INSERT INTO applications 
		(description, uri) 
		VALUES 
		($1, $2) 
		RETURNING *;`,
		create.Description, create.Uri)
}

type UpdateApplicationST struct {
	Description *string `json:"description"`
	Uri         *string `json:"uri"`
}

func UpdateApplication(id int32, update UpdateApplicationST) (*ApplicationRowST, error) {
	return GetOptional[ApplicationRowST](`UPDATE applications
		SET description=COALESCE($2, description), 
			uri=COALESCE($3, uri)
		WHERE id=$1
		RETURNING *;`,
		id, update.Description, update.Uri)
}

func DeleteApplication(id int32) (bool, error) {
	return Execute(`DELETE FROM applications WHERE id=$1;`, id)
}

func AddUserToApplication(applicationId, userId int32) (bool, error) {
	return Execute(`INSERT INTO application_users (application_id, user_id) VALUES ($1, $2);`, applicationId, userId)
}

func RemoveUserFromApplication(applicationId, userId int32) (bool, error) {
	return Execute(`DELETE FROM application_users WHERE application_id=$1 AND user_id=$2;`, applicationId, userId)
}
