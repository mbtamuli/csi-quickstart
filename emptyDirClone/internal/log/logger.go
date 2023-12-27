package log

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
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

func GRPCOpts(debug bool, logger logr.Logger) []grpc.ServerOption {
	var opts []grpc.ServerOption

	if debug {
		opts = append(opts, grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(
				interceptorLogger(logger),
				[]logging.Option{
					logging.WithLogOnEvents(
						logging.StartCall,
						logging.FinishCall,
						logging.PayloadReceived,
						logging.PayloadSent,
					),
				}...),
		))
	} else {
		opts = append(opts, grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(
				interceptorLogger(logger),
				[]logging.Option{
					logging.WithLogOnEvents(
						logging.FinishCall,
					),
				}...),
		))
	}

	return opts
}

// interceptorLogger adapts logr logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l logr.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.WithName("gRPCLogger").WithValues(fields...)
		switch lvl {
		case logging.LevelDebug:
			l.V(int(logging.LevelDebug)).Info(msg)
		case logging.LevelInfo:
			l.V(int(logging.LevelInfo)).Info(msg)
		case logging.LevelWarn:
			l.V(int(logging.LevelWarn)).Info(msg)
		case logging.LevelError:
			l.V(int(logging.LevelError)).Info(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
