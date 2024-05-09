package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/access"
	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// GetPermissions
//
//	@Summary		Get application permissions
//	@ID				application-permissions
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Success		200	{array}   	model.PermissionST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions [get]
//
//	@Security		Authorization
func GetPermissions(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	applications, err := repository.GetPermissions(int32(applicationId))
	if err != nil {
		log.Printf("failed to get applications: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	return c.JSON(util.Map(applications, model.PermissionFromRow))
}

// GetPermissionById
//
//	@Summary		Get application permission by id
//	@ID				application-permission-by-id
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		string	true	"application permission id"
//	@Success		200	{object}   	model.PermissionST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions/{id} [get]
//
//	@Security		Authorization
func GetPermissionById(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application, err := repository.GetPermissionById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.PermissionFromRow(*application))
}

// PostCreatePermission
//
//	@Summary		Create application permission
//	@ID				create-application-permission
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			application	body		model.CreatePermissionST	true	"create application"
//	@Success		201	{object}   	model.PermissionST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions [post]
//
//	@Security		Authorization
func PostCreatePermission(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	var createPermission model.CreatePermissionST
	if err := c.BodyParser(&createPermission); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application, err := repository.CreatePermission(int32(applicationId), createPermission.CreatePermissionST)
	if err != nil {
		log.Printf("failed to create permission: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.PermissionFromRow(application))
}

// PatchUpdatePermission
//
//	@Summary		Update application permission
//	@ID				update-application-permission
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application permission id"
//	@Param			application	body		model.UpdatePermissionST	true	"update application"
//	@Success		200	{object}   	model.PermissionST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions/{id} [patch]
//
//	@Security		Authorization
func PatchUpdatePermission(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	var updatePermission model.UpdatePermissionST
	if err := c.BodyParser(&updatePermission); err != nil {
		log.Printf("failed to parse body: %v\n", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	application, err := repository.UpdatePermission(int32(id), updatePermission.UpdatePermissionST)
	if err != nil {
		log.Printf("failed to update permission: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if application == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.PermissionFromRow(*application))
}

// DeletePermission
//
//	@Summary		Delete application permission
//	@ID				delete-application-permission
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		string	true	"application permission id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions/{id} [delete]
//
//	@Security		Authorization
func DeletePermission(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	_, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	deleted, err := repository.DeletePermission(int32(id))
	if err != nil {
		log.Printf("failed to delete permission: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// PatchAddPermissionToUser
//
//	@Summary		Add permission to user
//	@ID				add-application-permission-to-user
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application permission id"
//	@Param			userId	path		int	true	"user id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions/{id}/add-user/{userId} [patch]
//
//	@Security		Authorization
func PatchAddPermissionToUser(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	permission, err := repository.GetPermissionById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	added, err := repository.AddPermissionToUser(user.Id, permission.Id)
	if err != nil {
		log.Printf("failed to add permission to user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !added {
		return model.NewError(http.StatusNotFound).AddError("user", "found", "application")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}

// DeleteRemovePermissionFromUser
//
//	@Summary		Remove permission from user
//	@ID				remove-application-permission-from-user
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application permission id"
//	@Param			userId	path		int	true	"user id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/permissions/{id}/remove-user/{userId} [delete]
//
//	@Security		Authorization
func DeleteRemovePermissionFromUser(c *fiber.Ctx) error {
	if err := access.IsAdmin(c); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("id", "invalid")
	}
	application := middleware.GetApplication(c)
	permission, err := repository.GetPermissionById(int32(id))
	if err != nil {
		log.Printf("failed to get application: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("userId", "invalid")
	}
	user, err := repository.GetUserById(application.Id, int32(userId))
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	if user == nil {
		return model.NewError(http.StatusNotFound).AddError("userId", "invalid")
	}
	removed, err := repository.RemovePermissionFromUser(user.Id, permission.Id)
	if err != nil {
		log.Printf("failed to add permission to user: %v\n", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !removed {
		return model.NewError(http.StatusNotFound).AddError("user", "notFound")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
