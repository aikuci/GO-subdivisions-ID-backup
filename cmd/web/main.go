package main

import (
	"fmt"
	"log"

	"github.com/aikuci/go-subdivisions-id/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	zapLog := config.NewZapLog(viperConfig)
	db := config.NewDatabase(viperConfig, zapLog)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig, zapLog)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      zapLog,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
