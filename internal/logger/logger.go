package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger constructor.
func NewLogger() *logrus.Logger {
	log := logrus.New()
	formatter := &logrus.TextFormatter{FullTimestamp: true}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	return log
}
