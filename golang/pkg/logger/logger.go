package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logLevelEnvKey = "APP_LOG_LEVEL"

func Init() *zap.Logger {
	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "json"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	// set logger level
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(os.Getenv(logLevelEnvKey)))
	if err != nil { // if no env variable, set level info as a default
		logLevel = zapcore.DebugLevel
	}

	prodConfig.Level = zap.NewAtomicLevelAt(logLevel)
	logger, err := prodConfig.Build()
	if err != nil {
		log.Fatalln("build logger from config failed")
	}
	zap.ReplaceGlobals(logger)

	zap.L().Info("logger has been initiated", zap.String("level", logLevel.String()))

	return logger
}
