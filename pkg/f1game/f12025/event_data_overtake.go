package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataOvertake struct {
	OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

// GetOvertakingVehicleIdx returns the OvertakingVehicleIdx of *Overtake
func (data *EventDataOvertake) GetOvertakingVehicleIdx() uint8 { return data.OvertakingVehicleIdx }

// SetOvertakingVehicleIdx stores the OvertakingVehicleIdx of *Overtake
func (data *EventDataOvertake) SetOvertakingVehicleIdx(v uint8) { data.OvertakingVehicleIdx = v }

// GetBeingOvertakenVehicleIdx returns the BeingOvertakenVehicleIdx of *Overtake
func (data *EventDataOvertake) GetBeingOvertakenVehicleIdx() uint8 {
	return data.BeingOvertakenVehicleIdx
}

// SetBeingOvertakenVehicleIdx stores the BeingOvertakenVehicleIdx of *Overtake
func (data *EventDataOvertake) SetBeingOvertakenVehicleIdx(v uint8) {
	data.BeingOvertakenVehicleIdx = v
}

func (data *EventDataOvertake) Parse(reader *xbinary.LittleEndianReader) {
	data.OvertakingVehicleIdx = reader.ReadUint8()
	data.BeingOvertakenVehicleIdx = reader.ReadUint8()

}
