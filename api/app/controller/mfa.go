package controller

import (
	"strings"

	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/middleware"
	_ "github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

// PostValidateMFA
//
//	@Summary		Multi-factor authentication
//	@Description	Multi-factor authentication
//	@ID				validate-mfa
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			mfa	body	model.ValidateMFAST	true	"mfa validation"
//	@Success		200	{object}	model.TokenST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/mfa [post]
//
//	@Security		Authorization
func PostValidateMFA(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	tenent := middleware.GetTenent(c)
	claims := middleware.GetClaims[jwt.MFAClaims](c)
	return sendToken(c, sendTokenST{
		issuedTokenType: claims.GrantType,
		scope:           strings.Join(claims.Scope, " "),
		application:     middleware.GetApplication(c),
		tenent:          tenent,
		user:            user,
	})
}
