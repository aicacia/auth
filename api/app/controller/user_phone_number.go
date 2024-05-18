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

// PatchUserPhoneNumberSendConfirmation
//
//	@Summary		Send confirmation token to user phone_number
//	@ID				send-confirmation-to-phone-number
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"phone_number id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/phone-numbers/{id}/send-confirmation [patch]
//
//	@Security		Authorization
func PatchUserPhoneNumberSendConfirmation(c *fiber.Ctx) error {
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
	// TODO: send to phone_number
	log.Printf("userId=%d, emailId=%d, token=%s\n", userId, id, confirmationToken)
	_, err = repository.SetPhoneNumberConfirmation(user.Id, int32(id), confirmationToken)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchUserPhoneNumberConfirm
//
//	@Summary		Confirm phone_number with token
//	@ID				confirm-phone-number
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"phone_number id"
//	@Param			confirmPhoneNumber	body		model.ConfirmPhoneNumberST	true	"phone_number confirmation"
//	@Success		200 {object}	model.PhoneNumberST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/phone-numbers/{id}/confirm [patch]
//
//	@Security		Authorization
func PatchUserPhoneNumberConfirm(c *fiber.Ctx) error {
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
	var confirmPhoneNumber model.ConfirmPhoneNumberST
	if err := c.BodyParser(&confirmPhoneNumber); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	phone_number, err := repository.ConfirmPhoneNumber(user.Id, int32(id), strings.ToLower(strings.TrimSpace(confirmPhoneNumber.Token)))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(model.PhoneNumberFromRow(phone_number))
}

// PatchUserPhoneNumberSetPrimary
//
//	@Summary		Set a confirmed phone to primary
//	@ID				set-primary-phone-number
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
//	@Router			/applications/{applicationId}/users/{userId}/phone-numbers/{id}/set-primary [patch]
//
//	@Security		Authorization
func PatchUserPhoneNumberSetPrimary(c *fiber.Ctx) error {
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
	_, err = repository.SetPrimaryPhoneNumber(user.Id, int32(id))
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PostUserCreatePhoneNumber
//
//	@Summary		Create user phone number
//	@ID				create-phone-number
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			createPhoneNumber	body		model.CreatePhoneNumberST	true	"update phone_number"
//	@Success		201	{object}   	model.PhoneNumberST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/phone-numbers [post]
//
//	@Security		Authorization
func PostUserCreatePhoneNumber(c *fiber.Ctx) error {
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
	var createPhoneNumber model.CreatePhoneNumberST
	if err := c.BodyParser(&createPhoneNumber); err != nil {
		log.Printf("invalid request body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	confirmationToken, err := util.GenerateRandomHex(6)
	if err != nil {
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	phoneNumberRow, err := repository.CreatePhoneNumber(int32(applicationId), int32(userId), createPhoneNumber.PhoneNumber, confirmationToken)
	if err != nil {
		log.Printf("failed to create phone_number: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	// TODO: send to phone number
	log.Printf("userId=%d, emailId=%d, token=%s\n", userId, phoneNumberRow.Id, strings.ToUpper(*phoneNumberRow.ConfirmationToken))
	c.Status(http.StatusCreated)
	return c.JSON(model.PhoneNumberFromRow(phoneNumberRow))
}

// DeleteUserPhoneNumber
//
//	@Summary		Delete user phone number
//	@ID				delete-phone-number
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			userId	path		int	true	"user id"
//	@Param			id	path		int	true	"phone_number id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/users/{userId}/phone-numbers/{id} [delete]
//
//	@Security		Authorization
func DeleteUserPhoneNumber(c *fiber.Ctx) error {
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
	deleted, err := repository.DeletePhoneNumber(int32(userId), int32(id))
	if err != nil {
		log.Printf("failed to delete phone_number: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
