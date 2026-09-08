package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataSpeedTrap

type EventDataSpeedTrap struct {
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                      float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest  in this session
	FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest  in this session
}

func (data *EventDataSpeedTrap) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.Speed = reader.ReadFloat32()
	data.IsOverallFastestInSession = reader.ReadUint8()
	data.IsDriverFastestInSession = reader.ReadUint8()
	data.FastestVehicleIdxInSession = reader.ReadUint8()
	data.FastestSpeedInSession = reader.ReadFloat32()

}
