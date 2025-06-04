package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	context string
	entry   *logrus.Entry
}

func NewLogger(context string) *Logger {
	base := logrus.New()
	base.SetFormatter(&logrus.JSONFormatter{})
	base.SetLevel(logrus.DebugLevel)

	return &Logger{
		context: context,
		entry:   base.WithField("context", context),
	}
}

func (l *Logger) Info(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.entry.WithFields(logrus.Fields(fields[0])).Info(msg)
	} else {
		l.entry.Info(msg)
	}
}

func (l *Logger) Error(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.entry.WithFields(logrus.Fields(fields[0])).Error(msg)
	} else {
		l.entry.Error(msg)
	}
}

func (l *Logger) Debug(msg string, fields ...Fields) {
	if len(fields) > 0 {
		l.entry.WithFields(logrus.Fields(fields[0])).Debug(msg)
	} else {
		l.entry.Debug(msg)
	}
}
