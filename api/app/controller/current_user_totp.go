package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// GetCurrentUserTOTPs
//
//	@Summary		Get user TOTPs
//	@ID				totps
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Success		201	{array}   	model.TOTPST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp [get]
//
//	@Security		Authorization
func GetCurrentUserTOTPs(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	totps, err := repository.GetTOTPsByUserId(user.Id)
	if err != nil {
		slog.Error("failed to create TOTP", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(util.Map(totps, model.TOTPFromRow))
}

// PostCurrentUserCreateTOTP
//
//	@Summary		Create user TOTP
//	@ID				create-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			tenentId	path		int	true	"tenent id"
//	@Success		201	{object}   	model.TOTPWithSecretST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp/{tenentId} [post]
//
//	@Security		Authorization
func PostCurrentUserCreateTOTP(c *fiber.Ctx) error {
	tenentId, err := strconv.Atoi(c.Params("tenentId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("tenentId", "invalid")
	}
	user := middleware.GetUser(c)
	totp, err := repository.CreateTOTP(user.Id, int32(tenentId))
	if err != nil {
		slog.Error("failed to create TOTP", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	_, err = repository.UpsertMFA(user.Id, totp.Id, "totp")
	if err != nil {
		slog.Error("failed to enable MFA for TOTP", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.TOTPWithSecretFromRow(totp))
}

// PatchCurrentUserEnableTOTP
//
//	@Summary		Enables user TOTP
//	@ID				enable-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			tenentId	path		int	true	"tenent id"
//	@Success		200	{object}   	model.TOTPWithSecretST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp/{tenentId}/enable [patch]
//
//	@Security		Authorization
func PatchCurrentUserEnableTOTP(c *fiber.Ctx) error {
	tenentId, err := strconv.Atoi(c.Params("tenentId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("tenentId", "invalid")
	}
	user := middleware.GetUser(c)
	totp, err := repository.GetTOTPsByUserIdAndTenentId(user.Id, int32(tenentId))
	if err != nil {
		slog.Error("failed to find TOTP", "error", err)
		return model.NewError(http.StatusNotFound).AddError("tenentId", "invalid")
	}
	if totp == nil {
		return model.NewError(http.StatusNotFound).AddError("tenentId", "invalid")
	}
	_, err = repository.UpsertMFA(user.Id, totp.Id, "totp")
	if err != nil {
		slog.Error("failed to enable MFA for TOTP", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	totp.Enabled = true
	c.Status(http.StatusOK)
	return c.JSON(model.TOTPWithSecretFromRow(*totp))
}

// DeleteCurrentUserDisableTOTP
//
//	@Summary		Disables user TOTP
//	@ID				disalbe-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			tenentId	path		int	true	"tenent id"
//	@Success		200	{object}   	model.TOTPWithSecretST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp/{tenentId}/enable [delete]
//
//	@Security		Authorization
func DeleteCurrentUserDisableTOTP(c *fiber.Ctx) error {
	tenentId, err := strconv.Atoi(c.Params("tenentId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("tenentId", "invalid")
	}
	user := middleware.GetUser(c)
	totp, err := repository.GetTOTPsByUserIdAndTenentId(user.Id, int32(tenentId))
	if err != nil {
		slog.Error("failed to find TOTP", "error", err)
		return model.NewError(http.StatusNotFound).AddError("tenentId", "invalid")
	}
	if totp == nil {
		return model.NewError(http.StatusNotFound).AddError("tenentId", "invalid")
	}
	_, err = repository.DeleteMFA(user.Id)
	if err != nil {
		slog.Error("failed to disable MFA for TOTP", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	totp.Enabled = false
	c.Status(http.StatusOK)
	return c.JSON(model.TOTPWithSecretFromRow(*totp))
}

// DeleteUserTOTP
//
//	@Summary		Delete user TOTP
//	@ID				delete-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			tenentId	path		int	true	"tenent id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp/{tenentId} [delete]
//
//	@Security		Authorization
func DeleteCurrentUserTOTP(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	tenentId, err := strconv.Atoi(c.Params("tenentId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("tenentId", "invalid")
	}
	deleted, err := repository.DeleteTOTP(user.Id, int32(tenentId))
	if err != nil {
		slog.Error("failed to delete totp", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("tenentId", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
