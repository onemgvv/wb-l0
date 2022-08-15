package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
)

func New(filename string) *zap.SugaredLogger {
	filePath, err := getFilePath(filename)
	if err != nil {
		log.Fatalf("[LOGGER || [ERROR]: %s", err.Error())
	}

	writeSyncer := getLogWriter(filePath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    2,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getFilePath(location string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(ex, location), nil
}
