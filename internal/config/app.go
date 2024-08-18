package config

import (
	"github.com/aikuci/go-subdivisions-id/internal/delivery/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *zap.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories

	// setup use cases

	// setup controllers

	routeConfig := route.RouteConfig{
		App: config.App,
	}
	routeConfig.Setup()
}
