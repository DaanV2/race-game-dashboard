package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type EventDataPenalty

type EventDataPenalty struct {
	PenaltyType      f1common.PenaltyType        // Penalty type (uint8) – see Appendices
	InfringementType f1common.InfringementTypeID // Infringement type (uint8) – see Appendices
	VehicleIdx       uint8                       // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8                       // Vehicle index of the other car involved
	Time             uint8                       // Time gained, or time spent doing action in seconds
	LapNum           uint8                       // Lap the penalty occurred on
	PlacesGained     uint8                       // Number of places gained by this
}

func (data *EventDataPenalty) Parse(reader *xbinary.LittleEndianReader) {
	data.PenaltyType = f1common.PenaltyType(reader.ReadUint8())
	data.InfringementType = f1common.InfringementTypeID(reader.ReadUint8())
	data.VehicleIdx = reader.ReadUint8()
	data.OtherVehicleIdx = reader.ReadUint8()
	data.Time = reader.ReadUint8()
	data.LapNum = reader.ReadUint8()
	data.PlacesGained = reader.ReadUint8()
}
