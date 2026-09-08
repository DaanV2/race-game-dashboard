package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataSafetyCar

type EventDataSafetyCar struct {
	SafetyCarType uint8 // 0 = No Safety Car, 1 = Full Safety Car  2 = Virtual Safety Car, 3 = Formation Lap Safety Car
	EventType     uint8 // 0 = Deployed, 1 = Returning, 2 = Returned  3 = Resume Race
}

func (data *EventDataSafetyCar) Parse(reader *xbinary.LittleEndianReader) {
	data.SafetyCarType = reader.ReadUint8()
	data.EventType = reader.ReadUint8()

}
