package user_handler_rest

import (
	"social/constant"
	"social/domain"
	"social/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
func (u *userHandlerRest) Run(db *gorm.DB) {
	userRoute := u.app.Group("/user")

	userRoute.Post("/login", u.Login)
	userRoute.Post("/register", u.Register)

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
	user, err := u.usecase.LoginByEmail(ctx, req.Email, req.Password)
	if err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	// remove password from user
	user.Password = ""

	jwt, err := helper.GenerateToken(user)
	if err != nil {
		return c.Status(constant.HTTPResponseInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseInternalServerError,
		})
	}

	// return response
	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"user":  user,
			"token": jwt,
		},
		"code": constant.HTTPResponseOK,
	})
}

func (u *userHandlerRest) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	// recover
	defer func() {
		if r := recover(); r != nil {
			c.Status(constant.HTTPResponseInternalServerError).JSON(fiber.Map{
				"message": r,
				"code":    constant.HTTPResponseInternalServerError,
			})
		}
	}()

	// get email and password from request
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	// call usecase
	user, err := u.usecase.Register(ctx, domain.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	// remove password from user so it won't be sent to client
	user.Password = ""

	jwt, err := helper.GenerateToken(user)
	if err != nil {
		return c.Status(constant.HTTPResponseInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseInternalServerError,
		})
	}

	// return response
	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"user":  user,
			"token": jwt,
		},
		"code": constant.HTTPResponseOK,
	})

}
