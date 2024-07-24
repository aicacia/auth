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

// PatchCurrentUserEmailSendConfirmation
//
//	@Summary		Send confirmation token to user email
//	@ID				send-confirmation-to-email
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/emails/{id}/send-confirmation [patch]
//
//	@Security		Authorization
func PatchCurrentUserEmailSendConfirmation(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	confirmationToken, err := util.GenerateRandomHex(6)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to email
	slog.Info("sending email", "userId", user.Id, "emailId", id, "token", confirmationToken)
	_, err = repository.SetEmailConfirmation(user.Id, int32(id), confirmationToken)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchCurrentUserEmailConfirm
//
//	@Summary		Confirm email with token
//	@ID				confirm-email
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"email id"
//	@Param			confirmEmail	body		model.ConfirmEmailST	true	"email confirmation"
//	@Success		200 {object}	model.EmailST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/emails/{id}/confirm [patch]
//
//	@Security		Authorization
func PatchCurrentUserEmailConfirm(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var confirmEmail model.ConfirmEmailST
	if err := c.BodyParser(&confirmEmail); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	email, err := repository.ConfirmEmail(user.Id, int32(id), strings.ToLower(strings.TrimSpace(confirmEmail.Token)))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.EmailFromRow(email))
}

// PatchCurrentUserEmailSetPrimary
//
//	@Summary		Set a confirmed email to primary
//	@ID				set-primary-email
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/emails/{id}/set-primary [patch]
//
//	@Security		Authorization
func PatchCurrentUserEmailSetPrimary(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	_, err = repository.SetPrimaryEmail(user.Id, int32(id))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PostCurrentUserCreateEmail
//
//	@Summary		Create user email
//	@ID				create-email
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			createEmail	body		model.CreateEmailST	true	"update email"
//	@Success		201	{object}   	model.EmailST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/emails [post]
//
//	@Security		Authorization
func PostCurrentUserCreateEmail(c *fiber.Ctx) error {
	var createEmail model.CreateEmailST
	if err := c.BodyParser(&createEmail); err != nil {
		slog.Error("invalid request body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	user := middleware.GetUser(c)
	confirmationToken, err := util.GenerateRandomHex(6)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	email := strings.TrimSpace(createEmail.Email)
	if email == "" {
		return model.NewError(http.StatusBadRequest).AddError("email", "required")
	}
	emailRow, err := repository.CreateEmail(user.ApplicationId, user.Id, email, confirmationToken)
	if err != nil {
		slog.Error("failed to create email", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to email
	slog.Info("sending email", "userId", user.Id, "emailId", emailRow.Id, "token", strings.ToUpper(*emailRow.ConfirmationToken))
	c.Status(http.StatusCreated)
	return c.JSON(model.EmailFromRow(emailRow))
}

// DeleteUserEmail
//
//	@Summary		Delete user email
//	@ID				delete-email
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/emails/{id} [delete]
//
//	@Security		Authorization
func DeleteCurrentUserEmail(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	deleted, err := repository.DeleteEmail(user.Id, int32(id))
	if err != nil {
		slog.Error("failed to delete email", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
