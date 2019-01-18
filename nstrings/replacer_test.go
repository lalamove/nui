package nstrings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplacers(t *testing.T) {
	t.Run(
		"Upper case replacer",
		func(t *testing.T) {
			var s = "test "
			var ss = ReplacerToUpper.Replace(s)
			require.Equal(t, "TEST ", ss)
		},
	)

	t.Run(
		"Lower case replacer",
		func(t *testing.T) {
			var s = "TEST "
			var ss = ReplacerToLower.Replace(s)
			require.Equal(t, "test ", ss)
		},
	)

	t.Run(
		"Trim replacer",
		func(t *testing.T) {
			var s = "test "
			var ss = ReplacerTrimSpace.Replace(s)
			require.Equal(t, "test", ss)
		},
	)
}
