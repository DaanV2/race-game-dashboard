package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataDRSDisabled

type EventDataDRSDisabled struct {
	Reason uint8 // 0 = Wet track, 1 = Safety car deployed, 2 = Red flag  3 = Min lap not reached
}

func (data *EventDataDRSDisabled) Parse(reader *xbinary.LittleEndianReader) {
	data.Reason = reader.ReadUint8()

}
