package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
)

type ApplicationST struct {
	Id          int32     `json:"id" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Uri         string    `json:"uri" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name Application

func ApplicationFromApplicationRow(row repository.ApplicationRowST) ApplicationST {
	return ApplicationST{
		Id:          row.Id,
		Description: row.Description,
		Uri:         row.Uri,
		UpdatedAt:   row.UpdatedAt,
		CreatedAt:   row.CreatedAt,
	}
}

type CreateApplicationST struct {
	repository.CreateApplicationST
} // @name CreateApplication

type UpdateApplicationST struct {
	repository.UpdateApplicationST
} // @name UpdateApplication
