package nlogger

import (
	"context"
	"errors"
	"io"
	"log"
)

type loggerKey string

// LoggerKey is the key to access the logger in context
const LoggerKey loggerKey = "logger"

// ErrLoggerNotFoundInContext is the error when calling MustFromContext and the logger is not found
var ErrLoggerNotFoundInContext = errors.New("Logger not found in context")

// Logger is a generic Logger interface
type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type defaultLogger struct {
	*log.Logger
}

// New will return a default logger instance
func New(target io.Writer, prefix string) Logger {
	return &defaultLogger{log.New(target, prefix, log.LstdFlags)}
}

// Debug will print the message in debug level
func (dl *defaultLogger) Debug(msg string) {
	dl.Print(msg)
}

// Info will print the message in info level
func (dl *defaultLogger) Info(msg string) {
	dl.Print(msg)
}

// Warn will print the message in warning level
func (dl *defaultLogger) Warn(msg string) {
	dl.Print(msg)
}

// Error will print the message in error level
func (dl *defaultLogger) Error(msg string) {
	dl.Print(msg)
}

// Fatal will print the message in fatal level and kill the main process
func (dl *defaultLogger) Fatal(msg string) {
	dl.Logger.Fatal(msg)
}

// SetInContext sets the logger in a context and returns the new context
func SetInContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

// FromContext gets the logger from context.
// If context does not exist, it returns nil.
func FromContext(ctx context.Context) Logger {
	var v = ctx.Value(LoggerKey)
	if v == nil {
		return nil
	}
	return v.(Logger)
}

// MustFromContext gets the logger from context.
// If context does not exist, it panics with a ErrLoggerNotFoundInContext.
func MustFromContext(ctx context.Context) Logger {
	var v = ctx.Value(LoggerKey)
	if v == nil {
		panic(ErrLoggerNotFoundInContext)
	}
	return v.(Logger)
}
