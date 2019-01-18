package nio

import "errors"

var errDefaultRead = errors.New("Read error")

// ErrReader is a reader that returns an error instantly
// It is used for testing errors on io.Readers
type ErrReader struct {
	// Err is the error to return by the ErrReader
	// If nil, it returs a default error
	Err error
}

// Read implements the io.Reader interface
func (e ErrReader) Read([]byte) (int, error) {
	if e.Err == nil {
		return 0, errDefaultRead
	}
	return 0, e.Err
}
