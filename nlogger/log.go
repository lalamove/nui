package nlogger

import (
	"io"
	"log"
)

// Logger is a generic Logger interface
type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

// New will return a default logger instance
func New(target io.Writer, prefix string) Structured {
	// wrapedy wrap!
	return &basicStructured{&defaultLogger{log.New(target, prefix, log.LstdFlags)}}
}

// NewWithLog allows you to pass in a log.Logger which then gets snuggly wrapped in
// an interface that suits the nlogger.Structured interface.
func NewWithLog(log *log.Logger) Structured {
	return &basicStructured{&defaultLogger{log}}
}

type defaultLogger struct {
	*log.Logger
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
