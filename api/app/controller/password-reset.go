package controller

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/aicacia/auth/api/app/config"
	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// PostRequestPasswordReset
//
//	@Summary		Request Password Reset
//	@ID				request-password-reset
//	@Tags			password-reset
//	@Accept			json
//	@Produce		json
//	@Param			requestPasswordReset	body	model.RequestPasswordResetST	true	"request password reset body"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/password-reset/request [post]
func PostRequestPasswordReset(c *fiber.Ctx) error {
	var requestPasswordReset model.RequestPasswordResetST
	if err := c.BodyParser(&requestPasswordReset); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	email := strings.TrimSpace(requestPasswordReset.Email)
	phoneNumber := util.NumericRegex.ReplaceAllString(requestPasswordReset.PhoneNumber, "")
	application := middleware.GetApplication(c)
	var user *repository.UserRowST
	if email != "" {
		var err error
		user, err = repository.GetUserByEmail(application.Id, email)
		if err != nil {
			slog.Error("error fetching user by email", "error", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
		}
	}
	if user == nil && phoneNumber != "" {
		var err error
		user, err = repository.GetUserByPhoneNumber(application.Id, phoneNumber)
		if err != nil {
			slog.Error("error fetching user by email", "error", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
		}
	}
	if user == nil {
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	tenent := middleware.GetTenent(c)
	now := time.Now().UTC()
	claims := jwt.Claims{
		Type:             jwt.PasswordResetTokenType,
		Subject:          user.Id,
		NotBeforeSeconds: now.Unix(),
		IssuedAtSeconds:  now.Unix(),
		ExpiresAtSeconds: now.Unix() + int64(tenent.ExpiresInSeconds),
		Issuer:           config.Get().URL,
		Scope:            []string{},
	}
	passwordResetToken, err := jwt.CreateToken(&claims, tenent)
	if err != nil {
		slog.Error("failed to create access token", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send password reset email/phone number
	slog.Info("password reset token", "token", passwordResetToken)
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PostPasswordReset
//
//	@Summary		Request Password Reset
//	@ID				password-reset
//	@Tags			password-reset
//	@Accept			json
//	@Produce		json
//	@Param			passwordReset	body	model.PasswordResetST	true	"request password reset body"
//	@Success		200	{object}	model.TokenST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		404	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/password-reset [post]
func PostPasswordReset(c *fiber.Ctx) error {
	var passwordReset model.PasswordResetST
	if err := c.BodyParser(&passwordReset); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	password := strings.TrimSpace(passwordReset.Password)
	passwordConfirmation := strings.TrimSpace(passwordReset.PasswordConfirmation)
	errors := model.NewError(http.StatusBadRequest)
	if password == "" {
		errors.AddError("password", "required")
	}
	if password != passwordConfirmation {
		errors.AddError("passwordConfirmation", "mismatch")
	}
	if errors.HasErrors() {
		return errors
	}
	tenent := middleware.GetTenent(c)
	claims, err := jwt.ParseClaimsFromToken[jwt.Claims](passwordReset.Token, tenent)
	if err != nil {
		slog.Error("invalid password reset token", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	if claims.Type != jwt.PasswordResetTokenType {
		slog.Error("invalid password reset token", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.UpdateUserPassword(application.Id, claims.Subject, passwordReset.Password)
	if err != nil {
		slog.Error("error setting user reset password token", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return sendToken(c, sendTokenST{
		issuedTokenType: jwt.PasswordResetTokenType,
		scope:           "openid",
		application:     application,
		tenent:          tenent,
		user:            user,
	})
}
