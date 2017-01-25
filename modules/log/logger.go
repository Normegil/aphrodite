package log

import "github.com/Sirupsen/logrus"

func New(logLevel logrus.Level) logrus.FieldLogger {
	log := logrus.New()
	log.Level = logLevel
	log.Formatter = &logrus.TextFormatter{}
	return log
}
