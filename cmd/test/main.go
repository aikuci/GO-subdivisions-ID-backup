package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Mode string `yaml:"mode"`
	} `yaml:"app"`
	Web struct {
		Prefork bool `yaml:"prefork"`
		Port    int  `yaml:"port"`
	} `yaml:"app"`
	Log struct {
		Level int `yaml:"level"`
	} `yaml:"app"`
	Database struct {
		Dialect string `yaml:"dialect"`
		Dsn     string `yaml:"dsn"`
		Pool    struct {
			Idle     int `yaml:"idle"`
			Max      int `yaml:"max"`
			Lifetime int `yaml:"lifetime"`
		} `yaml:"pool"`
	} `yaml:"database"`
}

func main() {
	// https://stackoverflow.com/a/78709806

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config values: %+v\n", cfg)
}
