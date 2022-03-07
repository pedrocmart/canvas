package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// New create new logger
func New(service, env, version, lvl string) *Logger {
	var l Logger

	var logLevel zerolog.Level

	switch lvl {
	case "DEBUG":
		logLevel = zerolog.DebugLevel
	case "INFO":
		logLevel = zerolog.InfoLevel
	case "WARN":
		logLevel = zerolog.WarnLevel
	case "ERROR":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.DebugLevel
	}

	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(logLevel)

	l.Log = zerolog.New(os.Stdout).
		With().Timestamp().Logger()

	// additional log variables
	l.Log = l.Log.
		With().Str("service", service).Logger()
	l.Log = l.Log.
		With().Str("env", env).Logger()
	l.Log = l.Log.
		With().Str("version", version).Logger()

	return &l
}

// Logger wrapper
type Logger struct {
	Log zerolog.Logger
}

// Infof sends info level event with formatted msg
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Log.Info().Msgf(format, args...)
}

// Info sends info level event
func (l *Logger) Info(args ...interface{}) {
	l.Log.Info().Msg(fmt.Sprint(args...))
}

// Debugf sends debug level event with formatted msg
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Log.Debug().Msgf(format, args...)
}

// Debug sends debug level event
func (l *Logger) Debug(args ...interface{}) {
	l.Log.Debug().Msg(fmt.Sprint(args...))
}

// Errorf sends error level event with formatted msg
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Log.Error().Msgf(format, args...)
}

// Error sends error level event
func (l *Logger) Error(args ...interface{}) {
	l.Log.Error().Msg(fmt.Sprint(args...))
}

// Warnf sends warn level event with formatted msg
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Log.Warn().Msgf(format, args...)
}

// Warn sends warn level event
func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn().Msg(fmt.Sprint(args...))
}

// Fatalf sends fatal level event with formatted msg
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Log.Fatal().Msgf(format, args...)
}

// Fatal sends warn level event
func (l *Logger) Fatal(args ...interface{}) {
	l.Log.Fatal().Msg(fmt.Sprint(args...))
}
