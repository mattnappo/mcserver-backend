package common

import (
	"os"

	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
)

// NewLogger will create a new default loggo.Logger.
func NewLogger(context string) *loggo.Logger {
	// Create and setup a new logger
	logger := loggo.GetLogger(context)
	logger.SetLogLevel(loggo.DEBUG)
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr)) // Add colors

	return &logger
}
