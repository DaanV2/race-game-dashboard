package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataFastestLap

type EventDataFastestLap struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

func (data *EventDataFastestLap) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.LapTime = reader.ReadFloat32()

}
