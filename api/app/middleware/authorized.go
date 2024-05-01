package middleware

import (
	"log"
	"net/http"
	"slices"

	"github.com/aicacia/auth/api/app/jwt"
	"github.com/aicacia/auth/api/app/model"
	"github.com/aicacia/auth/api/app/repository"
	"github.com/gofiber/fiber/v2"
)

var baseClaimsLocalKey = "claims"
var userLocalKey = "user"
var serviceAccountLocalKey = "service-account"
var permissionsLocalKey = "permissions"
var permissionsMapLocalKey = "permissions.map"

func AuthorizedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := GetAuthorizationFromContext(c)
		unvalidatedClaims, err := jwt.ParseClaimsFromTokenNoValidation(tokenString)
		if err != nil {
			log.Printf("failed to get authorization header: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid").Send(c)
		}
		if slices.Contains(unvalidatedClaims.Scope, "refresh_token") {
			return model.NewError(http.StatusForbidden).AddError("authorization", "invalid").Send(c)
		}
		applicationTenent, err := repository.GetTenentByClientId(unvalidatedClaims.ClientId)
		if err != nil {
			log.Printf("failed to fetch application tenent: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid").Send(c)
		}
		claims, err := jwt.ParseClaimsFromToken(tokenString, applicationTenent)
		if err != nil {
			log.Printf("failed to parse claims from token: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid").Send(c)
		}
		application, err := repository.GetApplicationById(applicationTenent.ApplicationId)
		if err != nil {
			log.Printf("failed to fetch application: %v", err)
			return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
		}
		c.Locals(applicationLocalKey, application)
		c.Locals(applicationTenentLocalKey, applicationTenent)
		c.Locals(baseClaimsLocalKey, claims)

		switch claims.SubjectType {
		case jwt.UserSubject:
			user, err := repository.GetUserById(claims.Subject)
			if err != nil {
				log.Printf("failed to fetch user: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
			}
			permissions, err := repository.GetUserPermissions(user.Id, application.Id)
			if err != nil {
				log.Printf("failed to fetch user permissions: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
			}
			c.Locals(permissionsLocalKey, permissions)
			c.Locals(permissionsMapLocalKey, repository.PermissionsToMap(permissions))
			c.Locals(userLocalKey, user)
		case jwt.ServiceAccountSubject:
			serviceAccount, err := repository.GetServiceAccountById(claims.Subject)
			if err != nil {
				log.Printf("failed to fetch service account: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
			}
			permissions, err := repository.GetServiceAccountPermissions(serviceAccount.Id, application.Id)
			if err != nil {
				log.Printf("failed to fetch user permissions: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application").Send(c)
			}
			c.Locals(permissionsLocalKey, permissions)
			c.Locals(permissionsMapLocalKey, repository.PermissionsToMap(permissions))
			c.Locals(serviceAccountLocalKey, serviceAccount)
		}
		return c.Next()
	}
}

func OpenIdMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if HasScope(c, "openid") {
			return c.Next()
		}
		return model.NewError(http.StatusUnauthorized).AddError("token", "invalid").Send(c)
	}
}

func GetClaims(c *fiber.Ctx) *jwt.Claims {
	claims := c.Locals(baseClaimsLocalKey)
	return claims.(*jwt.Claims)
}

func HasScope(c *fiber.Ctx, scope ...string) bool {
	claims := GetClaims(c)
	for _, s := range scope {
		if !slices.Contains(claims.Scope, s) {
			return false
		}
	}
	return true
}

func IsUserSubject(c *fiber.Ctx) bool {
	return GetClaims(c).SubjectType == jwt.UserSubject
}

func GetUser(c *fiber.Ctx) *repository.UserRowST {
	user := c.Locals(userLocalKey)
	return user.(*repository.UserRowST)
}

func IsServiceAccount(c *fiber.Ctx) bool {
	return GetClaims(c).SubjectType == jwt.ServiceAccountSubject
}

func GetServiceAccount(c *fiber.Ctx) *repository.ServiceAccountRowST {
	serviceAccount := c.Locals(serviceAccountLocalKey)
	return serviceAccount.(*repository.ServiceAccountRowST)
}

func GetPermissions(c *fiber.Ctx) []repository.PermissionRowST {
	permissions := c.Locals(permissionsLocalKey)
	return permissions.([]repository.PermissionRowST)
}

func GetPermissionsMap(c *fiber.Ctx) map[string]bool {
	permissionsMap := c.Locals(permissionsMapLocalKey)
	return permissionsMap.(map[string]bool)
}

func IsUserMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if IsUserSubject(c) {
			return c.Next()
		}
		return model.NewError(http.StatusForbidden).AddError("authorization", "invalid").Send(c)
	}
}

func IsServiceAccountMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if IsServiceAccount(c) {
			return c.Next()
		}
		return model.NewError(http.StatusForbidden).AddError("authorization", "invalid").Send(c)
	}
}
