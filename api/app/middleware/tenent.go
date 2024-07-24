package middleware

import (
	"log/slog"
	"net/http"

	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var applicationLocalKey = "application"
var tenentLocalKey = "tenent"

func TenentMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenentIdString := c.Get("Tenent-Id")
		tenentId, err := uuid.Parse(tenentIdString)
		if err != nil {
			slog.Error("invalid tenent id", "tenentId", tenentIdString, "error", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid")
		}
		tenent, err := repository.GetTenentByClientId(tenentId)
		if err != nil {
			slog.Error("failed to fetch application tenent", "error", err)
			return model.NewError(http.StatusNotFound).AddError("Tenent-Id", "invalid")
		}
		if tenent == nil {
			slog.Error("tenent not found", "error", err)
			return model.NewError(http.StatusNotFound).AddError("Tenent-Id", "invalid")
		}
		application, err := repository.GetApplicationById(tenent.ApplicationId)
		if err != nil {
			slog.Error("failed to fetch application", "error", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid")
		}
		if application == nil {
			slog.Error("application not found", "error", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid")
		}
		c.Locals(applicationLocalKey, application)
		c.Locals(tenentLocalKey, tenent)
		return c.Next()
	}
}

func GetApplication(c *fiber.Ctx) *repository.ApplicationRowST {
	application := c.Locals(applicationLocalKey)
	return application.(*repository.ApplicationRowST)
}

func GetTenent(c *fiber.Ctx) *repository.TenentRowST {
	tenent := c.Locals(tenentLocalKey)
	return tenent.(*repository.TenentRowST)
}
