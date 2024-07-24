package controller

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/xlzd/gotp"
)

// PostValidateMFA
//
//	@Summary		Multi-factor authentication
//	@Description	Multi-factor authentication
//	@ID				validate-mfa
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			mfa	body	    model.ValidateMFAST	true	"mfa validation"
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
	mfa, err := repository.GetMFA(user.Id)
	if err != nil {
		slog.Error("failed to find MFA", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if mfa == nil || !mfa.Enabled {
		slog.Error("MFA is not enabled")
		return model.NewError(http.StatusForbidden).AddError("mfa", "disabled")
	}
	var body model.ValidateMFAST
	if err := c.BodyParser(&body); err != nil {
		slog.Error("failed to parse body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("invalid", "body")
	}
	tenent := middleware.GetTenent(c)
	switch mfa.Type {
	case "totp":
		{
			totp, err := repository.GetTOTPsByUserIdAndTenentId(user.Id, tenent.Id)
			if err != nil {
				slog.Error("failed to find TOTP", "error", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			if totp == nil {
				slog.Error("TOTP is not enabled")
				return model.NewError(http.StatusForbidden).AddError("mfa", "disabled")
			}
			if gotp.NewDefaultTOTP(totp.Secret).Now() != body.Code {
				slog.Error("failed to validate MFA", "error", err)
				return model.NewError(http.StatusForbidden).AddError("mfa", "invalid")
			}
		}
	default:
		slog.Error("unknown MFA type", "type", mfa.Type)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	claims := middleware.GetClaims[jwt.MFAClaims](c)
	return sendToken(c, sendTokenST{
		issuedTokenType: claims.GrantType,
		scope:           strings.Join(claims.Scope, " "),
		application:     middleware.GetApplication(c),
		tenent:          tenent,
		user:            user,
	})
}
