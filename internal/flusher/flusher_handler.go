package flusher

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guaychou/flusher/internal/type"
	rds "github.com/guaychou/flusher/pkg/redis"
)

// RedisFlushHandler for redis flushall
// Flusher Flush Handler godoc
// @Summary Flush Handler
// @Description Flush handler for flusher
// @ID flush-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/flush [post]
// @securityDefinitions.basic BasicAuth
func RedisFlushHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.Flushkey(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: respond,
	})
}

// RedisFlushAsyncHandler for redis flush all async
// RedisFlushHandler for redis flushallasync
// Flusher FlushAsync Handler godoc
// @Summary FlushAsync Handler
// @Description Flush handler for flusher required redis version 4.0+
// @ID flush-async-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/flushasync [post]
// @securityDefinitions.basic BasicAuth
func RedisFlushAsyncHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.FlushKeyAsync(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: respond,
	})
}

// RedisFlusDBhHandler for redis flushdb
// Flusher FlushDB Handler godoc
// @Summary FlushDB Handler
// @Description FlushDB handler for flusher
// @ID flush-db-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/flushdb [post]
// @securityDefinitions.basic BasicAuth
func RedisFlushDBHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.FlushDB(c.Context(), redis)
	if err != nil {
		return (err)
	}
	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: respond,
	})
}

// RedisFlusDbAsynchHandler for redis flushdbAsync
// Flusher FlushDBAsync Handler godoc
// @Summary Flush DB async Handler
// @Description FlushDBasync handler for flusher
// @ID flush-db-async-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/flushdbasync [post]
// @securityDefinitions.basic BasicAuth
func RedisFlushDbAsyncHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.FlushDbAsync(c.Context(), redis)
	if err != nil {
		return (err)
	}
	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: respond,
	})
}

// RedisGetkeyHandler for redis get specific key
func RedisGetkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.Getkey(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"value":   respond,
	})
}

//RedisAllkeyHandler for redis get all key
func RedisAllkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.Allkeys(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"value":   respond,
	})
}

func RedisSetkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.Setkeys(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": respond,
	})
}

//RedisSentinelGetMaster is a handler that we use for returning Master address of sentinel
func RedisSentinelGetMaster(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.GetMasterName(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"value":   respond,
	})
}

// RedisDelkeyHandler for redis del key
// Flusher DelKey Handler godoc
// @Summary DelKey Handler
// @Description Delete handler for flusher
// @ID del-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/delkey [post]
// @securityDefinitions.basic BasicAuth
func RedisDelkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	err := rds.Delkey(c.Context(), redis)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: "Key Deleted",
	})
}

// RedisPrefixDelkeyHandler for redis prefix del key
// Flusher PrefixDelKey Handler godoc
// @Summary Flush by prefix/pattern Handler
// @Description Flush handler for flusher
// @ID prefix-del-key-handler
// @Accept  json
// @Produce  json
// @Param data body rds.Rds true "Redis Data"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /api/v1/flush [post]
// @securityDefinitions.basic BasicAuth
func RedisPrefixDelkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	total_deleted_key, err := rds.Prefixdelkey(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&types.SuccessResponse{
		Success: true,
		Message: "Total deleted key: " + strconv.Itoa(int(total_deleted_key)),
	})
}
