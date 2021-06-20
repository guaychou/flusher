package server

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

var whitelisted = map[string]bool{
	"/api/v1/health": true,
}

// FiberConfig is custom config that we use for this apps
func FiberConfig() fiber.Config {
	preforkmode, _ := strconv.ParseBool(os.Getenv("PREFORK_MODE"))
	log.Info("Initialize fiber app")
	fiberconfig := fiber.Config{
		Prefork:               preforkmode,
		ServerHeader:          "Fiber Global Digital Niaga",
		StrictRouting:         true,
		CaseSensitive:         true,
		DisableStartupMessage: true,
		ProxyHeader:           "X-Real-IP",
		ErrorHandler:          handleError,
	}
	return fiberconfig
}

// BasicAuthConfig is auth middlewar config that set by vault
func BasicAuthConfig(username string, password string) basicauth.Config {
	log.Info("setting Up the basic auth credential from vault")
	authconfig := basicauth.Config{
		Users: map[string]string{
			username: password,
		},
		Next: WhiteListPath,
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(&fiber.Map{
				"status": false,
				"message":  "Oops, you are unauthorized",
			})
		},
	}
	return authconfig
}

// FiberLoggerConfig internal for internal use of logging structure in apps mimicking the logrus
func FiberLoggerConfig() logger.Config {
	log.Info("initialize logger config")
	loggerconfig := logger.Config{
		Format:     "time=\"${time}\" level=info from=${ip} code=${status} latency=${latency} method=${method} path=${path} pid=${pid} error=\"${error}\"\n",
		TimeFormat: "Mon, 02 Jan 2006 15:04:05",
	}
	return loggerconfig
}

//GracefullyShutdown is a method to do Graceful shutdown fiber and any other task
func GracefullyShutdown(app *fiber.App, signalChannel chan os.Signal) {
	<-signalChannel
	log.Info("gracefully shutting down")
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
	// in this section current webserver already killed , put your cleanup task like remove something here !
}

//WhiteListPath is a method than whitelist from basicauth
func WhiteListPath(c *fiber.Ctx) bool {
	return whitelisted[c.Path()]
}

func handleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retreive the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if err != nil {
		ctx.Status(code).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return nil
}
