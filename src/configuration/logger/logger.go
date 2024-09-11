package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logFilePath := "logfile.log"
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de log: %v\n", err)
		return
	}
	defer file.Close()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"

	zapcore.NewJSONEncoder(encoderConfig)

	cfg := zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: encoderConfig,
		OutputPaths: []string{
			"stdout",
			logFilePath,
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	log, err = cfg.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		fmt.Printf("fail to create log: %v\n", err)
		return
	}
}

func Info(message string, tags ...zap.Field) {
	defer log.Sync()

	log.Info(message, tags...)
}

func Error(message string, err error, tags ...zap.Field) {
	defer log.Sync()

	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
}
