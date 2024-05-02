package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAuthorizationFromContext(c *fiber.Ctx) (string, string) {
	authorizationHeader := strings.TrimSpace(c.Get("Authorization"))
	if len(authorizationHeader) != 0 {
		parts := strings.SplitN(authorizationHeader, " ", 2)
		if len(parts) == 2 {
			tokenType := strings.TrimSpace(parts[0])
			token := strings.TrimSpace(parts[1])
			return tokenType, token
		}
	}
	return "", ""
}
