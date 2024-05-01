package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
)

type PermissionST struct {
	Id            int32     `json:"id" validate:"required"`
	ApplicationId int32     `json:"application_id" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	Uri           string    `json:"uri" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt     time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name Permission

func PermissionFromPermissionRow(row repository.PermissionRowST) PermissionST {
	return PermissionST{
		Id:            row.Id,
		ApplicationId: row.ApplicationId,
		Description:   row.Description,
		Uri:           row.Uri,
		UpdatedAt:     row.UpdatedAt,
		CreatedAt:     row.CreatedAt,
	}
}

type CreatePermissionST struct {
	repository.CreatePermissionST
} // @name CreatePermission

type UpdatePermissionST struct {
	repository.UpdatePermissionST
} // @name UpdatePermission
