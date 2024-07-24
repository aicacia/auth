package middleware

import (
	"log/slog"
	"net/http"

	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
)

func MFAAuthorizedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, tokenString := GetAuthorizationFromContext(c)
		unvalidatedClaims, err := jwt.ParseClaimsFromTokenNoValidation(tokenString)
		if err != nil {
			slog.Error("failed to get authorization header", "error", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		if unvalidatedClaims.Type != jwt.MFATokenType {
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		tenent, err := repository.GetTenentByClientId(unvalidatedClaims.ClientId)
		if err != nil {
			slog.Error("failed to fetch application tenent", "error", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		claims, err := jwt.ParseClaimsFromToken[jwt.MFAClaims](tokenString, tenent)
		if err != nil {
			slog.Error("failed to parse claims from token", "error", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		application, err := repository.GetApplicationById(tenent.ApplicationId)
		if err != nil {
			slog.Error("failed to fetch application", "error", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
		}
		c.Locals(applicationLocalKey, application)
		c.Locals(tenentLocalKey, tenent)
		c.Locals(baseClaimsLocalKey, claims)

		switch claims.SubjectType {
		case jwt.UserSubject:
			user, err := repository.GetUserById(application.Id, claims.Subject)
			if err != nil {
				slog.Error("failed to fetch user", "error", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			permissions, err := repository.GetUserPermissions(user.Id)
			if err != nil {
				slog.Error("failed to fetch user permissions", "error", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			c.Locals(permissionsLocalKey, permissions)
			c.Locals(permissionsMapLocalKey, PermissionsFromRows(permissions))
			c.Locals(userLocalKey, user)
		default:
			slog.Error("invalid subject type", "type", claims.SubjectType)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		return c.Next()
	}
}
