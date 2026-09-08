package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataFlashback

type EventDataFlashback struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

func (data *EventDataFlashback) Parse(reader *xbinary.LittleEndianReader) {
	data.FlashbackFrameIdentifier = reader.ReadUint32()
	data.FlashbackSessionTime = reader.ReadFloat32()
}
