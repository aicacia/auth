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
var applicationTenentLocalKey = "application.tenent"
var applicationConfigLocalKey = "application.config"

func TenentMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenentIdString := c.Get("Tenent-Id")
		tenentId, err := uuid.Parse(tenentIdString)
		if err != nil {
			log.Printf("invalid tenent id %s: %v\n", tenentIdString, err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid").Send(c)
		}
		applicationTenent, err := repository.GetTenentByClientId(tenentId)
		if err != nil {
			log.Printf("failed to fetch application tenent: %v\n", err)
			return model.NewError(http.StatusNotFound).AddError("Tenent-Id", "invalid").Send(c)
		}
		application, err := repository.GetApplicationById(applicationTenent.ApplicationId)
		if err != nil {
			log.Printf("failed to fetch application: %v\n", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid").Send(c)
		}
		applicationConfigRows, err := repository.GetApplicationConfigs(applicationTenent.ApplicationId)
		if err != nil {
			log.Printf("failed to fetch application config rows: %v\n", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid").Send(c)
		}
		applicationConfig, err := model.ApplicationConfigFromApplicationConfigRows(applicationConfigRows)
		if err != nil {
			log.Printf("failed to convert application config rows to application config: %v\n", err)
			return model.NewError(http.StatusBadRequest).AddError("Tenent-Id", "invalid").Send(c)
		}
		c.Locals(applicationLocalKey, application)
		c.Locals(applicationTenentLocalKey, applicationTenent)
		c.Locals(applicationConfigLocalKey, applicationConfig)
		return c.Next()
	}
}

func GetApplication(c *fiber.Ctx) *repository.ApplicationRowST {
	application := c.Locals(applicationLocalKey)
	return application.(*repository.ApplicationRowST)
}

func GetTenent(c *fiber.Ctx) *repository.TenentRowST {
	applicationTenent := c.Locals(applicationTenentLocalKey)
	return applicationTenent.(*repository.TenentRowST)
}

func GetApplicationConfig(c *fiber.Ctx) *model.ApplicationConfigST {
	applicationConfig := c.Locals(applicationConfigLocalKey)
	return applicationConfig.(*model.ApplicationConfigST)
}
