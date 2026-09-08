package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type FinalClassificationData

type FinalClassificationData struct {
	Position          uint8                     // Finishing position
	NumLaps           uint8                     // Number of laps completed
	GridPosition      uint8                     // Grid position of the car
	Points            uint8                     // Number of points scored
	NumPitStops       uint8                     // Number of pit stops made
	ResultStatus      uint8                     // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
	ResultReason      uint8                     // Result reason - 0 = invalid, 1 = retired, 2 = finished, 3 = terminal damage, 4 = inactive, 5 = not enough laps completed, 6 = black flagged, 7 = red flagged, 8 = mechanical failure, 9 = session skipped, 10 = session simulated
	BestLapTimeInMS   uint32                    // Best lap time of the session in milliseconds
	TotalRaceTime     float64                   // Total race time in seconds without penalties
	PenaltiesTime     uint8                     // Total penalties accumulated in seconds
	NumPenalties      uint8                     // Number of penalties applied to this driver
	NumTyreStints     uint8                     // Number of tyres stints up to maximum
	TyreStintsActual  [CS_MAX_TYRE_STINTS]uint8 // Actual tyres used by this driver
	TyreStintsVisual  [CS_MAX_TYRE_STINTS]uint8 // Visual tyres used by this driver
	TyreStintsEndLaps [CS_MAX_TYRE_STINTS]uint8 // The lap number stints end on
}

func (data *FinalClassificationData) Parse(reader *xbinary.LittleEndianReader) {
	data.Position = reader.ReadUint8()
	data.NumLaps = reader.ReadUint8()
	data.GridPosition = reader.ReadUint8()
	data.Points = reader.ReadUint8()
	data.NumPitStops = reader.ReadUint8()
	data.ResultStatus = reader.ReadUint8()
	data.ResultReason = reader.ReadUint8()
	data.BestLapTimeInMS = reader.ReadUint32()
	data.TotalRaceTime = reader.ReadFloat64()
	data.PenaltiesTime = reader.ReadUint8()
	data.NumPenalties = reader.ReadUint8()
	data.NumTyreStints = reader.ReadUint8()
	reader.Read(data.TyreStintsActual[:])
	reader.Read(data.TyreStintsVisual[:])
	reader.Read(data.TyreStintsEndLaps[:])
}
