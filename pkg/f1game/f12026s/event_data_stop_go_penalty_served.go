package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataStopGoPenaltyServed struct {
	VehicleIdx uint8   // Vehicle index of the vehicle serving stop go
	StopTime   float32 // Time spent serving stop go in seconds
}

// GetVehicleIdx returns the VehicleIdx of *StopGoPenaltyServed
func (data *EventDataStopGoPenaltyServed) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *StopGoPenaltyServed
func (data *EventDataStopGoPenaltyServed) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

// GetStopTime returns the StopTime of *StopGoPenaltyServed
func (data *EventDataStopGoPenaltyServed) GetStopTime() float32 { return data.StopTime }

// SetStopTime stores the StopTime of *StopGoPenaltyServed
func (data *EventDataStopGoPenaltyServed) SetStopTime(v float32) { data.StopTime = v }

func (data *EventDataStopGoPenaltyServed) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.StopTime = reader.ReadFloat32()

}
