package middleware

import (
	"github.com/autumnleaf-ra/go-blogger-app/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	c.Next()
}
