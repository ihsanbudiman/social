package thread_handler_rest

import (
	"social/constant"
	"social/domain"
	"social/middlewares"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type threadHandler struct {
	// fiber app
	app *fiber.App

	usecase domain.ThreadUseCase
}

// Run implements domain.Handler
func (t *threadHandler) Run(db *gorm.DB) {
	threadRoute := t.app.Group("/thread")

	// ! use this route for testing purpose only
	// threadRoute.Post("/create", middlewares.GormTransaction(db), t.CreateThread)
	threadRoute.Post("/create", middlewares.JWT(), middlewares.GormTransaction(db), t.CreateThread)

	threadRoute.Get("/", t.GetThread)
	threadRoute.Get("/replies", t.GetReplies)

}

func (t *threadHandler) CreateThread(c *fiber.Ctx) error {

	// get trx
	trx := c.Locals("get_trx").(*gorm.DB)

	// recover
	defer func() {
		if r := recover(); r != nil {
			// rollback trx
			trx.Rollback()

			c.Status(constant.HTTPResponseInternalServerError).JSON(fiber.Map{
				"message": r,
				"code":    constant.HTTPResponseInternalServerError,
			})

		}
	}()

	user := c.Locals("user").(domain.User)
	// check if user is empty
	if user.ID == 0 {
		trx.Rollback()
		return c.Status(constant.HTTPResponseUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
			"code":    constant.HTTPResponseUnauthorized,
		})
	}

	ctx := c.Context()

	// get email and password from request
	req := CreateThreadRequest{}
	if err := c.BodyParser(&req); err != nil {
		trx.Rollback()
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	threadPhotos := []domain.ThreadPhoto{}

	err := PhotoSaver(c, &threadPhotos, "image1", "image2", "image3", "image4", "image5")
	if err != nil {
		trx.Rollback()
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	thread := domain.Thread{
		Content: req.Content,
		UserID:  user.ID,
	}

	if req.ReplyTo != 0 {
		thread.ReplyTo = null.IntFrom(int64(req.ReplyTo))
	}

	// call usecase
	thread, err = t.usecase.WithTx(ctx, trx).CreateThread(ctx, thread)
	if err != nil {
		trx.Rollback()
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	for i := 0; i < len(threadPhotos); i++ {
		photo := domain.ThreadPhoto{
			ThreadID: thread.ID,
			FileUrl:  threadPhotos[i].FileUrl,
		}

		err = t.usecase.WithTx(ctx, trx).InsertThreadPhoto(ctx, &photo)
		if err != nil {
			trx.Rollback()
			return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
				"message": err.Error(),
				"code":    constant.HTTPResponseBadRequest,
			})
		}

	}

	thread, err = t.usecase.WithTx(ctx, trx).GetThread(ctx, thread.ID)
	if err != nil {
		trx.Rollback()
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"data":    thread,
		"code":    constant.HTTPResponseOK,
	})
}

func (t *threadHandler) GetThread(c *fiber.Ctx) error {
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

	req := GetThreadRequest{}

	if err := c.QueryParser(&req); err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	thread, err := t.usecase.GetThread(ctx, req.ThreadID)
	if err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"data":    thread,
		"code":    constant.HTTPResponseOK,
	})
}

func (t *threadHandler) GetReplies(c *fiber.Ctx) error {
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

	req := GetRepliesRequest{}

	if err := c.QueryParser(&req); err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	replies, err := t.usecase.GetReplies(ctx, req.ThreadID, req.Page)
	if err != nil {
		return c.Status(constant.HTTPResponseBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    constant.HTTPResponseBadRequest,
		})
	}

	return c.Status(constant.HTTPResponseOK).JSON(fiber.Map{
		"message": "success",
		"data":    replies,
		"code":    constant.HTTPResponseOK,
	})
}

func NewThreadHandler(app *fiber.App, usecase domain.ThreadUseCase) domain.Handler {
	return &threadHandler{
		app:     app,
		usecase: usecase,
	}
}
