package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataCollision

type EventDataCollision struct {
	Vehicle1Idx uint8 // Vehicle index of the first vehicle involved in the collision
	Vehicle2Idx uint8 // Vehicle index of the second vehicle involved in the collision
}

func (data *EventDataCollision) Parse(reader *xbinary.LittleEndianReader) {
	data.Vehicle1Idx = reader.ReadUint8()
	data.Vehicle2Idx = reader.ReadUint8()

}
