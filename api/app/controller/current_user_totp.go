package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
)

// PostCurrentUserCreateTOTP
//
//	@Summary		Create user TOTP
//	@ID				create-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}   	model.TOTPST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp [post]
//
//	@Security		Authorization
func PostCurrentUserCreateTOTP(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	tenent := middleware.GetTenent(c)
	totp, err := repository.CreateTOTP(user.Id, tenent.Id)
	if err != nil {
		log.Printf("failed to create TOTP: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.TOTPST{
		TenentId:  totp.TenentId,
		UserId:    totp.UserId,
		Secret:    totp.Secret,
		UpdatedAt: totp.UpdatedAt,
		CreatedAt: totp.CreatedAt,
	})
}

// DeleteUserTOTP
//
//	@Summary		Create user TOTP
//	@ID				delete-totp
//	@Tags			current-user
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/user/totp [delete]
//
//	@Security		Authorization
func DeleteCurrentUserTOTP(c *fiber.Ctx) error {
	user := middleware.GetUser(c)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	deleted, err := repository.DeleteEmail(user.Id, int32(id))
	if err != nil {
		log.Printf("failed to delete totp: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
