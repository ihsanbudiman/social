package app

import (
	"social/config"
	user_handler_rest "social/user/handler/rest"
	user_repo_pg "social/user/repo/pg"
	user_usecase "social/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) Start() error {

	db, err := config.NewGormConnection()
	if err != nil {
		panic(err)
	}

	userRepoPg := user_repo_pg.NewUserRepoPG(db)
	userUseCase := user_usecase.NewUserUseCase(userRepoPg)
	userHandlerRest := user_handler_rest.NewUserHandlerRest(s.app, userUseCase)
	userHandlerRest.Run()

	return s.app.Listen(":3000")
}
