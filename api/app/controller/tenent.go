package controller

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/aicacia/auth/api/app/access"
	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/app/util"
	"github.com/gofiber/fiber/v2"
)

// GetTenents
//
//	@Summary		Get application tenents
//	@ID				tenents
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			query	query		model.OffsetAndLimitQueryST	false	"query"
//	@Success		200	{object}   	model.PaginationST[model.TenentST]
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents [get]
//
//	@Security		Authorization
func GetTenents(c *fiber.Ctx) error {
	if err := access.HasAction(c, "tenents", "read"); err != nil {
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		slog.Error("failed to parse applicationId", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	var offsetAndLimit model.OffsetAndLimitQueryST
	if err := c.QueryParser(&offsetAndLimit); err != nil {
		slog.Error("failed to parse query", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("query", "invalid")
	}
	tenents, err := repository.GetTenents(int32(applicationId), offsetAndLimit.Limit, offsetAndLimit.Offset)
	if err != nil {
		slog.Error("failed to get applications", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	hasMore := false
	if offsetAndLimit.Limit != nil && *offsetAndLimit.Limit == len(tenents) {
		hasMore = true
	}
	return c.JSON(model.PaginationST[model.TenentST]{
		HasMore: hasMore,
		Items:   util.Map(tenents, model.TenentFromRow),
	})
}

// GetTenentById
//
//	@Summary		Get application tenent by id
//	@ID				tenent
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
	if err := access.HasAction(c, "tenents", "read"); err != nil {
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
	tenent, err := repository.GetTenentById(int32(id))
	if err != nil {
		slog.Error("failed to get tenent", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if tenent == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.TenentFromRow(*tenent))
}

// GetTenentPrivateKeyById
//
//	@Summary		Get application tenent by id
//	@ID				tenent-private-key
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application tenent id"
//	@Success		200	{object}    string
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents/{id}/private-key [get]
//
//	@Security		Authorization
func GetTenentPrivateKeyById(c *fiber.Ctx) error {
	if err := access.HasAction(c, "tenents", "read"); err != nil {
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
	tenent, err := repository.GetTenentById(int32(id))
	if err != nil {
		slog.Error("failed to get tenent", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if tenent == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(tenent.PrivateKey)
}

// PostCreateTenent
//
//	@Summary		Create application tenent
//	@ID				create-tenent
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			tenent	body		model.CreateTenentST	true	"create application"
//	@Success		201	{object}   	model.TenentST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents [post]
//
//	@Security		Authorization
func PostCreateTenent(c *fiber.Ctx) error {
	if err := access.HasAction(c, "tenents", "write"); err != nil {
		return err
	}
	applicationId, err := strconv.Atoi(c.Params("applicationId"))
	if err != nil {
		return model.NewError(http.StatusBadRequest).AddError("applicationId", "invalid")
	}
	var createTenent model.CreateTenentST
	if err := c.BodyParser(&createTenent); err != nil {
		slog.Error("failed to parse body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	tenent, err := repository.CreateTenent(int32(applicationId), createTenent.CreateTenentST)
	if err != nil {
		slog.Error("failed to create tenent", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	c.Status(http.StatusCreated)
	return c.JSON(model.TenentFromRow(tenent))
}

// PatchUpdateTenent
//
//	@Summary		Update application tenent
//	@ID				update-tenent
//	@Tags			tenent
//	@Accept			json
//	@Produce		json
//	@Param			applicationId	path		int	true	"application id"
//	@Param			id	path		int	true	"application tenent id"
//	@Param			tenent	body		model.UpdateTenentST	true	"update application"
//	@Success		200	{object}   	model.TenentST
//	@Failure		400	{object}	model.ErrorST
//	@Failure		401	{object}	model.ErrorST
//	@Failure		403	{object}	model.ErrorST
//	@Failure		500	{object}	model.ErrorST
//	@Router			/applications/{applicationId}/tenents/{id} [patch]
//
//	@Security		Authorization
func PatchUpdateTenent(c *fiber.Ctx) error {
	if err := access.HasAction(c, "tenents", "write"); err != nil {
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
	var updateTenent model.UpdateTenentST
	if err := c.BodyParser(&updateTenent); err != nil {
		slog.Error("failed to parse body", "error", err)
		return model.NewError(http.StatusBadRequest).AddError("request", "invalid")
	}
	currentTenent, err := repository.GetTenentById(int32(id))
	if err != nil {
		slog.Error("failed to find tenent", "error", err)
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	if currentTenent == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	errors := model.NewError(http.StatusBadRequest)
	alg := currentTenent.Algorithm
	if updateTenent.Algorithm != nil {
		alg = *updateTenent.Algorithm
	}
	if updateTenent.PrivateKey != nil {
		privateKey := strings.TrimSpace(*updateTenent.PrivateKey)
		updateTenent.PrivateKey = &privateKey
		_, err := jwt.ParsePrivateKey(alg, *updateTenent.PrivateKey)
		if err != nil {
			slog.Error("failed to parse private key", "error", err)
			errors.AddError("privateKey", "invalid")
		}
	}
	if updateTenent.PublicKey != nil {
		publicKey := strings.TrimSpace(*updateTenent.PublicKey)
		updateTenent.PublicKey = &publicKey
		_, err := jwt.ParsePublicKey(alg, *updateTenent.PublicKey, *updateTenent.PrivateKey)
		if err != nil {
			slog.Error("failed to parse public key", "error", err)
			errors.AddError("publicKey", "invalid")
		}
	}
	if errors.HasErrors() {
		return errors
	}
	tenent, err := repository.UpdateTenent(int32(id), updateTenent.UpdateTenentST)
	if err != nil {
		slog.Error("failed to update tenent", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if tenent == nil {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	return c.JSON(model.TenentFromRow(*tenent))
}

// DeleteTenent
//
//	@Summary		Delete application tenent
//	@ID				delete-tenent
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
	if err := access.HasAction(c, "tenents", "write"); err != nil {
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
	deleted, err := repository.DeleteTenent(int32(id))
	if err != nil {
		slog.Error("failed to delete tenent", "error", err)
		return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
	}
	if !deleted {
		return model.NewError(http.StatusNotFound).AddError("id", "invalid")
	}
	c.Status(http.StatusNoContent)
	return c.Send(nil)
}
