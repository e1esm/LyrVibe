package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"sync"
)

var (
	loggerFilePath = "log.json"
	logger         *zap.Logger
	once           sync.Once
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		file, err := os.Create(loggerFilePath)
		if err != nil {
			log.Fatalf(err.Error())
		}
		setUpLogger(file)
	})

	return logger
}

func setUpLogger(file io.Writer) {
	config := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.InfoLevel))
	logger = zap.New(core, zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel))
}
