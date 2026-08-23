package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataSpeedTrap struct {
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
	Speed                      float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
	IsDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest  in this session
	FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest  in this session
}

// GetVehicleIdx returns the VehicleIdx of *SpeedTrap
func (data *EventDataSpeedTrap) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *SpeedTrap
func (data *EventDataSpeedTrap) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

// GetSpeed returns the Speed of *SpeedTrap
func (data *EventDataSpeedTrap) GetSpeed() float32 { return data.Speed }

// SetSpeed stores the Speed of *SpeedTrap
func (data *EventDataSpeedTrap) SetSpeed(v float32) { data.Speed = v }

// GetIsOverallFastestInSession returns the IsOverallFastestInSession of *SpeedTrap
func (data *EventDataSpeedTrap) GetIsOverallFastestInSession() uint8 {
	return data.IsOverallFastestInSession
}

// SetIsOverallFastestInSession stores the IsOverallFastestInSession of *SpeedTrap
func (data *EventDataSpeedTrap) SetIsOverallFastestInSession(v uint8) {
	data.IsOverallFastestInSession = v
}

// GetIsDriverFastestInSession returns the IsDriverFastestInSession of *SpeedTrap
func (data *EventDataSpeedTrap) GetIsDriverFastestInSession() uint8 {
	return data.IsDriverFastestInSession
}

// SetIsDriverFastestInSession stores the IsDriverFastestInSession of *SpeedTrap
func (data *EventDataSpeedTrap) SetIsDriverFastestInSession(v uint8) {
	data.IsDriverFastestInSession = v
}

// GetFastestVehicleIdxInSession returns the FastestVehicleIdxInSession of *SpeedTrap
func (data *EventDataSpeedTrap) GetFastestVehicleIdxInSession() uint8 {
	return data.FastestVehicleIdxInSession
}

// SetFastestVehicleIdxInSession stores the FastestVehicleIdxInSession of *SpeedTrap
func (data *EventDataSpeedTrap) SetFastestVehicleIdxInSession(v uint8) {
	data.FastestVehicleIdxInSession = v
}

// GetFastestSpeedInSession returns the FastestSpeedInSession of *SpeedTrap
func (data *EventDataSpeedTrap) GetFastestSpeedInSession() float32 { return data.FastestSpeedInSession }

// SetFastestSpeedInSession stores the FastestSpeedInSession of *SpeedTrap
func (data *EventDataSpeedTrap) SetFastestSpeedInSession(v float32) { data.FastestSpeedInSession = v }

func (data *EventDataSpeedTrap) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.Speed = reader.ReadFloat32()
	data.IsOverallFastestInSession = reader.ReadUint8()
	data.IsDriverFastestInSession = reader.ReadUint8()
	data.FastestVehicleIdxInSession = reader.ReadUint8()
	data.FastestSpeedInSession = reader.ReadFloat32()

}
