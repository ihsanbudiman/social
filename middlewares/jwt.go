package middlewares

import (
	"social/constant"
	"social/helper"

	"github.com/gofiber/fiber/v2"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// get Auth header
		authHeader := c.Get("Authorization")

		// check if auth header is empty
		if authHeader == "" {
			return c.Status(constant.HTTPResponseUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization Header",
				"code":    constant.HTTPResponseUnauthorized,
			})
		}

		// check if auth header is valid
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(constant.HTTPResponseUnauthorized).JSON(fiber.Map{
				"message": "Invalid Authorization Header",
				"code":    constant.HTTPResponseUnauthorized,
			})
		}

		// get token from auth header
		token := authHeader[7:]

		// parse token
		claims, err := helper.ParseToken(token)
		if err != nil {
			return c.Status(constant.HTTPResponseUnauthorized).JSON(fiber.Map{
				"message": "Invalid Token",
				"code":    constant.HTTPResponseUnauthorized,
			})

		}

		// set user data in context
		c.Locals("user", claims.User)

		return c.Next()
	}
}
