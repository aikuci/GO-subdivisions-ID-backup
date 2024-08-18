package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, zapLog *zap.Logger) *gorm.DB {
	dbFile := viper.GetString("database.file")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.New(&zapLogWriter{Logger: zapLog}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		zapLog.Fatal("failed to connect database:", zap.Error(err))
	}

	return db
}

type zapLogWriter struct {
	Logger *zap.Logger
}

func (l *zapLogWriter) Printf(message string, args ...interface{}) {
	l.Logger.Debug(message, zap.Field{Interface: args})
}
