package nlogger

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type loggerKey string

// LoggerKey is the key to access the logger in context
const LoggerKey loggerKey = "nlogger.Structured"

// ErrLoggerNotFoundInContext is the error when calling MustFromContext and the logger is not found
var ErrLoggerNotFoundInContext = errors.New("Logger not found in context")

// Entry represents a single log line and is what you use in the *WithFields callback.
// You would call any of the functions in the interface to add a keyed value to the log line.
type Entry interface {
	String(key string, value string) Entry
	Int(key string, value int) Entry
	Int64(key string, value int64) Entry
	Float(key string, value float64) Entry
	Bool(key string, value bool) Entry
	Err(key string, value error) Entry
}

// EntryFunc defines the callback that you implement when using the *WithFields function
type EntryFunc func(Entry) Entry

// Structured is a generic structured logger that allows you to add fields to log messages
type Structured interface {
	Debug(string)
	DebugWithFields(string, EntryFunc)
	Info(string)
	InfoWithFields(string, EntryFunc)
	Warn(string)
	WarnWithFields(string, EntryFunc)
	Error(string)
	ErrorWithFields(string, EntryFunc)
	Fatal(string)
	FatalWithFields(string, EntryFunc)
}

// SetInContext sets the logger in a context and returns the new context
func SetInContext(ctx context.Context, logger Structured) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

// FromContext gets the logger from context.
// If context does not exist, it returns nil.
func FromContext(ctx context.Context) Structured {
	var v = ctx.Value(LoggerKey)
	if v == nil {
		return nil
	}
	return v.(Structured)
}

// MustFromContext gets the logger from context.
// If context does not exist, it panics with a ErrLoggerNotFoundInContext.
func MustFromContext(ctx context.Context) Structured {
	var v = ctx.Value(LoggerKey)
	if v == nil {
		panic(ErrLoggerNotFoundInContext)
	}
	return v.(Structured)
}

// Provider is an interface to get a thread safe logger provider with
// the ability to replace the internal logger provided
type Provider interface {
	// Get returns the Provider's attached logger in a thread safe manner
	Get() Structured
	// Replace replaces the provider's internal logger in a thread safe manner
	Replace(Structured)
}

type provider struct {
	v   *atomic.Value
	mut *sync.Mutex
}

// NewProvider returns a new Structured Provider from the given Logger l
func NewProvider(l Structured) Provider {
	var v atomic.Value
	v.Store(l)

	return &provider{
		v:   &v,
		mut: &sync.Mutex{},
	}
}

// Get returns the attached Structured logger
func (s *provider) Get() Structured {
	return s.v.Load().(Structured)
}

// Replace the logger inside the Provider
func (s *provider) Replace(l Structured) {
	s.mut.Lock()
	s.v.Store(l)
	s.mut.Unlock()
}
