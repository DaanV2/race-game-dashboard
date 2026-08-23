package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataCollision struct {
	Vehicle1Idx uint8 // Vehicle index of the first vehicle involved in the collision
	Vehicle2Idx uint8 // Vehicle index of the second vehicle involved in the collision
	Severity    uint8 // Severity of the collision - 0 = low, 1 = medium, 2 = high
}

// GetVehicle1Idx returns the Vehicle1Idx of *Collision
func (data *EventDataCollision) GetVehicle1Idx() uint8 { return data.Vehicle1Idx }

// SetVehicle1Idx stores the Vehicle1Idx of *Collision
func (data *EventDataCollision) SetVehicle1Idx(v uint8) { data.Vehicle1Idx = v }

// GetVehicle2Idx returns the Vehicle2Idx of *Collision
func (data *EventDataCollision) GetVehicle2Idx() uint8 { return data.Vehicle2Idx }

// SetVehicle2Idx stores the Vehicle2Idx of *Collision
func (data *EventDataCollision) SetVehicle2Idx(v uint8) { data.Vehicle2Idx = v }

// GetSeverity returns the Severity of *Collision
func (data *EventDataCollision) GetSeverity() uint8 { return data.Severity }

// SetSeverity stores the Severity of *Collision
func (data *EventDataCollision) SetSeverity(v uint8) { data.Severity = v }

func (data *EventDataCollision) Parse(reader *xbinary.LittleEndianReader) {
	data.Vehicle1Idx = reader.ReadUint8()
	data.Vehicle2Idx = reader.ReadUint8()
	data.Severity = reader.ReadUint8()

}
