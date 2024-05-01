package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
	"github.com/google/uuid"
)

type ServiceAccountST struct {
	Id        int32     `json:"id" validate:"required"`
	Name      string    `json:"email" validate:"required"`
	Key       uuid.UUID `json:"key" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name ServiceAccount

func ServiceAccountFromServiceAccountRow(row repository.ServiceAccountRowST) ServiceAccountST {
	return ServiceAccountST{
		Id:        row.Id,
		Name:      row.Name,
		Key:       row.Key,
		UpdatedAt: row.UpdatedAt,
		CreatedAt: row.CreatedAt,
	}
}
