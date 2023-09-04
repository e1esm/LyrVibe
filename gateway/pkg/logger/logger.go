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
		setUpLogger(file)
	})

	return logger
}

func setUpLogger(file io.Writer) {
	config := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.InfoLevel))
	logger = zap.New(core, zap.AddStacktrace(zapcore.InfoLevel), zap.AddCaller())
}
