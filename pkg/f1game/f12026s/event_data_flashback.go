package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataFlashback struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

// GetFlashbackFrameIdentifier returns the FlashbackFrameIdentifier of *Flashback
func (data *EventDataFlashback) GetFlashbackFrameIdentifier() uint32 {
	return data.FlashbackFrameIdentifier
}

// SetFlashbackFrameIdentifier stores the FlashbackFrameIdentifier of *Flashback
func (data *EventDataFlashback) SetFlashbackFrameIdentifier(v uint32) {
	data.FlashbackFrameIdentifier = v
}

// GetFlashbackSessionTime returns the FlashbackSessionTime of *Flashback
func (data *EventDataFlashback) GetFlashbackSessionTime() float32 { return data.FlashbackSessionTime }

// SetFlashbackSessionTime stores the FlashbackSessionTime of *Flashback
func (data *EventDataFlashback) SetFlashbackSessionTime(v float32) { data.FlashbackSessionTime = v }

func (data *EventDataFlashback) Parse(reader *xbinary.LittleEndianReader) {
	data.FlashbackFrameIdentifier = reader.ReadUint32()
	data.FlashbackSessionTime = reader.ReadFloat32()

}
