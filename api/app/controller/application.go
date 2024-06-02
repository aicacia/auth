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
//	@Param			query	query		model.OffsetAndLimitQueryST	false	"query"
//	@Success		200	{object}   	model.PaginationST[model.ApplicationST]
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications [get]
//
//	@Security		Authorization
func GetApplications(c *fiber.Ctx) error {
	if err := access.HasAction(c, "applications", "read"); err != nil {
		return err
	}
	var offsetAndLimit model.OffsetAndLimitQueryST
	if err := c.QueryParser(&offsetAndLimit); err != nil {
		log.Printf("failed to parse query: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("query", "invalid")
	}
	applications, err := repository.GetApplications(offsetAndLimit.Limit, offsetAndLimit.Offset)
	if err != nil {
		log.Printf("failed to get applications: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	hasMore := false
	if offsetAndLimit.Limit != nil && *offsetAndLimit.Limit == len(applications) {
		hasMore = true
	}
	return c.JSON(model.PaginationST[model.ApplicationST]{
		HasMore: hasMore,
		Items:   util.Map(applications, model.ApplicationFromRow),
	})
}

// GetApplicationById
//
//	@Summary		Get application by id
//	@ID				application
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
	if err := access.HasAction(c, "applications", "read"); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application, err := repository.GetApplicationById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.ApplicationFromRow(*application))
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
	if err := access.HasAction(c, "applications", "write"); err != nil {
		return err
	}
	var createApplication model.CreateApplicationST
	if err := c.BodyParser(&createApplication); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application, err := repository.CreateApplication(createApplication.CreateApplicationST)
	if err != nil {
		log.Printf("failed to create application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.ApplicationFromRow(application))
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
	if err := access.HasAction(c, "applications", "write"); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var updateApplication model.UpdateApplicationST
	if err := c.BodyParser(&updateApplication); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application, err := repository.UpdateApplication(int32(id), updateApplication.UpdateApplicationST)
	if err != nil {
		log.Printf("failed to update application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.ApplicationFromRow(*application))
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
	if err := access.HasAction(c, "applications", "write"); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application, err := repository.GetApplicationById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	if application.IsAdmin {
		return model.NewError(http.StatusForbidden).AddError("internal", "cannotDeleteAdmin")
	}
	deleted, err := repository.DeleteApplication(int32(id))
	if err != nil {
		log.Printf("failed to delete application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
