package app

import (
	"social/config"
	thread_repo_pg "social/services/thread/repo/pg"
	user_handler_rest "social/services/user/handler/rest"
	user_repo_pg "social/services/user/repo/pg"
	user_usecase "social/services/user/usecase"

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

	// ====================== REPOSITORIES ======================
	userRepoPg := user_repo_pg.NewUserRepoPG(db)
	thread_repo_pg.NewThreadRepoPg(db)

	// ====================== USECASES ======================
	userUseCase := user_usecase.NewUserUseCase(userRepoPg)

	// ====================== HANDLERS ======================
	userHandlerRest := user_handler_rest.NewUserHandlerRest(s.app, userUseCase)

	// ====================== ROUTES ======================
	userHandlerRest.Run()

	return s.app.Listen(":3000")
}
