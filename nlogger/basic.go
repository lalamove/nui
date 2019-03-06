package nlogger

import (
	"fmt"
	"strconv"
	"strings"
)

// Basic logger interface provides a very basic logger with functions that prints
// messages corresponding to the log level picked
type Basic interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

// ToStructured lets you use a simple logger that conforms to the Basic
// interface to satisfy the needs of packages that use the Structured logger
func ToStructured(logger Basic) Structured {
	return &basicStructured{logger: logger}
}

// From here and below follows an implementation that wraps a basic logger into a
// private struct that conforms to the Structured logger interface. Some speed
// considerations have been made, but not a lot
type entry struct {
	builder strings.Builder
}

func (e *entry) format(key string, value string) {
	e.builder.WriteRune(' ')
	e.builder.WriteString(key)
	e.builder.WriteRune('=')
	e.builder.WriteString(value)
}

func (e *entry) String(key string, value string) {
	e.format(key, value)
}

func (e *entry) Int(key string, value int) {
	e.format(key, strconv.Itoa(value))
}

func (e *entry) Int64(key string, value int64) {
	e.format(key, fmt.Sprintf("%d", value))
}

func (e *entry) Float(key string, value float64) {
	e.format(key, fmt.Sprintf("%f", value))
}

func (e *entry) Bool(key string, value bool) {
	e.format(key, strconv.FormatBool(value))
}

func (e *entry) Err(key string, value error) {
	e.format(key, value.Error())
}

func (e *entry) ObjectFunc(key string, value EntryFunc) {
	e.builder.WriteRune(' ')
	e.builder.WriteString(key)
	e.builder.WriteString("={")
	value(e)
	e.builder.WriteString(" }")
}

type basicStructured struct {
	logger Basic
}

func (b *basicStructured) Debug(msg string) {
	b.logger.Debug(msg)
}

func (b *basicStructured) DebugWithFields(msg string, ef EntryFunc) {
	var e = &entry{}
	e.builder.WriteString(msg)
	ef(e)
	b.logger.Debug(e.builder.String())
}

func (b *basicStructured) Info(msg string) {
	b.logger.Info(msg)
}

func (b *basicStructured) InfoWithFields(msg string, ef EntryFunc) {
	var e = &entry{}
	e.builder.WriteString(msg)
	ef(e)
	b.logger.Info(e.builder.String())
}

func (b *basicStructured) Warn(msg string) {
	b.logger.Warn(msg)
}

func (b *basicStructured) WarnWithFields(msg string, ef EntryFunc) {
	var e = &entry{}
	e.builder.WriteString(msg)
	ef(e)
	b.logger.Warn(e.builder.String())
}

func (b *basicStructured) Error(msg string) {
	b.logger.Error(msg)
}

func (b *basicStructured) ErrorWithFields(msg string, ef EntryFunc) {
	var e = &entry{}
	e.builder.WriteString(msg)
	ef(e)
	b.logger.Error(e.builder.String())
}

func (b *basicStructured) Fatal(msg string) {
	b.logger.Fatal(msg)
}

func (b *basicStructured) FatalWithFields(msg string, ef EntryFunc) {
	var e = &entry{}
	e.builder.WriteString(msg)
	ef(e)
	b.logger.Fatal(e.builder.String())
}
