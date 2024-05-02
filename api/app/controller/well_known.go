package controller

import (
	"github.com/aicacia/auth/api/app/config"
	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

// GetOpenIDConfiguration
//
//	@Summary		Get openid configuration
//	@ID				openid-configuration
//	@Tags			well-known
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}   	model.OpenIDConfigurationST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/.well-known/openid-configuration [get]
//
//	@Security		TenentId
func GetOpenIDConfiguration(c *fiber.Ctx) error {
	tenent := middleware.GetTenent(c)
	url := config.Get().URL
	grantTypesSupported := []string{"refresh_token"}
	if tenent.RegistrationWebsite != nil {
		grantTypesSupported = append(grantTypesSupported, "password")
	}
	return c.JSON(model.OpenIDConfigurationST{
		Issuer:                url,
		JwksUri:               url + "/.well-known/jwks.json",
		RegistrationEndpoint:  tenent.RegistrationWebsite,
		AuthorizationEndpoint: tenent.AuthorizationWebsite,
		TokenEndpoint:         url + "/token",
		UserInfoEndpoint:      url + "/userinfo",
		ScopesSupported: []string{
			"openid",
		},
		GrantTypesSupported: grantTypesSupported,
		ResponseTypesSupported: []string{
			"id_token",
			"access_token",
			"refresh_token",
		},
		SubjectTypesSupported: []string{
			jwt.UserSubject,
			jwt.ServiceAccountSubject,
		},
		IdTokenSigningAlgValuesSupported: []string{
			tenent.Algorithm,
		},
		TokenEndpointAuthMethodsSupported: []string{
			"client_secret_basic",
			"client_secret_post",
			"client_secret_jwt",
		},
		ClaimsSupported: []string{
			"sub",
			"type",
			"client_id",
			"aud",
			"nbf",
			"iat",
			"iss",
			"exp",
			"scope",
			"email",
			"email_verified",
			"phone",
			"phone_verified",
			"name",
			"given_name",
			"family_name",
			"middle_name",
			"nickname",
			"profile",
			"picture",
			"website",
			"gender",
			"birthdate",
			"zoneinfo",
			"locale",
			"address",
		},
		CodeChallengeMethodsSupported: []string{
			"plain",
			"S256",
		},
	})
}
