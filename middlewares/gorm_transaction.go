package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GormTransaction is a middleware that wraps a request in a transaction
func GormTransaction(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start transaction
		trx := db.Begin()

		// Check for errors
		if trx.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": trx.Error.Error(),
				"code":    fiber.StatusInternalServerError,
			})
		}
		// Set transaction in context
		c.Locals("get_trx", trx)
		// Call next handler
		if err := c.Next(); err != nil {
			// Rollback transaction
			trx.Rollback()
			return err
		}
		// Commit transaction
		trx.Commit()
		return nil
	}
}
