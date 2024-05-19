package middleware

import (
	"log"
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
			log.Printf("failed to get authorization header: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		if unvalidatedClaims.Type != jwt.MFATokenType {
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		tenent, err := repository.GetTenentByClientId(unvalidatedClaims.ClientId)
		if err != nil {
			log.Printf("failed to fetch application tenent: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		claims, err := jwt.ParseClaimsFromToken[jwt.MFAClaims](tokenString, tenent)
		if err != nil {
			log.Printf("failed to parse claims from token: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		application, err := repository.GetApplicationById(tenent.ApplicationId)
		if err != nil {
			log.Printf("failed to fetch application: %v", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
		}
		c.Locals(applicationLocalKey, application)
		c.Locals(tenentLocalKey, tenent)
		c.Locals(baseClaimsLocalKey, claims)

		switch claims.SubjectType {
		case jwt.UserSubject:
			user, err := repository.GetUserById(application.Id, claims.Subject)
			if err != nil {
				log.Printf("failed to fetch user: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			permissions, err := repository.GetUserPermissions(user.Id)
			if err != nil {
				log.Printf("failed to fetch user permissions: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			c.Locals(permissionsLocalKey, permissions)
			c.Locals(permissionsMapLocalKey, PermissionsFromRows(permissions))
			c.Locals(userLocalKey, user)
		default:
			log.Printf("invalid subject type: %v\n", claims.SubjectType)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		return c.Next()
	}
}
