package app

import (
	"social/config"
	thread_handler_rest "social/services/thread/handler/rest"
	thread_repo_pg "social/services/thread/repo/pg"
	thread_usecase "social/services/thread/usecase"
	user_handler_rest "social/services/user/handler/rest"

	user_repo_pg "social/services/user/repo/pg"
	user_usecase "social/services/user/usecase"

	"github.com/gofiber/contrib/otelfiber"
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
	s.app.Use(otelfiber.Middleware("my-server"))

	// ====================== REPOSITORIES ======================
	userRepoPg := user_repo_pg.NewUserRepoPG(db)
	threadRepoPg := thread_repo_pg.NewThreadRepoPg(db)

	// ====================== USECASES ======================
	userUseCase := user_usecase.NewUserUseCase(userRepoPg)
	threadUseCase := thread_usecase.NewThreadUseCase(threadRepoPg)

	// ====================== HANDLERS ======================
	userHandlerRest := user_handler_rest.NewUserHandlerRest(s.app, userUseCase)
	threadHandlerRest := thread_handler_rest.NewThreadHandler(s.app, threadUseCase)

	// ====================== ROUTES ======================
	userHandlerRest.Run(db)
	threadHandlerRest.Run(db)

	return s.app.Listen(":3000")
}
