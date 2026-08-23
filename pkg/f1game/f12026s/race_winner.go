package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataRaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

// GetVehicleIdx returns the VehicleIdx of *RaceWinner
func (data *EventDataRaceWinner) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *RaceWinner
func (data *EventDataRaceWinner) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

func (data *EventDataRaceWinner) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()

}
