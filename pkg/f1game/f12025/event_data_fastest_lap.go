package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataFastestLap struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float32 // Lap time is in seconds
}

// GetVehicleIdx returns the VehicleIdx of *FastestLap
func (data *EventDataFastestLap) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *FastestLap
func (data *EventDataFastestLap) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

// GetLapTime returns the LapTime of *FastestLap
func (data *EventDataFastestLap) GetLapTime() float32 { return data.LapTime }

// SetLapTime stores the LapTime of *FastestLap
func (data *EventDataFastestLap) SetLapTime(v float32) { data.LapTime = v }

func (data *EventDataFastestLap) Parse(reader *xbinary.LittleEndianReader) {
	data.VehicleIdx = reader.ReadUint8()
	data.LapTime = reader.ReadFloat32()

}
