package config

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapio"
)

func NewFiber(config *viper.Viper, zapLog *zap.Logger) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.GetBool("web.prefork"),
	})

	zapLogWriter := &zapio.Writer{Log: zapLog}
	defer zapLogWriter.Close()
	app.Use(logger.New(logger.Config{
		Output: zapLogWriter,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:                    10,
		SkipSuccessfulRequests: true,
	}))
	app.Use(recover.New(recover.Config{
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fmt.Println(c.Request().URI())
			stacks := fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack())
			log.Println(stacks)
		},
		EnableStackTrace: true,
	}))

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
