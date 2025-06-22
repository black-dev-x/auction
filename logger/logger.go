package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logConfiguration := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		Development:      true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfiguration.Build()
}

func Info(msg string) {
	log.Info(msg)
	log.Sync()
}

func Debug(msg string) {
	log.Debug(msg)
	log.Sync()
}

func Error(msg string, err error) {
	log.Error(msg, zap.Error(err))
	log.Sync()
}

func Warn(msg string) {
	log.Warn(msg)
	log.Sync()
}
