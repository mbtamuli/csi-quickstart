package log

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(level int, environment string) logr.Logger {
	var zc zap.Config
	if environment == "production" {
		zc = zap.NewProductionConfig()
	}
	if environment == "development" {
		zc = zap.NewDevelopmentConfig()
	}

	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(level))
	zapLog, err := zc.Build()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	return zapr.NewLogger(zapLog)
}
