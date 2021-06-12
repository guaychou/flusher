package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guaychou/flusher/internal/server"
	"github.com/guaychou/flusher/pkg/auth"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	apps *fiber.App
}

func Init() *Server {
	app := fiber.New(server.FiberConfig())
	username, password := auth.GetAuthCredential()
	app.Use(basicauth.New(server.BasicAuthConfig(username, password)))
	app.Use(recover.New())
	app.Use(logger.New(server.FiberLoggerConfig()))
	api := app.Group("/api/v1")
	server.SetupRoutes(api)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(
		signalChannel,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL)
	go server.GracefullyShutdown(app, signalChannel)
	return &Server{
		apps: app,
	}
}

func (server *Server) Run() {
	port := getPort()
	myFigure := figure.NewFigure("Flusher", "roman", true)
	myFigure.Print()
	log.WithFields(log.Fields{
		"version": os.Getenv("FLUSHER_VERSION"),
		"port":    port,
	}).Info("Flusher is running")
	if err := server.apps.Listen(":" + port); err != nil {
		log.Panic(err)
	}
}

func getPort() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		return value
	}
	return "3000"
}
