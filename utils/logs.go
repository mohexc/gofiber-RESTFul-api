package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func initLogger() {
	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func LogInfo(message string, field ...zap.Field) {
	log.Info(message, field...)
}

func LogDebug(message string, field ...zap.Field) {
	log.Debug(message, field...)
}

func LogWarn(message string, field ...zap.Field) {
	log.Warn(message, field...)
}

func LogError(message interface{}, field ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), field...)
	case string:
		log.Error(v, field...)
	}
}
