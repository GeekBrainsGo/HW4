package main

import "github.com/sirupsen/logrus"

// NewLogger - Создаёт новый логгер
func NewLogger() *logrus.Logger {
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	lg.SetLevel(logrus.DebugLevel)
	return lg
}
