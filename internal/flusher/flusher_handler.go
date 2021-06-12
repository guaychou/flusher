package flusher

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	rds "github.com/guaychou/flusher/pkg/redis"
)

// RedisFlushHandler for redis flushall
func RedisFlushHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.Flushkey(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": respond,
	})
}

// RedisFlushAsyncHandler for redis flush all async
func RedisFlushAsyncHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}

	respond, err := rds.FlushKeyAsync(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": respond,
	})
}

// RedisFlusDBhHandler for redis flushdb
func RedisFlushDBHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.FlushDB(c.Context(), redis)
	if err != nil {
		return (err)
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": respond,
	})
}

// RedisFlusDbAsynchHandler for redis flushdbAsync
func RedisFlushDbAsyncHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	respond, err := rds.FlushDbAsync(c.Context(), redis)
	if err != nil {
		return (err)
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": respond,
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

// RedisSetkeyHandler for redis set key
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
func RedisDelkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	err := rds.Delkey(c.Context(), redis)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Key Deleted",
	})
}

// RedisPrefixDelkeyHandler for redis prefix del key
func RedisPrefixDelkeyHandler(c *fiber.Ctx) error {
	redis := &rds.Rds{}
	if err := c.BodyParser(redis); err != nil {
		return (err)
	}
	total_deleted_key, err := rds.Prefixdelkey(c.Context(), redis)
	if err != nil {
		return (err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Total deleted key: " + strconv.Itoa(int(total_deleted_key)),
	})
}
