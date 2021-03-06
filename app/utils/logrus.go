package utils

import (
	"io"
	"strings"

	"github.com/sirupsen/logrus"
)

// InitLogrus set logrus options
// logLevel set level which will be logged, values: debug, info(default), warning, error, fatal, panic
// logFormatter set log format, values: text(default), json
func InitLogrus(output io.Writer, logLevel string, logFormatter string) {
	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logrus.SetLevel(level)
	}

	if strings.ToLower(logFormatter) == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	logrus.SetOutput(output)
}
