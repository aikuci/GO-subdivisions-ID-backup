package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aikuci/go-subdivisions-id/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	zapLog := config.NewZapLog(viperConfig)
	db := config.NewDatabase(viperConfig, zapLog)
	validate := config.NewValidator(viperConfig)

	// Must be called within same function, check the [docs](https://go.dev/tour/flowcontrol/12)
	// Ref: https://www.reddit.com/r/golang/comments/bkqxgm/comment/emivhxl/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	file, err := os.OpenFile("storage/log/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app := config.NewFiber(viperConfig, &config.AppOptions{LogWriter: file})

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      zapLog,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
