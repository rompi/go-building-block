package main

import (
	"building-block/logger"
	"bytes"
)

func main() {
	// call logger.NewLogger()
	var buf bytes.Buffer
	log := logger.NewLogger(&buf, nil)

	// call logger.Info()
	log.Info("This is an info message")
	// call logger.Warning()
	log.Warning("This is a warning message")
	// call logger.Error()
	log.Error("This is an error message")

}
