package access

import (
	"net/http"
	"slices"

	"github.com/aicacia/auth/api/app/middleware"
	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

func HasAction(c *fiber.Ctx, resource string, actions ...string) *model.ErrorST {
	if allowedActions, ok := middleware.GetPermissionsMap(c)[resource]; ok {
		for _, action := range allowedActions {
			if !slices.Contains(allowedActions, action) {
				return model.NewError(http.StatusForbidden).AddError("authorization", "invalid")
			}
		}
	}
	return nil
}

func UserIsOwnerOrHasAction(c *fiber.Ctx, userId int32, actions ...string) *model.ErrorST {
	user := middleware.GetUser(c)
	if user.Id == userId {
		return nil
	}
	return HasAction(c, "users", actions...)
}
