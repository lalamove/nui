package nio

import "errors"

var errDefaultWrite = errors.New("Write error")

// ErrWriter is a writer that returns an error instantly
// It is used for testing errors on io.Writers
type ErrWriter struct {
	Err error
}

// Write implements io.Writer
func (e ErrWriter) Write([]byte) (int, error) {
	if e.Err == nil {
		return 0, errDefaultWrite
	}
	return 0, e.Err
}
