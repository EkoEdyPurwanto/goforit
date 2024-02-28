package delivery

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"goforit/config"
	"goforit/delivery/controller"
	"goforit/manager"
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
	err := s.engine.Use(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	//s.engine.Use(middleware.LogRequest(s.log))
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

	// instance repository
	repositoryManager := manager.NewRepositoryManager(infraManager)
	// instance usecase
	useCaseManager := manager.NewUseCaseManager(repositoryManager)

	hostAndPort := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	logger := logrus.New()

	app := fiber.New()
	return &Server{
		useCaseM: useCaseManager,
		engine:   app,
		host:     hostAndPort,
		log:      logger,
	}
}
