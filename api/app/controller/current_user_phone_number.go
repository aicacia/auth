package controller

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// PatchCurrentUserPhoneNumberSendConfirmation
//
//	@Summary		Send confirmation token to user phone_number
//	@ID				send-confirmation-to-phone-number
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"phone_number id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/phone-numbers/{id}/send-confirmation [patch]
//
//	@Security		Authorization
func PatchCurrentUserPhoneNumberSendConfirmation(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	confirmationToken, err := util.GenerateRandomHex(6)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to phone_number
	slog.Info("sending text", "userId", user.Id, "phoneNumberId", id, "token", confirmationToken)
	_, err = repository.SetPhoneNumberConfirmation(user.Id, int32(id), confirmationToken)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchCurrentUserPhoneNumberConfirm
//
//	@Summary		Confirm phone_number with token
//	@ID				confirm-phone-number
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"phone_number id"
//	@Param			confirmPhoneNumber	body		model.ConfirmPhoneNumberST	true	"phone_number confirmation"
//	@Success		200 {object}	model.PhoneNumberST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/phone-numbers/{id}/confirm [patch]
//
//	@Security		Authorization
func PatchCurrentUserPhoneNumberConfirm(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var confirmPhoneNumber model.ConfirmPhoneNumberST
	if err := c.BodyParser(&confirmPhoneNumber); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	phone_number, err := repository.ConfirmPhoneNumber(user.Id, int32(id), strings.ToLower(strings.TrimSpace(confirmPhoneNumber.Token)))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.PhoneNumberFromRow(phone_number))
}

// PatchCurrentUserPhoneNumberSetPrimary
//
//	@Summary		Set a confirmed phone to primary
//	@ID				set-primary-phone-number
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/phone-numbers/{id}/set-primary [patch]
//
//	@Security		Authorization
func PatchCurrentUserPhoneNumberSetPrimary(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	_, err = repository.SetPrimaryPhoneNumber(user.Id, int32(id))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PostCurrentUserCreatePhoneNumber
//
//	@Summary		Create user phone number
//	@ID				create-phone-number
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			createPhoneNumber	body		model.CreatePhoneNumberST	true	"update phone_number"
//	@Success		201	{object}   	model.PhoneNumberST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/phone-numbers [post]
//
//	@Security		Authorization
func PostCurrentUserCreatePhoneNumber(c *fiber.Ctx) error {
	var createPhoneNumber model.CreatePhoneNumberST
	if err := c.BodyParser(&createPhoneNumber); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	phoneNumber := util.NumericRegex.ReplaceAllString(createPhoneNumber.PhoneNumber, "")
	if phoneNumber == "" {
		return model.NewError(http.StatusBadRequest).AddError("phoneNumber", "required")
	}
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return model.NewError(http.StatusBadRequest).AddError("phoneNumber", "invalid")
	}
	user := middleware.GetUser(c)
	confirmationToken, err := util.GenerateRandomHex(6)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	phoneNumberRow, err := repository.CreatePhoneNumber(user.ApplicationId, user.Id, phoneNumber, confirmationToken)
	if err != nil {
		slog.Error("failed to create phone_number", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to phone number
	slog.Info("send text", "userId", user.Id, "phoneNumberId", phoneNumberRow.Id, "token", strings.ToUpper(*phoneNumberRow.ConfirmationToken))
	c.Status(http.StatusCreated)
	return c.JSON(model.PhoneNumberFromRow(phoneNumberRow))
}

// DeleteCurrentUserPhoneNumber
//
//	@Summary		Delete user phone number
//	@ID				delete-phone-number
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"phone_number id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/phone-numbers/{id} [delete]
//
//	@Security		Authorization
func DeleteCurrentUserPhoneNumber(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	deleted, err := repository.DeletePhoneNumber(user.Id, int32(id))
	if err != nil {
		slog.Error("failed to delete phone_number", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
