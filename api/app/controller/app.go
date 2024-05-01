package controller

import (
	"net/http"
	"time"

	"github.com/aicacia/auth/api/app"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/aicacia/auth/api/docs"
	"github.com/gofiber/fiber/v2"
)

func GetOpenAPI(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.SendString(docs.SwaggerInfo.ReadDoc())
}

// GetHealthCheck
//
//	@ID				healthCheck
//	@Summary		Get Health Check
//	@Tags			app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.HealthST
//	@Failure		500	{object}	model.HealthST
//	@Router			/health [get]
func GetHealthCheck(c *fiber.Ctx) error {
	health := model.HealthST{
		DB:   repository.ValidConnection(),
		Date: time.Now().UTC(),
	}
	if health.IsHealthy() {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusInternalServerError)
	}
	return c.JSON(health)
}

// GetVersion
//
//	@ID				version
//	@Summary		Get Version
//	@Tags			app
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.VersionST
//	@Router			/version [get]
func GetVersion(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(app.Version)
}
