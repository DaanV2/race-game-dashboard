package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataTeamMateInPits

type EventDataTeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of team mate
}

func (data *EventDataTeamMateInPits) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()

}
