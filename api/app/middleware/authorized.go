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
		_, tokenString := GetAuthorizationFromContext(c)
		unvalidatedClaims, err := jwt.ParseClaimsFromTokenNoValidation(tokenString)
		if err != nil {
			log.Printf("failed to get authorization header: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		if unvalidatedClaims.Type != jwt.BearerTokenType {
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		tenent, err := repository.GetTenentByClientId(unvalidatedClaims.ClientId)
		if err != nil {
			log.Printf("failed to fetch application tenent: %v", err)
			return model.NewError(http.StatusUnauthorized).AddError("authorization", "invalid")
		}
		claims, err := jwt.ParseClaimsFromToken[jwt.Claims](tokenString, tenent)
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
		case jwt.ServiceAccountSubject:
			serviceAccount, err := repository.GetServiceAccountById(claims.Subject)
			if err != nil {
				log.Printf("failed to fetch service account: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			permissions, err := repository.GetServiceAccountPermissions(serviceAccount.Id)
			if err != nil {
				log.Printf("failed to fetch user permissions: %v", err)
				return model.NewError(http.StatusInternalServerError).AddError("internal", "application")
			}
			c.Locals(permissionsLocalKey, permissions)
			c.Locals(permissionsMapLocalKey, PermissionsFromRows(permissions))
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
		return model.NewError(http.StatusUnauthorized).AddError("token", "invalid")
	}
}

func GetClaims[C any](c *fiber.Ctx) *C {
	claims := c.Locals(baseClaimsLocalKey)
	return claims.(*C)
}

func HasScope(c *fiber.Ctx, scope ...string) bool {
	claims := GetClaims[jwt.Claims](c)
	for _, s := range scope {
		if !slices.Contains(claims.Scope, s) {
			return false
		}
	}
	return true
}

func IsUserSubject(c *fiber.Ctx) bool {
	return GetClaims[jwt.Claims](c).SubjectType == jwt.UserSubject
}

func GetUser(c *fiber.Ctx) *repository.UserRowST {
	user := c.Locals(userLocalKey)
	return user.(*repository.UserRowST)
}

func IsServiceAccount(c *fiber.Ctx) bool {
	return GetClaims[jwt.Claims](c).SubjectType == jwt.ServiceAccountSubject
}

func GetServiceAccount(c *fiber.Ctx) *repository.ServiceAccountRowST {
	serviceAccount := c.Locals(serviceAccountLocalKey)
	return serviceAccount.(*repository.ServiceAccountRowST)
}

func GetPermissions(c *fiber.Ctx) []repository.PermissionRowST {
	permissions := c.Locals(permissionsLocalKey)
	return permissions.([]repository.PermissionRowST)
}

func GetPermissionsMap(c *fiber.Ctx) map[string][]string {
	return c.Locals(permissionsMapLocalKey).(map[string][]string)
}

func IsUserMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if IsUserSubject(c) {
			return c.Next()
		}
		return model.NewError(http.StatusForbidden).AddError("authorization", "invalid")
	}
}

func IsServiceAccountMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if IsServiceAccount(c) {
			return c.Next()
		}
		return model.NewError(http.StatusForbidden).AddError("authorization", "invalid")
	}
}

type PermissionsST = map[string][]string

func PermissionsFromRows(rows []repository.PermissionRowST) PermissionsST {
	permissionsMap := make(PermissionsST)
	for _, row := range rows {
		if actions, ok := permissionsMap[row.Resource]; ok {
			permissionsMap[row.Resource] = append(actions, row.Actions...)
		} else {
			permissionsMap[row.Resource] = row.Actions
		}
	}
	return permissionsMap
}
