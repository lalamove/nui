package ntime

import "time"

// DefaultTimer is a Timer wrapping the `time` package
var DefaultTimer = Time{}

// Timer is an interface that provides wrappers to get times
// It is useful to mock time.Now()
type Timer interface {
	Now() time.Time
}

// Time is an implementation of Timer wrapping `time` standard package
type Time struct{}

// Now returns time.Now()
func (Time) Now() time.Time {
	return time.Now()
}
