package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Logger *zap.Logger

func init() {
	file, err := os.Create("log.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	config := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.InfoLevel))
	Logger = zap.New(core)
}
