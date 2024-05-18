package model

import (
	"github.com/aicacia/auth/api/app/repository"
)

type PermissionST struct {
	Resource string   `json:"resource" validate:"required"`
	Actions  []string `json:"actions" validate:"required"`
} // @name Permission

func PermissionFromRow(row repository.PermissionRowST) PermissionST {
	return PermissionST{
		Resource: row.Resource,
		Actions:  row.Actions,
	}
}
