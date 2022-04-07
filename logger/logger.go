package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// These consts are refering levels/severity from GCP
const (
	DEFAULT   = "Default"
	DEBUG     = "Debug"
	INFO      = "Info"
	NOTICE    = "Notice"
	WARNING   = "Warning"
	ERROR     = "Error"
	ALERT     = "Alert"
	CRITICAL  = "Critical"
	EMERGENCY = "Emergency"
)

var (
	log              *logrus.Logger
	useSeverityField bool
)

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	useSeverityField = true
}

// SetLevel altera o level do logger
func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	log.SetLevel(lvl)
}

// GetLevel recupera o level do logger
func GetLevel() logrus.Level {
	return log.GetLevel()
}

// DisableSeverityField it will do disable severity fields on log
func DisableSeverityField() {
	useSeverityField = false
}

// setFieldSeverity will set severity field when it was availble
func setFieldSeverity(lvl string) logrus.Fields {
	if useSeverityField {
		return logrus.Fields{"severity": lvl}
	} else {
		return logrus.Fields{}
	}
}

// Default show detail of log
func Default(args ...interface{}) {
	log.WithFields(setFieldSeverity(DEFAULT)).Info(args...)
}

// Notice show detail of log
func Notice(args ...interface{}) {
	log.WithFields(setFieldSeverity(NOTICE)).Info(args...)
}

// Warning show detail of log
func Warning(args ...interface{}) {
	log.WithFields(setFieldSeverity(WARNING)).Warning(args...)
}

// Alert show detail of errors
func Alert(args ...interface{}) {
	log.WithFields(setFieldSeverity(ALERT)).Error(args...)
}

// Critical show detail of log and generate a fatal
func Critical(args ...interface{}) {
	log.WithFields(setFieldSeverity(CRITICAL)).Fatal(args...)
}

// Emergency show detail of log and generate a panic
func Emergency(args ...interface{}) {
	log.WithFields(setFieldSeverity(EMERGENCY)).Panic(args...)
}

// Error exibe detalhes do erro
func Error(args ...interface{}) {
	log.WithFields(setFieldSeverity(ERROR)).Error(args...)
}

// Info exibe detalhes do log info
func Info(args ...interface{}) {
	log.WithFields(setFieldSeverity(INFO)).Info(args...)
}

// Debug exibe detalhes do log debug
func Debug(args ...interface{}) {
	log.WithFields(setFieldSeverity(DEBUG)).Debug(args...)
}

// Trace exibe detalhes do log trace
func Trace(args ...interface{}) {
	log.WithFields(setFieldSeverity(DEBUG)).Trace(args...)
}

// Fatal exibe detalhes do erro
func Fatal(args ...interface{}) {
	log.WithFields(setFieldSeverity(EMERGENCY)).Panic(args...)
}
