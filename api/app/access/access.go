package access

import (
	"net/http"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

func HasPermission(c *fiber.Ctx, permissions ...string) bool {
	permissionsMap := middleware.GetPermissionsMap(c)
	for _, permission := range permissions {
		if !permissionsMap[permission] {
			return false
		}
	}
	return true
}

func IsAdmin(c *fiber.Ctx) *model.ErrorST {
	if !HasPermission(c, "admin") {
		return model.NewError(http.StatusForbidden).AddError("authorization", "invalid")
	}
	return nil
}
