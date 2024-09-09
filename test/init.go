package test

import (
	"log"
	"os"

	"github.com/aikuci/go-subdivisions-id/internal/config"
	"github.com/aikuci/go-subdivisions-id/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var zapLog *zap.Logger

func init() {
	viperConfig = config.NewViper()
	zapLog = config.NewZapLog(viperConfig)
	app = config.NewFiber(viperConfig, &config.AppOptions{LogWriter: os.Stdout})
	db = config.NewDatabase(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
		Log:    zapLog,
		DB:     db,
	})

	// initTables()
}

func initTables() {
	err := db.AutoMigrate(&entity.Province{}, &entity.City{}, &entity.District{}, &entity.Village{})
	if err != nil {
		log.Fatalf("failed to initialize tables: %v", err)
	}
}
