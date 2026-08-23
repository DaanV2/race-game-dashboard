package xstrings_test

import (
	"testing"

	xstrings "github.com/daanv2/race-game-dashboard/pkg/extensions/strings"
	"github.com/stretchr/testify/require"
)

func Test_NullTerminated(t *testing.T) {
	v := xstrings.NullTerminated([]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 0, 0, 0, 0})
	require.Equal(t, "0123456789", v)

	v = xstrings.NullTerminated([]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	require.Equal(t, "0123456789", v)
}
