package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataOvertake

type EventDataOvertake struct {
	OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

func (data *EventDataOvertake) Parse(reader *xbinary.LittleEndianReader) {
	data.OvertakingVehicleIdx = reader.ReadUint8()
	data.BeingOvertakenVehicleIdx = reader.ReadUint8()

}
