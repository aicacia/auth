package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/access"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// GetTenents
//
//	@Summary		Get application tenents
//	@ID				application-tenents
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			limit	query		int	false	"limit"
//	@Param			offset	query		int	false	"offset"
//	@Success		200	{object}   	model.PaginationST[model.TenentST]
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents [get]
//
//	@Security		Authorization
func GetTenents(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid").Send(c)
	}
	limit, offset, err := GetLimitAndOffset(c, 20)
	if err != nil {
		if err == errParseLimitOffset {
			return nil
		}
		return err
	}
	applicationTenents, err := repository.GetTenents(int32(applicationId), limit, offset)
	if err != nil {
		log.Printf("failed to get applications: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	return c.JSON(model.PaginationST[model.TenentST]{
		HasMore: len(applicationTenents) == limit,
		Items:   util.Map(applicationTenents, model.TenentFromTenentRow),
	})
}

// GetTenentById
//
//	@Summary		Get application tenent by id
//	@ID				application-tenent-by-id
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application tenent id"
//	@Success		200	{object}   	model.TenentST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents/{id} [get]
//
//	@Security		Authorization
func GetTenentById(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid").Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	application, err := repository.GetTenentById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	return c.JSON(model.TenentFromTenentRow(*application))
}

// PostCreateTenent
//
//	@Summary		Create application tenent
//	@ID				create-application-tenent
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			application	body		model.CreateTenentST	true	"create application"
//	@Success		201	{object}   	model.TenentST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents [post]
//
//	@Security		Authorization
func PostCreateTenent(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid").Send(c)
	}
	var createTenent model.CreateTenentST
	if err := c.BodyParser(&createTenent); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	application, err := repository.CreateTenent(int32(applicationId), createTenent.CreateTenentST)
	if err != nil {
		log.Printf("failed to create tenent: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.TenentFromTenentRow(application))
}

// PatchUpdateTenent
//
//	@Summary		Update application tenent
//	@ID				update-application-tenent
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application tenent id"
//	@Param			application	body		model.UpdateTenentST	true	"update application"
//	@Success		200	{object}   	model.TenentST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents/{id} [patch]
//
//	@Security		Authorization
func PatchUpdateTenent(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid").Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	var updateTenent model.UpdateTenentST
	if err := c.BodyParser(&updateTenent); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	application, err := repository.UpdateTenent(int32(id), updateTenent.UpdateTenentST)
	if err != nil {
		log.Printf("failed to update tenent: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	return c.JSON(model.TenentFromTenentRow(*application))
}

// DeleteTenent
//
//	@Summary		Delete application tenent
//	@ID				delete-application-tenent
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application tenent id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents/{id} [delete]
//
//	@Security		Authorization
func DeleteTenent(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid").Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	deleted, err := repository.DeleteTenent(int32(id))
	if err != nil {
		log.Printf("failed to delete tenent: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
