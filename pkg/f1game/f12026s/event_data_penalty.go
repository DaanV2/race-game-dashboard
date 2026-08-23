package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataPenalty struct {
	PenaltyType      uint8 // Penalty type – see Appendices
	InfringementType uint8 // Infringement type – see Appendices
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

// GetPenaltyType returns the PenaltyType of *Penalty
func (data *EventDataPenalty) GetPenaltyType() uint8 { return data.PenaltyType }

// SetPenaltyType stores the PenaltyType of *Penalty
func (data *EventDataPenalty) SetPenaltyType(v uint8) { data.PenaltyType = v }

// GetInfringementType returns the InfringementType of *Penalty
func (data *EventDataPenalty) GetInfringementType() uint8 { return data.InfringementType }

// SetInfringementType stores the InfringementType of *Penalty
func (data *EventDataPenalty) SetInfringementType(v uint8) { data.InfringementType = v }

// GetVehicleIdx returns the VehicleIdx of *Penalty
func (data *EventDataPenalty) GetVehicleIdx() uint8 { return data.VehicleIdx }

// SetVehicleIdx stores the VehicleIdx of *Penalty
func (data *EventDataPenalty) SetVehicleIdx(v uint8) { data.VehicleIdx = v }

// GetOtherVehicleIdx returns the OtherVehicleIdx of *Penalty
func (data *EventDataPenalty) GetOtherVehicleIdx() uint8 { return data.OtherVehicleIdx }

// SetOtherVehicleIdx stores the OtherVehicleIdx of *Penalty
func (data *EventDataPenalty) SetOtherVehicleIdx(v uint8) { data.OtherVehicleIdx = v }

// GetTime returns the Time of *Penalty
func (data *EventDataPenalty) GetTime() uint8 { return data.Time }

// SetTime stores the Time of *Penalty
func (data *EventDataPenalty) SetTime(v uint8) { data.Time = v }

// GetLapNum returns the LapNum of *Penalty
func (data *EventDataPenalty) GetLapNum() uint8 { return data.LapNum }

// SetLapNum stores the LapNum of *Penalty
func (data *EventDataPenalty) SetLapNum(v uint8) { data.LapNum = v }

// GetPlacesGained returns the PlacesGained of *Penalty
func (data *EventDataPenalty) GetPlacesGained() uint8 { return data.PlacesGained }

// SetPlacesGained stores the PlacesGained of *Penalty
func (data *EventDataPenalty) SetPlacesGained(v uint8) { data.PlacesGained = v }

func (data *EventDataPenalty) Parse(reader *xbinary.LittleEndianReader) {
	data.PenaltyType = reader.ReadUint8()
	data.InfringementType = reader.ReadUint8()
	data.VehicleIdx = reader.ReadUint8()
	data.OtherVehicleIdx = reader.ReadUint8()
	data.Time = reader.ReadUint8()
	data.LapNum = reader.ReadUint8()
	data.PlacesGained = reader.ReadUint8()

}
