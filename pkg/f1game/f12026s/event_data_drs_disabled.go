package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataDRSDisabled struct {
	Reason uint8 // 0 = Wet track, 1 = Safety car deployed, 2 = Red flag  3 = Min lap not reached
}

// GetReason returns the Reason of *DRSDisabled
func (data *EventDataDRSDisabled) GetReason() uint8 { return data.Reason }

// SetReason stores the Reason of *DRSDisabled
func (data *EventDataDRSDisabled) SetReason(v uint8) { data.Reason = v }

func (data *EventDataDRSDisabled) Parse(reader *xbinary.LittleEndianReader) {
	data.Reason = reader.ReadUint8()

}
