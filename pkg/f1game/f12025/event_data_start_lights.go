package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataStartLights

type EventDataStartLights struct {
	NumLights uint8 // Number of lights showing
}

func (data *EventDataStartLights) Parse(reader *xbinary.LittleEndianReader) {
	data.NumLights = reader.ReadUint8()

}
