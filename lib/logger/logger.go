package logger

import (
	"strings"

	"github.com/goodcodeguy/bottomline/config"
	"github.com/juju/loggo"
)

// New provides a logger to be used by the application
func New(module string) loggo.Logger {
	var logger loggo.Logger

	cfg := config.GetConfiguration()
	logger = loggo.GetLogger(module)
	logger.SetLogLevel(logLevel(cfg.LogLevel))

	return logger
}

func logLevel(cfgLogLevel string) loggo.Level {
	var logInfo loggo.Level
	switch strings.ToUpper(cfgLogLevel) {
	case "INFO":
		logInfo = loggo.INFO
	case "DEBUG":
		logInfo = loggo.DEBUG
	case "TRACE":
		logInfo = loggo.TRACE
	case "ERROR":
		logInfo = loggo.ERROR
	case "WARNING":
		logInfo = loggo.WARNING
	case "CRITICAL":
		logInfo = loggo.CRITICAL
	default:
		logInfo = loggo.CRITICAL
	}
	return logInfo
}
