package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataStopGoPenaltyServed

type EventDataStopGoPenaltyServed struct {
	VehicleIdx uint8 // Vehicle index of the vehicle serving stop go
}

func (data *EventDataStopGoPenaltyServed) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
}
