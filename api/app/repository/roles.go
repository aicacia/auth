package repository

import (
	"time"
)

type RoleRowST struct {
	Id            int32     `db:"id"`
	ApplicationId int32     `db:"application_id"`
	Description   string    `db:"description"`
	URI           string    `db:"uri"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatedAt     time.Time `db:"created_at"`
}

func GetRoles(applicationId int32) ([]RoleRowST, error) {
	return All[RoleRowST](`SELECT r.*
		FROM roles r
		WHERE r.application_id = $1 
		ORDER BY r.updated_at DESC;`, applicationId)
}

func GetRoleById(id int32) ([]RoleRowST, error) {
	return All[RoleRowST](`SELECT r.*
		FROM roles r
		WHERE r.id = $1 
		LIMIT 1;`, id)
}

type CreateRoleST struct {
	ApplicationId int32  `db:"application_id"`
	Description   string `db:"description"`
	URI           string `db:"uri"`
}

func CreateRole(params CreateRoleST) (RoleRowST, error) {
	return NamedGet[RoleRowST](`INSERT INTO roles (application_id, description, uri, actions)
		VALUES (:application_id, :description, :uri, :actions)
		RETURNING *;`, params)
}

type UpdateRoleST struct {
	Id          int32     `db:"id"`
	Description *string   `db:"description"`
	URI         *string   `db:"uri"`
	Actions     *[]string `db:"actions"`
}

func UpdateRole(params UpdateRoleST) (*RoleRowST, error) {
	return NamedGetOptional[RoleRowST](`UPDATE resources SET
		description=COALESCE(:description, description)
		uri=COALESCE(:uri, uri)
		actions=COALESCE(:id, actions)
	WHERE id=:id
	RETURNING *;`, params)
}
