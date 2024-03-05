package delivery

import (
	"fmt"
	"github.com/EkoEdyPurwanto/goforit/config"
	"github.com/EkoEdyPurwanto/goforit/delivery/controller"
	"github.com/EkoEdyPurwanto/goforit/delivery/middleware"
	"github.com/EkoEdyPurwanto/goforit/manager"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		useCaseM manager.UseCaseManager
		engine   *fiber.App
		host     string
		log      *logrus.Logger
	}
)

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Listen(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.Logging(s.log))
}

func (s *Server) initControllers() {
	controller.NewUsersController(s.useCaseM.UsersUseCase(), s.engine).Route()
}

// NewServer Constructor
func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	var (
		// instance repository
		repositoryManager = manager.NewRepositoryManager(infraManager)

		// instance useCase
		v              = validator.New()
		useCaseManager = manager.NewUseCaseManager(repositoryManager, v)

		hostAndPort = fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)

		logger = logrus.New()
		app    = fiber.New()
	)
	return &Server{
		useCaseM: useCaseManager,
		engine:   app,
		host:     hostAndPort,
		log:      logger,
	}
}
