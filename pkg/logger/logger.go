package logger

import (
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type loggerImpl struct {
	loggerClient *log.Logger
}

func NewLogger() Logger {
	return loggerImpl{
		loggerClient: log.New(),
	}
}

func (logger loggerImpl) Info(format string, args ...interface{}) {
	logger.loggerClient.Infof(format, args...)
}

func (logger loggerImpl) Warn(format string, args ...interface{}) {
	logger.loggerClient.Warnf(format, args...)
}

func (logger loggerImpl) Error(format string, args ...interface{}) {
	logger.loggerClient.Errorf(format, args...)
}
