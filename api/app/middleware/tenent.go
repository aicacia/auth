package middleware

import (
	"log"
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
			log.Printf("invalid tenent id %s: %v\n", tenentIdString, err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid")
		}
		tenent, err := repository.GetTenentByClientId(tenentId)
		if err != nil {
			log.Printf("failed to fetch application tenent: %v\n", err)
			return model.NewError(http.StatusNotFound).AddError("Tenent-Id", "invalid")
		}
		if tenent == nil {
			log.Printf("tenent not found: %v\n", err)
			return model.NewError(http.StatusNotFound).AddError("Tenent-Id", "invalid")
		}
		application, err := repository.GetApplicationById(tenent.ApplicationId)
		if err != nil {
			log.Printf("failed to fetch application: %v\n", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid")
		}
		if application == nil {
			log.Printf("application not found: %v\n", err)
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
