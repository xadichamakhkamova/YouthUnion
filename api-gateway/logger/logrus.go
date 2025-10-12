package logger 

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()

	//stdout/stderr
	log.SetOutput(os.Stdout)

	// JSON format: 
	log.SetFormatter(&logrus.JSONFormatter{})

	// Log level: Info yoki Debug
	log.SetLevel(logrus.InfoLevel)

	return log
}
