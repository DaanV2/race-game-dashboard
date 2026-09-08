package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataRaceWinner

type EventDataRaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

func (data *EventDataRaceWinner) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
}
