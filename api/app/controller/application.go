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

// GetApplications
//
//	@Summary		Get applications
//	@ID				applications
//	@Tags			application
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"limit"
//	@Param			offset	query		int	false	"offset"
//	@Success		200	{object}   	model.PaginationST[model.ApplicationST]
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications [get]
//
//	@Security		Authorization
func GetApplications(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	limit, offset, err := GetLimitAndOffset(c, 20)
	if err != nil {
		if err == errParseLimitOffset {
			return nil
		}
		return err
	}
	applications, err := repository.GetApplications(limit, offset)
	if err != nil {
		log.Printf("failed to get applications: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	return c.JSON(model.PaginationST[model.ApplicationST]{
		HasMore: len(applications) == limit,
		Items:   util.Map(applications, model.ApplicationFromApplicationRow),
	})
}

// GetApplicationById
//
//	@Summary		Get application by id
//	@ID				application-by-id
//	@Tags			application
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"application id"
//	@Success		200	{object}   	model.ApplicationST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{id} [get]
//
//	@Security		Authorization
func GetApplicationById(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	application, err := repository.GetApplicationById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	return c.JSON(model.ApplicationFromApplicationRow(*application))
}

// PostCreateApplication
//
//	@Summary		Create application
//	@ID				create-application
//	@Tags			application
//	@Accept			json
//	@Produce		json
//	@Param			application	body		model.CreateApplicationST	true	"create application"
//	@Success		201	{object}   	model.ApplicationST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications [post]
//
//	@Security		Authorization
func PostCreateApplication(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	var createApplication model.CreateApplicationST
	if err := c.BodyParser(&createApplication); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	application, err := repository.CreateApplication(createApplication.CreateApplicationST)
	if err != nil {
		log.Printf("failed to create application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.ApplicationFromApplicationRow(application))
}

// PatchUpdateApplication
//
//	@Summary		Update application
//	@ID				update-application
//	@Tags			application
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"application id"
//	@Param			application	body		model.UpdateApplicationST	true	"update application"
//	@Success		200	{object}   	model.ApplicationST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{id} [patch]
//
//	@Security		Authorization
func PatchUpdateApplication(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	var updateApplication model.UpdateApplicationST
	if err := c.BodyParser(&updateApplication); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid").Send(c)
	}
	application, err := repository.UpdateApplication(int32(id), updateApplication.UpdateApplicationST)
	if err != nil {
		log.Printf("failed to update application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	return c.JSON(model.ApplicationFromApplicationRow(*application))
}

// DeleteApplication
//
//	@Summary		Delete application
//	@ID				delete-application
//	@Tags			application
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"application id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{id} [delete]
//
//	@Security		Authorization
func DeleteApplication(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err.Send(c)
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid").Send(c)
	}
	application, err := repository.GetApplicationById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	if application.IsAdmin {
		return model.NewError(http.StatusForbidden).AddError("internal", "cannotDeleteAdmin").Send(c)
	}
	deleted, err := repository.DeleteApplication(int32(id))
	if err != nil {
		log.Printf("failed to delete application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid").Send(c)
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
