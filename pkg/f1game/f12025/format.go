package f12025

import xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"

// Returns a new byte reader wrapped around the buffer
func NewByteReader(buf []byte) *xbinary.LittleEndianReader {
	return xbinary.NewLittleEndianReader(buf)
}
