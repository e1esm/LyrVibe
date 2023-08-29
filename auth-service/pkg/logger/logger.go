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
	logger         *zap.Logger
	once           sync.Once
	loggerFilePath = "log.json"
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		file, err := os.Create(loggerFilePath)
		if err != nil {
			log.Fatalf(err.Error())
		}
		setupLogger(file)
	})

	return logger
}

func setupLogger(fileWriter io.Writer) {
	config := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), zapcore.InfoLevel))
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
