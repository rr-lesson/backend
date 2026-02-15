package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ParseIncludes(c *fiber.Ctx) map[string]bool {
	result := make(map[string]bool)

	includes := c.Query("includes")
	if includes == "" {
		return result
	}

	items := strings.Split(includes, ",")

	for _, item := range items {
		result[strings.TrimSpace(item)] = true
	}

	return result
}
