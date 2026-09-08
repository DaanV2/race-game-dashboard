package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataPenalty

type EventDataPenalty struct {
	PenaltyType      uint8 // Penalty type – see Appendices
	InfringementType uint8 // Infringement type – see Appendices
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

func (data *EventDataPenalty) Parse(reader *xbinary.LittleEndianReader) {
	data.PenaltyType = reader.ReadUint8()
	data.InfringementType = reader.ReadUint8()
	data.VehicleIdx = reader.ReadUint8()
	data.OtherVehicleIdx = reader.ReadUint8()
	data.Time = reader.ReadUint8()
	data.LapNum = reader.ReadUint8()
	data.PlacesGained = reader.ReadUint8()

}
