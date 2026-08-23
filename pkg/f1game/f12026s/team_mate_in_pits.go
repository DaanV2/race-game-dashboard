package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataTeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of team mate
}

// GetVehicleIdx returns the VehicleIdx of *TeamMateInPits
func (data *EventDataTeamMateInPits) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *TeamMateInPits
func (data *EventDataTeamMateInPits) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

func (data *EventDataTeamMateInPits) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()

}
