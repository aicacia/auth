package model

import (
	"time"

	"github.com/aicacia/auth/api/app/repository"
)

type UserInfoAddressST struct {
	StreetAddress *string `json:"street_address,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	Country       *string `json:"country,omitempty"`
} // @name UserInfoAddress

type UpdateUserInfoRequestST struct {
	Name       *string            `json:"name"`
	GivenName  *string            `json:"given_name"`
	FamilyName *string            `json:"family_name"`
	MiddleName *string            `json:"middle_name"`
	Nickname   *string            `json:"nickname"`
	Profile    *string            `json:"profile"`
	Picture    *string            `json:"picture"`
	Website    *string            `json:"website"`
	Gender     *string            `json:"gender"`
	Birthdate  *time.Time         `json:"birthdate"`
	Zoneinfo   *string            `json:"zoneinfo"`
	Locale     *string            `json:"locale"`
	Address    *UserInfoAddressST `json:"address"`
} // @name UpdateUserInfoRequest

type UserInfoST struct {
	UserId            int32             `json:"user_id" validate:"required"`
	PreferredUsername string            `json:"preferred_username" validate:"required"`
	Name              *string           `json:"name,omitempty"`
	GivenName         *string           `json:"given_name,omitempty"`
	FamilyName        *string           `json:"family_name,omitempty"`
	MiddleName        *string           `json:"middle_name,omitempty"`
	Nickname          *string           `json:"nickname,omitempty"`
	Profile           *string           `json:"profile,omitempty"`
	Picture           *string           `json:"picture,omitempty"`
	Website           *string           `json:"website,omitempty"`
	Gender            *string           `json:"gender,omitempty"`
	Birthdate         *time.Time        `json:"birthdate,omitempty" format:"date-time"`
	Zoneinfo          *string           `json:"zoneinfo,omitempty"`
	Locale            *string           `json:"locale,omitempty"`
	Address           UserInfoAddressST `json:"address" validate:"required"`
	UpdatedAt         time.Time         `json:"updated_at" validate:"required" format:"date-time"`
	CreatedAt         time.Time         `json:"created_at" validate:"required" format:"date-time"`
} // @name UserInfo

func UserInfoFromRow(user *repository.UserRowST, row *repository.UserInfoRowST) UserInfoST {
	return UserInfoST{
		UserId:            user.Id,
		PreferredUsername: user.Username,
		Name:              row.Name,
		GivenName:         row.GivenName,
		FamilyName:        row.FamilyName,
		MiddleName:        row.MiddleName,
		Nickname:          row.Nickname,
		Profile:           row.Profile,
		Picture:           row.Picture,
		Website:           row.Website,
		Gender:            row.Gender,
		Birthdate:         row.Birthdate,
		Zoneinfo:          row.Zoneinfo,
		Locale:            row.Locale,
		Address: UserInfoAddressST{
			StreetAddress: row.StreetAddress,
			Locality:      row.Locality,
			Region:        row.Region,
			PostalCode:    row.PostalCode,
			Country:       row.Country,
		},
		UpdatedAt: row.UpdatedAt,
		CreatedAt: row.CreatedAt,
	}
}
