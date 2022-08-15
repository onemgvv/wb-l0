package logger

import (
	"github.com/onemgvv/wb-l0/pkg/logger"
	"go.uber.org/zap"
	"path/filepath"
)

var (
	infoPath  = filepath.Join("logs", "info", "info.log")
	errorPath = filepath.Join("logs", "error", "error.log")

	InfoLogger  = New(infoPath)
	ErrorLogger = New(errorPath)
)

func New(path string) *zap.SugaredLogger {
	return logger.New(path)
}
