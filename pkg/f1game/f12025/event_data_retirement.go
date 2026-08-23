package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataRetirement struct {
	VehicleIdx uint8 // Vehicle index of car retiring
	Reason     uint8 // Reason - 0 = invalid, 1 = retired, 2 = finished  3 = terminal damage, 4 = inactive, 5 = not enough laps completed  6 = black flagged, 7 = red flagged, 8 = mechanical failure  9 = session skipped, 10 = session simulated
}

// GetVehicleIdx returns the VehicleIdx of *Retirement
func (data *EventDataRetirement) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *Retirement
func (data *EventDataRetirement) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

// GetReason returns the Reason of *Retirement
func (data *EventDataRetirement) GetReason() uint8 { return data.Reason }

// SetReason stores the Reason of *Retirement
func (data *EventDataRetirement) SetReason(v uint8) { data.Reason = v }

func (data *EventDataRetirement) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.Reason = reader.ReadUint8()

}
