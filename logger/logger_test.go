package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// TestNewLogger tests the NewLogger function
func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, nil)
	assert.NotNil(t, logger)
	assert.IsType(t, &logrus.Logger{}, logger.logger)
}

// TestLoggerInfo tests the Info method
func TestLoggerInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, nil)

	logger.Info("This is an info message")
	assert.Contains(t, buf.String(), "This is an info message")
}

// TestLoggerWarning tests the Warning method
func TestLoggerWarning(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, nil)

	logger.Warning("This is a warning message")
	assert.Contains(t, buf.String(), "This is a warning message")
}

// TestLoggerError tests the Error method
func TestLoggerError(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, nil)

	logger.Error("This is an error message")
	assert.Contains(t, buf.String(), "This is an error message")
}

// TestNewLoggerWithFileOutput tests the NewLogger function with a file output
func TestNewLoggerWithFileOutput(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "logfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name()) // Clean up the file after the test

	logger := NewLogger(file, nil)
	assert.NotNil(t, logger)
	assert.IsType(t, &logrus.Logger{}, logger.logger)

	// Log a message
	logger.Info("This is an info message")

	// Close the file to flush the content
	file.Close()

	// Read the file content
	content, err := os.ReadFile(file.Name())
	assert.NoError(t, err)
	assert.Contains(t, string(content), "This is an info message")
}

// TestLoggerJSONOutput tests the logger with JSON formatter
func TestLoggerJSONOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, &logrus.JSONFormatter{})

	logger.Info("This is an info message in JSON format")
	assert.Contains(t, buf.String(), `msg="This is an info message in JSON format"`)
	assert.Contains(t, buf.String(), `level=info`)
}
