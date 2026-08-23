package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type Overtake struct {
	OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

// GetOvertakingVehicleIdx returns the OvertakingVehicleIdx of *Overtake
func (data *Overtake) GetOvertakingVehicleIdx() uint8 { return data.OvertakingVehicleIdx }

// SetOvertakingVehicleIdx stores the OvertakingVehicleIdx of *Overtake
func (data *Overtake) SetOvertakingVehicleIdx(v uint8) { data.OvertakingVehicleIdx = v }

// GetBeingOvertakenVehicleIdx returns the BeingOvertakenVehicleIdx of *Overtake
func (data *Overtake) GetBeingOvertakenVehicleIdx() uint8 { return data.BeingOvertakenVehicleIdx }

// SetBeingOvertakenVehicleIdx stores the BeingOvertakenVehicleIdx of *Overtake
func (data *Overtake) SetBeingOvertakenVehicleIdx(v uint8) { data.BeingOvertakenVehicleIdx = v }

func (data *Overtake) Parse(reader *xbinary.LittleEndianReader) {
	data.OvertakingVehicleIdx = reader.ReadUint8()
	data.BeingOvertakenVehicleIdx = reader.ReadUint8()

}
