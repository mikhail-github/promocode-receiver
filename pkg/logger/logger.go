package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Init function provides log configuration
func Init(level string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}
	log.Debugf("Log level: %s", level)
}
