package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetPosition returns the Position of *FinalClassificationData
func (data *FinalClassificationData) GetPosition() uint8 { return data.Position }

// SetPosition stores the Position of *FinalClassificationData
func (data *FinalClassificationData) SetPosition(v uint8) { data.Position = v }

// GetNumLaps returns the NumLaps of *FinalClassificationData
func (data *FinalClassificationData) GetNumLaps() uint8 { return data.NumLaps }

// SetNumLaps stores the NumLaps of *FinalClassificationData
func (data *FinalClassificationData) SetNumLaps(v uint8) { data.NumLaps = v }

// GetGridPosition returns the GridPosition of *FinalClassificationData
func (data *FinalClassificationData) GetGridPosition() uint8 { return data.GridPosition }

// SetGridPosition stores the GridPosition of *FinalClassificationData
func (data *FinalClassificationData) SetGridPosition(v uint8) { data.GridPosition = v }

// GetPoints returns the Points of *FinalClassificationData
func (data *FinalClassificationData) GetPoints() uint8 { return data.Points }

// SetPoints stores the Points of *FinalClassificationData
func (data *FinalClassificationData) SetPoints(v uint8) { data.Points = v }

// GetNumPitStops returns the NumPitStops of *FinalClassificationData
func (data *FinalClassificationData) GetNumPitStops() uint8 { return data.NumPitStops }

// SetNumPitStops stores the NumPitStops of *FinalClassificationData
func (data *FinalClassificationData) SetNumPitStops(v uint8) { data.NumPitStops = v }

// GetResultStatus returns the ResultStatus of *FinalClassificationData
func (data *FinalClassificationData) GetResultStatus() uint8 { return data.ResultStatus }

// SetResultStatus stores the ResultStatus of *FinalClassificationData
func (data *FinalClassificationData) SetResultStatus(v uint8) { data.ResultStatus = v }

// GetResultReason returns the ResultReason of *FinalClassificationData
func (data *FinalClassificationData) GetResultReason() uint8 { return data.ResultReason }

// SetResultReason stores the ResultReason of *FinalClassificationData
func (data *FinalClassificationData) SetResultReason(v uint8) { data.ResultReason = v }

// GetBestLapTimeInMS returns the BestLapTimeInMS of *FinalClassificationData
func (data *FinalClassificationData) GetBestLapTimeInMS() uint32 { return data.BestLapTimeInMS }

// SetBestLapTimeInMS stores the BestLapTimeInMS of *FinalClassificationData
func (data *FinalClassificationData) SetBestLapTimeInMS(v uint32) { data.BestLapTimeInMS = v }

// GetTotalRaceTime returns the TotalRaceTime of *FinalClassificationData
func (data *FinalClassificationData) GetTotalRaceTime() float64 { return data.TotalRaceTime }

// SetTotalRaceTime stores the TotalRaceTime of *FinalClassificationData
func (data *FinalClassificationData) SetTotalRaceTime(v float64) { data.TotalRaceTime = v }

// GetPenaltiesTime returns the PenaltiesTime of *FinalClassificationData
func (data *FinalClassificationData) GetPenaltiesTime() uint8 { return data.PenaltiesTime }

// SetPenaltiesTime stores the PenaltiesTime of *FinalClassificationData
func (data *FinalClassificationData) SetPenaltiesTime(v uint8) { data.PenaltiesTime = v }

// GetNumPenalties returns the NumPenalties of *FinalClassificationData
func (data *FinalClassificationData) GetNumPenalties() uint8 { return data.NumPenalties }

// SetNumPenalties stores the NumPenalties of *FinalClassificationData
func (data *FinalClassificationData) SetNumPenalties(v uint8) { data.NumPenalties = v }

// GetNumTyreStints returns the NumTyreStints of *FinalClassificationData
func (data *FinalClassificationData) GetNumTyreStints() uint8 { return data.NumTyreStints }

// SetNumTyreStints stores the NumTyreStints of *FinalClassificationData
func (data *FinalClassificationData) SetNumTyreStints(v uint8) { data.NumTyreStints = v }

// GetTyreStintsActual returns the TyreStintsActual of *FinalClassificationData
func (data *FinalClassificationData) GetTyreStintsActual() [CS_MAX_TYRE_STINTS]uint8 {
	return data.TyreStintsActual
}

// SetTyreStintsActual stores the TyreStintsActual of *FinalClassificationData
func (data *FinalClassificationData) SetTyreStintsActual(v [CS_MAX_TYRE_STINTS]uint8) {
	data.TyreStintsActual = v
}

// GetTyreStintsVisual returns the TyreStintsVisual of *FinalClassificationData
func (data *FinalClassificationData) GetTyreStintsVisual() [CS_MAX_TYRE_STINTS]uint8 {
	return data.TyreStintsVisual
}

// SetTyreStintsVisual stores the TyreStintsVisual of *FinalClassificationData
func (data *FinalClassificationData) SetTyreStintsVisual(v [CS_MAX_TYRE_STINTS]uint8) {
	data.TyreStintsVisual = v
}

// GetTyreStintsEndLaps returns the TyreStintsEndLaps of *FinalClassificationData
func (data *FinalClassificationData) GetTyreStintsEndLaps() [CS_MAX_TYRE_STINTS]uint8 {
	return data.TyreStintsEndLaps
}

// SetTyreStintsEndLaps stores the TyreStintsEndLaps of *FinalClassificationData
func (data *FinalClassificationData) SetTyreStintsEndLaps(v [CS_MAX_TYRE_STINTS]uint8) {
	data.TyreStintsEndLaps = v
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
