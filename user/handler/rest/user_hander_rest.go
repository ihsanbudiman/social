package user_handler_rest

import (
	"social/constant"
	"social/domain"

	"github.com/gofiber/fiber/v2"
)

type userHandlerRest struct {
	// fiber app
	app *fiber.App

	usecase domain.UserUseCase
}

func NewUserHandlerRest(app *fiber.App, usecase domain.UserUseCase) domain.Handler {
	return &userHandlerRest{
		app:     app,
		usecase: usecase,
	}
}

// Run define need to run
func (u *userHandlerRest) Run() {
	userRoute := u.app.Group("/user")

	userRoute.Get("/login", u.Login)
}

func (u *userHandlerRest) Login(c *fiber.Ctx) error {

	// recover
	defer func() {
		if r := recover(); r != nil {
			c.Status(constant.HTTPResponseInternalServerError).JSON(fiber.Map{
				"message": r,
				"code":    constant.HTTPResponseInternalServerError,
			})
		}
	}()

	ctx := c.Context()

	// get email and password from request
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	// call usecase
	_, err := u.usecase.LoginByEmail(ctx, req.Email, req.Password)
	if err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	// return response
	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"code":    constant.HTTPResponseOK,
	})
}
