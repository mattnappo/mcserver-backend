package api

import (
	"os"

	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
)

// newLogger will create a new default loggo.Logger.
func newLogger(context string) *loggo.Logger {
	// Create and setup a new logger
	logger := loggo.GetLogger(context)
	logger.SetLogLevel(loggo.INFO)
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr)) // Add colors

	return &logger
}
