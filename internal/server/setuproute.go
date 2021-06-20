package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guaychou/flusher/internal/flusher"
	"github.com/guaychou/flusher/internal/handler"
)

// SetupRoutes for fiber for flusher api
func SetupRoutes(api fiber.Router) {
	api.Get("/health", handler.HealthHanlder)
	api.Post("/flush", flusher.RedisFlushHandler)
	api.Post("/flushasync", flusher.RedisFlushAsyncHandler)
	api.Post("/flushdb", flusher.RedisFlushDBHandler)
	api.Post("/flushdbasync", flusher.RedisFlushDbAsyncHandler)
	api.Post("/getkey", flusher.RedisGetkeyHandler)
	api.Post("/allkey", flusher.RedisAllkeyHandler)
	api.Post("/setkey", flusher.RedisSetkeyHandler)
	api.Post("/getmaster", flusher.RedisSentinelGetMaster)
	api.Post("/delkey", flusher.RedisDelkeyHandler)
	api.Post("/prefixdelkey", flusher.RedisPrefixDelkeyHandler)
}
