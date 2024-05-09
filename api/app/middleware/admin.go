package middleware

import (
	"net/http"

	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

func AdminApplicationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		application := GetApplication(c)
		if !application.IsAdmin {
			return model.NewError(http.StatusForbidden).AddError("authorization", "invalid")
		}
		return c.Next()
	}
}
