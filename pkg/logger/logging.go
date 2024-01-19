package logging

import (
	"github.com/sirupsen/logrus"
)

type Logging struct {
	Config Config
	logger *logrus.Logger
}

func New(config Config) *Logging {
	logger := logrus.New()
	logger.SetOutput(config.Out)
	logger.SetFormatter(config.Formatter)

	return &Logging{config, logger}
}

func (l *Logging) Log(level uint32, code int, message string, detailed interface{}) {
	l.logger.WithFields(logrus.Fields{
		"detailed": detailed,
		"code":     code,
	}).Log(logrus.Level(level), message)
}

func (l *Logging) Trace(code int, message string, detailed interface{}) {
	l.Log(LevelTrace, code, message, detailed)
}

func (l *Logging) Debug(code int, message string, detailed interface{}) {
	l.Log(LevelDebug, code, message, detailed)
}

func (l *Logging) Info(code int, message string, detailed interface{}) {
	l.Log(LevelInfo, code, message, detailed)
}

func (l *Logging) Warn(code int, message string, detailed interface{}) {
	l.Log(LevelWarn, code, message, detailed)
}

func (l *Logging) Error(code int, message string, detailed interface{}) {
	l.Log(LevelError, code, message, detailed)
}

func (l *Logging) Fatal(code int, message string, detailed interface{}) {
	l.Log(LevelFatal, code, message, detailed)
}

func (l *Logging) Panic(code int, message string, detailed interface{}) {
	l.Log(LevelPanic, code, message, detailed)
}
