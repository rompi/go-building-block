package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Logger is a custom logger that wraps the logrus.Logger.
type Logger struct {
	logger *logrus.Logger
}

// NewLogger initializes and returns a new Logger instance.
// It sets the output to the provided io.Writer and uses the provided logrus.Formatter.
// If no formatter is provided, it defaults to a text formatter with full timestamps.
func NewLogger(output io.Writer, formatter logrus.Formatter) *Logger {
	logger := logrus.New()
	logger.SetOutput(output)
	if formatter == nil {
		formatter = &logrus.TextFormatter{
			FullTimestamp: true,
		}
	}
	// logger.SetFormatter(formatter)
	logger.SetLevel(logrus.InfoLevel)

	return &Logger{logger: logger}
}

// Info logs informational messages.
// It accepts a string message as input and logs it with the Info level.
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

// Warning logs warning messages.
// It accepts a string message as input and logs it with the Warning level.
func (l *Logger) Warning(msg string) {
	l.logger.Warn(msg)
}

// Error logs error messages.
// It accepts a string message as input and logs it with the Error level.
func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}
