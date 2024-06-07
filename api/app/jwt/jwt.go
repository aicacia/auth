package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/aicacia/auth/api/app/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ToMapClaims interface {
	ToMapClaims() (jwt.MapClaims, error)
}

func anyToMapClaims(value any) (jwt.MapClaims, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var result jwt.MapClaims
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

var (
	UserSubject           = "user"
	ServiceAccountSubject = "service-account"
)

var (
	BearerTokenType        = "bearer"
	RefreshTokenType       = "refresh"
	PasswordResetTokenType = "password-reset"
	MFATokenType           = "mfa"
)

type Claims struct {
	Subject          int32     `json:"sub" validate:"required"`
	SubjectType      string    `json:"sub_type" validate:"required"`
	Type             string    `json:"type" validate:"required"`
	ClientId         uuid.UUID `json:"client_id" validate:"required"`
	Audiences        []string  `json:"aud" validate:"required"`
	NotBeforeSeconds int64     `json:"nbf" validate:"required"`
	IssuedAtSeconds  int64     `json:"iat" validate:"required"`
	Issuer           string    `json:"iss" validate:"required"`
	ExpiresAtSeconds int64     `json:"exp" validate:"required"`
	Scope            []string  `json:"scope" validate:"required"`
}

func (claims *Claims) ToMapClaims() (jwt.MapClaims, error) {
	return anyToMapClaims(claims)
}

func (claims *Claims) ToRefreshClaims(application *repository.ApplicationRowST, tenent *repository.TenentRowST) *Claims {
	claims.ExpiresAtSeconds = claims.IssuedAtSeconds + tenent.RefreshExpiresInSeconds
	claims.Type = RefreshTokenType
	return claims
}

type MFAClaims struct {
	Claims
	GrantType string `json:"grant_type" validate:"required"`
}

func (claims *MFAClaims) ToMapClaims() (jwt.MapClaims, error) {
	return anyToMapClaims(claims)
}

type OpenIdClaimsAddress struct {
	StreetAddress *string `json:"street_address"`
	Locality      *string `json:"locality"`
	Region        *string `json:"region"`
	PostalCode    *string `json:"postal_code"`
	Country       *string `json:"country"`
}

type OpenIdClaims struct {
	Claims
	Email         *string             `json:"email"`
	EmailVerified *bool               `json:"email_verified"`
	Phone         *string             `json:"phone"`
	PhoneVerified *bool               `json:"phone_verified"`
	Name          *string             `json:"name"`
	GivenName     *string             `json:"given_name"`
	FamilyName    *string             `json:"family_name"`
	MiddleName    *string             `json:"middle_name"`
	Nickname      *string             `json:"nickname"`
	Profile       *string             `json:"profile"`
	Picture       *string             `json:"picture"`
	Website       *string             `json:"website"`
	Gender        *string             `json:"gender"`
	Birthdate     *time.Time          `json:"birthdate" format:"date-time"`
	Zoneinfo      *string             `json:"zoneinfo"`
	Locale        *string             `json:"locale"`
	Address       OpenIdClaimsAddress `json:"address"`
}

func (claims *OpenIdClaims) ToMapClaims() (jwt.MapClaims, error) {
	return anyToMapClaims(claims)
}

func OpenIdClaimsForUser(claims *Claims, userId int32) (*OpenIdClaims, error) {
	userInfoRow, err := repository.GetUserInfoByUserId(userId)
	if err != nil {
		return nil, err
	}
	emailRow, err := repository.GetUserPrimaryEmail(userId)
	if err != nil {
		return nil, err
	}
	var email *string
	var emailVerified *bool
	if emailRow != nil {
		email = &emailRow.Email
		emailVerified = &emailRow.Confirmed
	}
	phoneRow, err := repository.GetUserPrimaryPhoneNumber(userId)
	if err != nil {
		return nil, err
	}
	var phone *string
	var phoneVerified *bool
	if phoneRow != nil {
		phone = &phoneRow.PhoneNumber
		phoneVerified = &phoneRow.Confirmed
	}
	return &OpenIdClaims{
		Claims:        *claims,
		Email:         email,
		EmailVerified: emailVerified,
		Phone:         phone,
		PhoneVerified: phoneVerified,
		Name:          userInfoRow.Name,
		GivenName:     userInfoRow.GivenName,
		FamilyName:    userInfoRow.FamilyName,
		MiddleName:    userInfoRow.MiddleName,
		Nickname:      userInfoRow.Nickname,
		Profile:       userInfoRow.Profile,
		Picture:       userInfoRow.Picture,
		Website:       userInfoRow.Website,
		Gender:        userInfoRow.Gender,
		Birthdate:     userInfoRow.Birthdate,
		Zoneinfo:      userInfoRow.Zoneinfo,
		Locale:        userInfoRow.Locale,
		Address: OpenIdClaimsAddress{
			StreetAddress: userInfoRow.StreetAddress,
			Locality:      userInfoRow.Locality,
			Region:        userInfoRow.Region,
			PostalCode:    userInfoRow.PostalCode,
			Country:       userInfoRow.Country,
		},
	}, nil
}

func ParseScopes(scope string) []string {
	parts := strings.Split(scope, " ")
	scopes := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" && !slices.Contains(scopes, part) {
			part := strings.ToLower(part)
			scopes = append(scopes, part)
		}
	}
	return scopes
}

func ParseClaimsFromToken[C any](tokenString string, tenent *repository.TenentRowST) (*C, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		alg := token.Method.Alg()
		if alg != tenent.Algorithm {
			return nil, fmt.Errorf("invalid algorithm")
		}
		if alg == "HS256" || alg == "HS384" || alg == "HS512" {
			return []byte(tenent.PrivateKey), nil
		} else if tenent.PublicKey != nil {
			return []byte(*tenent.PublicKey), nil
		} else {
			return nil, fmt.Errorf("invalid algorithm")
		}
	})
	if err != nil {
		return nil, err
	}
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		bytes, err := json.Marshal(&mapClaims)
		if err != nil {
			return nil, err
		}
		var claims C
		if err := json.Unmarshal(bytes, &claims); err != nil {
			return nil, err
		}
		return &claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func ParseClaimsFromTokenNoValidation[C Claims](tokenString string) (*C, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token")
	}
	bytes, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims C
	if err := json.Unmarshal(bytes, &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}

func CreateToken[C ToMapClaims](claims C, tenent *repository.TenentRowST) (string, error) {
	mapClaims, err := claims.ToMapClaims()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(tenent.Algorithm), mapClaims)
	tokenString, err := token.SignedString([]byte(tenent.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
