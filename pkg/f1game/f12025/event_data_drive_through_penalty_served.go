package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataDriveThroughPenaltyServed struct {
	VehicleIdx uint8 // Vehicle index of the vehicle serving drive through
}

// GetVehicleIdx returns the VehicleIdx of *DriveThroughPenaltyServed
func (data *EventDataDriveThroughPenaltyServed) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *DriveThroughPenaltyServed
func (data *EventDataDriveThroughPenaltyServed) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

func (data *EventDataDriveThroughPenaltyServed) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()

}
