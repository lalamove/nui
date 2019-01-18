package nio

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrReader(t *testing.T) {
	t.Run(
		"err reader default err",
		func(t *testing.T) {
			var errReader = ErrReader{}
			var _, err = errReader.Read(nil)
			require.Equal(t, errDefaultRead, err)
		},
	)
	t.Run(
		"err reader custom err",
		func(t *testing.T) {
			var customErr = errors.New("custom err")
			var errReader = ErrReader{Err: customErr}
			var _, err = errReader.Read(nil)
			require.Equal(t, customErr, err)
		},
	)
}

func TestErrWriter(t *testing.T) {
	t.Run(
		"err writer default err",
		func(t *testing.T) {
			var errWriter = ErrWriter{}
			var _, err = errWriter.Write(nil)
			require.Equal(t, errDefaultWrite, err)
		},
	)
	t.Run(
		"err writer custom err",
		func(t *testing.T) {
			var customErr = errors.New("custom err")
			var errWriter = ErrWriter{Err: customErr}
			var _, err = errWriter.Write(nil)
			require.Equal(t, customErr, err)
		},
	)
}
