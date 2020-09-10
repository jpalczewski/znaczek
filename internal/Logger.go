package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Vreate logger creates a logger.
func CreateLogger(enableDebug bool) *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if enableDebug == false {
		config.Level.SetLevel(zap.WarnLevel)
	}
	logger, _ := config.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}
