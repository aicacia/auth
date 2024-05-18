package repository

import "github.com/lib/pq"

type PermissionRowST struct {
	Resource string         `db:"resource"`
	Actions  pq.StringArray `db:"actions"`
}

func GetUserPermissions(userId int32) ([]PermissionRowST, error) {
	return All[PermissionRowST](`SELECT resources.uri as resource, rrp.actions
		FROM user_roles ur
		JOIN role_resource_permissions rrp ON rrp.role_id = ur.role_id
		JOIN resources ON resources.id = rrp.resource_id
		WHERE ur.user_id = $1
		ORDER BY rrp.updated_at DESC;`, userId)
}

func GetServiceAccountPermissions(userId int32) ([]PermissionRowST, error) {
	return All[PermissionRowST](`SELECT resources.uri as resource, rrp.actions
		FROM service_account_roles sar
		JOIN role_resource_permissions rrp ON rrp.role_id = sar.role_id
		JOIN resources ON resources.id = rrp.resource_id
		WHERE sar.user_id = $1
		ORDER BY rrp.updated_at DESC;`, userId)
}
