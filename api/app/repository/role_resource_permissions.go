package repository

import (
	"time"

	"github.com/lib/pq"
)

type RoleResourcePermissionRowST struct {
	RoleId     int32          `db:"role_id"`
	ResourceId int32          `db:"resource_id"`
	Actions    pq.StringArray `db:"actions"`
	UpdatedAt  time.Time      `db:"updated_at"`
	CreatedAt  time.Time      `db:"created_at"`
}

func GetRoleResourcePermissions(applicationId int32) ([]RoleResourcePermissionRowST, error) {
	return All[RoleResourcePermissionRowST](`SELECT rrp.*
		FROM role_resource_permissions rrp
		WHERE rrp.role_id = $1 AND rrp.resource_id = $1 
		ORDER BY rrp.updated_at DESC;`, applicationId)
}

func GetRoleResourcePermissionById(id int32) ([]RoleResourcePermissionRowST, error) {
	return All[RoleResourcePermissionRowST](`SELECT rrp.*
		FROM role_resource_permissions rrp
		WHERE rrp.id = $1 
		LIMIT 1;`, id)
}

type CreateRoleResourcePermissionST struct {
	RoleId     int32          `db:"role_id"`
	ResourceId int32          `db:"resource_id"`
	Actions    pq.StringArray `db:"actions"`
}

func CreateRoleResourcePermission(params CreateRoleResourcePermissionST) (RoleResourcePermissionRowST, error) {
	return NamedGet[RoleResourcePermissionRowST](`INSERT INTO role_resource_permissions (role_id, resource_id, actions)
		VALUES (:role_id, :resource_id, :actions)
		RETURNING *;`, params)
}

type UpdateRoleResourcePermissionST struct {
	RoleId     int32          `db:"role_id"`
	ResourceId int32          `db:"resource_id"`
	Actions    pq.StringArray `db:"actions"`
}

func UpdateRoleResourcePermission(params UpdateRoleResourcePermissionST) (*RoleResourcePermissionRowST, error) {
	return NamedGetOptional[RoleResourcePermissionRowST](`UPDATE role_resource_permissions SET
		actions=COALESCE(:actions, actions)
	WHERE role_id=:role_id AND resource_id=:resource_id
	RETURNING *;`, params)
}
