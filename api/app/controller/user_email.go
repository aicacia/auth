package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// PatchUserEmailSendConfirmation
//
//	@Summary		Send confirmation token to user email
//	@ID				send-confirmation-to-email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/emails/{id}/send-confirmation [patch]
//
//	@Security		Authorization
func PatchUserEmailSendConfirmation(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	confirmationToken, err := util.GenerateRandomHex(8)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to email
	log.Printf("userId=%d, emailId=%d, token=%s\n", userId, id, confirmationToken)
	_, err = repository.SetEmailConfirmation(user.Id, int32(id), confirmationToken)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchUserEmailConfirm
//
//	@Summary		Confirm email with token
//	@ID				confirm-email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"email id"
//	@Param			confirmEmail	body		model.ConfirmEmailST	true	"email confirmation"
//	@Success		200 {object}	model.EmailST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/emails/{id}/confirm [patch]
//
//	@Security		Authorization
func PatchUserEmailConfirm(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var confirmEmail model.ConfirmEmailST
	if err := c.BodyParser(&confirmEmail); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	email, err := repository.ConfirmEmail(user.Id, int32(id), strings.ToLower(strings.TrimSpace(confirmEmail.Token)))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.EmailFromRow(email))
}

// PatchUserEmailSetPrimary
//
//	@Summary		Set a confirmed email to primary
//	@ID				set-primary-email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/emails/{id}/set-primary [patch]
//
//	@Security		Authorization
func PatchUserEmailSetPrimary(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
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

// PostUserCreateEmail
//
//	@Summary		Create user email
//	@ID				create-email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			createEmail	body		model.CreateEmailST	true	"update email"
//	@Success		201	{object}   	model.EmailST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/emails [post]
//
//	@Security		Authorization
func PostUserCreateEmail(c *fiber.Ctx) error {
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	var createEmail model.CreateEmailST
	if err := c.BodyParser(&createEmail); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	emailRow, err := repository.CreateEmail(int32(applicationId), int32(userId), createEmail.Email)
	if err != nil {
		log.Printf("failed to create email: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to email
	log.Printf("userId=%d, emailId=%d, token=%s\n", userId, emailRow.Id, emailRow.ConfirmationToken)
	c.Status(http.StatusCreated)
	return c.JSON(model.EmailFromRow(emailRow))
}

// DeleteUserEmail
//
//	@Summary		Delete user email
//	@ID				delete-email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"email id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/emails/{id} [delete]
//
//	@Security		Authorization
func DeleteUserEmail(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	application := middleware.GetApplication(c)
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	deleted, err := repository.DeleteEmail(int32(userId), int32(id))
	if err != nil {
		log.Printf("failed to delete email: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
