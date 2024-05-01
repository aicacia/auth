package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aicacia/auth/api/app/model"
	"github.com/gofiber/fiber/v2"
)

var errParseLimitOffset = fmt.Errorf("failed to parse limit or offset")

func GetLimitAndOffset(c *fiber.Ctx, limit int) (int, int, error) {
	offset := 0

	limitString := c.Query("limit")
	if limitString != "" {
		integer, err := strconv.Atoi(limitString)
		if err != nil {
			log.Printf("failed to parse limit: %v\n", err)
			err = model.NewError(http.StatusUnauthorized).AddError("limit", "invalid").Send(c)
			if err == nil {
				err = errParseLimitOffset
			}
			return 0, 0, err
		}
		limit = integer
	}
	offsetString := c.Query("offset")
	if offsetString != "" {
		integer, err := strconv.Atoi(offsetString)
		if err != nil {
			log.Printf("failed to parse offset: %v\n", err)
			err = model.NewError(http.StatusUnauthorized).AddError("offset", "invalid").Send(c)
			if err == nil {
				err = errParseLimitOffset
			}
			return 0, 0, err
		}
		offset = integer
	}

	return limit, offset, nil
}
