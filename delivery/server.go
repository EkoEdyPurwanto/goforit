package delivery

import (
	"fmt"
	"github.com/EkoEdyPurwanto/goforit/config"
	"github.com/EkoEdyPurwanto/goforit/delivery/controller"
	"github.com/EkoEdyPurwanto/goforit/delivery/middleware"
	"github.com/EkoEdyPurwanto/goforit/manager"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		UseCaseM manager.UseCaseManager
		Engine   *fiber.App
		Host     string
		Log      *logrus.Logger
	}
)

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.Engine.Listen(s.Host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.Engine.Use(middleware.Logging(s.Log))
}

func (s *Server) initControllers() {
	controller.NewUsersController(s.UseCaseM.UsersUseCase(), s.Engine).Route()
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
		UseCaseM: useCaseManager,
		Engine:   app,
		Host:     hostAndPort,
		Log:      logger,
	}
}
