package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
)

type EmailST struct {
	Id            int32     `json:"id" validate:"required"`
	ApplicationId int32     `json:"application_id" validate:"required"`
	Email         string    `json:"email" validate:"required"`
	Confirmed     bool      `json:"confirmed" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt     time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name Email

type PhoneNumberST struct {
	Id            int32     `json:"id" validate:"required"`
	ApplicationId int32     `json:"application_id" validate:"required"`
	PhoneNumber   string    `json:"phone_number" validate:"required"`
	Confirmed     bool      `json:"confirmed" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt     time.Time `json:"created_at" validate:"required" format:"date-time"`
} // @name PhoneNumber

type UserST struct {
	Id            int32           `json:"id" validate:"required"`
	ApplicationId int32           `json:"application_id" validate:"required"`
	Email         *EmailST        `json:"email"`
	Emails        []EmailST       `json:"emails" validate:"required"`
	PhoneNumber   *PhoneNumberST  `json:"phone_number"`
	PhoneNumbers  []PhoneNumberST `json:"phone_numbers" validate:"required"`
	Username      string          `json:"username" validate:"required"`
	UpdatedAt     time.Time       `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt     time.Time       `json:"created_at" validate:"required" format:"date-time"`
} // @name User

type UserWithPermissionsST struct {
	UserST
	Permissions []string `json:"permissions" validate:"required"`
} // @name UserWithPermissions

func UserFromRow(userRow repository.UserRowST, emailRows []repository.EmailRowST, phoneNumberRows []repository.PhoneNumberRowST) UserST {
	var primaryEmail *EmailST
	emails := make([]EmailST, 0, len(emailRows))
	for _, emailRow := range emailRows {
		email := EmailFromRow(emailRow)
		if userRow.EmailId.Valid && userRow.EmailId.Int32 == emailRow.Id {
			primaryEmail = &email
		} else {
			emails = append(emails, email)
		}
	}
	var primaryPhoneNumber *PhoneNumberST
	phoneNumbers := make([]PhoneNumberST, 0, len(phoneNumberRows))
	for _, phoneNumberRow := range phoneNumberRows {
		phoneNumber := PhoneNumberFromRow(phoneNumberRow)
		if userRow.PhoneNumberId.Valid && userRow.PhoneNumberId.Int32 == phoneNumberRow.Id {
			primaryPhoneNumber = &phoneNumber
		} else {
			phoneNumbers = append(phoneNumbers, phoneNumber)
		}
	}
	return UserST{
		Id:            userRow.Id,
		ApplicationId: userRow.ApplicationId,
		Email:         primaryEmail,
		Emails:        emails,
		PhoneNumber:   primaryPhoneNumber,
		PhoneNumbers:  phoneNumbers,
		Username:      userRow.Username,
		UpdatedAt:     userRow.UpdatedAt,
		CreatedAt:     userRow.CreatedAt,
	}
}

func EmailFromRow(row repository.EmailRowST) EmailST {
	return EmailST{
		Id:            row.Id,
		ApplicationId: row.ApplicationId,
		Email:         row.Email,
		Confirmed:     row.Confirmed,
		UpdatedAt:     row.UpdatedAt,
		CreatedAt:     row.CreatedAt,
	}
}

func PhoneNumberFromRow(row repository.PhoneNumberRowST) PhoneNumberST {
	return PhoneNumberST{
		Id:            row.Id,
		ApplicationId: row.ApplicationId,
		PhoneNumber:   row.PhoneNumber,
		Confirmed:     row.Confirmed,
		UpdatedAt:     row.UpdatedAt,
		CreatedAt:     row.CreatedAt,
	}
}

type CreateUserST struct {
	Username string `json:"username" validate:"required"`
} // @name CreateUser

type UpdateUserST struct {
	Username string `json:"username" validate:"required"`
} // @name UpdateUser

type CreateEmailST struct {
	Email string `json:"email" validate:"required"`
} // @name CreateEmail

type ConfirmEmailST struct {
	Token string `json:"token" validate:"required"`
} // @name ConfirmEmail

type CreatePhoneNumberST struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
} // @name CreatePhoneNumber

type ConfirmPhoneNumberST struct {
	Token string `json:"token" validate:"required"`
} // @name ConfirmPhoneNumber

type ResetPasswordST struct {
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
} // @name ResetPassword
