package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	logger         *zap.Logger
	once           sync.Once
	loggerFilePath = "log.json"
)

func newLogger() {
	file, err := os.Create(loggerFilePath)
	if err != nil {
		panic(err)
	}
	config := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.InfoLevel),
	)
	logger = zap.New(core, zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel))
}

func GetLogger() *zap.Logger {
	once.Do(func() {
		newLogger()
	})
	return logger
}
