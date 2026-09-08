package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataRetirement

type EventDataRetirement struct {
	VehicleIdx uint8 // Vehicle index of car retiring
	Reason     uint8 // Reason - 0 = invalid, 1 = retired, 2 = finished  3 = terminal damage, 4 = inactive, 5 = not enough laps completed  6 = black flagged, 7 = red flagged, 8 = mechanical failure  9 = session skipped, 10 = session simulated
}

func (data *EventDataRetirement) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.Reason = reader.ReadUint8()

}
