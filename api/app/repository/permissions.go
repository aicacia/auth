package repository

import (
	"time"
)

type PermissionRowST struct {
	Id            int32     `db:"id"`
	ApplicationId int32     `db:"application_id"`
	Description   string    `db:"description"`
	Uri           string    `db:"uri"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatedAt     time.Time `db:"created_at"`
}

func GetPermissions(applicationId int32) ([]PermissionRowST, error) {
	return All[PermissionRowST](`SELECT p.* 
		FROM permissions p
		WHERE p.application_id = $1 
		ORDER BY p.updated_at DESC;`, applicationId)
}

func GetPermissionById(id int32) (*PermissionRowST, error) {
	return GetOptional[PermissionRowST](`SELECT p.*
		FROM permissions p
		WHERE p.id = $1;`, id)
}

func GetUserPermissions(userId, applicationId int32) ([]PermissionRowST, error) {
	return All[PermissionRowST](`SELECT p.* 
		FROM permissions p
		JOIN user_permissions up ON up.permission_id = p.id
		WHERE up.user_id = $1 AND p.application_id = $2
		ORDER BY p.updated_at DESC;`,
		userId, applicationId)
}

func GetServiceAccountPermissions(servieAccountId, applicationId int32) ([]PermissionRowST, error) {
	return All[PermissionRowST](`SELECT p.* 
		FROM permissions p
		JOIN service_account_permissions sap ON sap.permission_id = p.id
		WHERE sap.service_account_id = $1 AND p.application_id = $2
		ORDER BY p.updated_at DESC;`,
		servieAccountId, applicationId)
}

type CreatePermissionST struct {
	Description string `json:"description" validate:"required"`
	Uri         string `json:"uri" validate:"required"`
}

func CreatePermission(applicationId int32, create CreatePermissionST) (PermissionRowST, error) {
	return Get[PermissionRowST](`INSERT INTO permissions 
		(applicationId, description, uri)
		VALUES ($1, $2, $3) 
		RETURNING *;`,
		applicationId, create.Description, create.Uri,
	)
}

type UpdatePermissionST struct {
	Description *string `json:"description"`
	Uri         *string `json:"uri"`
}

func UpdatePermission(id int32, update UpdatePermissionST) (*PermissionRowST, error) {
	return GetOptional[PermissionRowST](`UPDATE permissions SET
		description = COALESCE($2, description),
		uri = COALESCE($3, uri)
		WHERE id = $1
		RETURNING *;`,
		id, update.Description, update.Uri,
	)
}

func DeletePermission(id int32) (bool, error) {
	return Execute(`DELETE FROM permissions WHERE id = $1;`, id)
}

func AddPermissionToUser(userId, applicationPermissionId int32) (bool, error) {
	return Execute(`INSERT INTO user_permissions
		(user_id, permission_id)
		VALUES ($1, $2);`, userId, applicationPermissionId)
}

func RemovePermissionFromUser(userId, applicationPermissionId int32) (bool, error) {
	return Execute(`DELETE FROM user_permissions
		WHERE user_id = $1 AND permission_id = $2;`, userId, applicationPermissionId)
}

func PermissionsToMap(permissions []PermissionRowST) map[string]bool {
	m := make(map[string]bool, len(permissions))
	for _, permission := range permissions {
		m[permission.Uri] = true
	}
	return m
}
