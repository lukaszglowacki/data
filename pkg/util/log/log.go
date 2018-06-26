package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = os.Stdout
	log.SetLevel(logrus.FatalLevel)
}

func Get() *logrus.Logger {
	return log
}

func Info(args ...interface{}) {
	Get().Info(args)
}

func Warn(args ...interface{}) {
	Get().Warn(args)
}

func Error(args ...interface{}) {
	Get().Error(args)
}

func Fatal(args ...interface{}) {
	Get().Fatal(args)
}
