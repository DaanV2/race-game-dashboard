package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type LapData struct {
	LastLapTimeInMS              uint32  // Last lap time in milliseconds
	CurrentLapTimeInMS           uint32  // Current time around the lap in milliseconds
	Sector1TimeMSPart            uint16  // Sector 1 time milliseconds part
	Sector1TimeMinutesPart       uint8   // Sector 1 whole minute part
	Sector2TimeMSPart            uint16  // Sector 2 time milliseconds part
	Sector2TimeMinutesPart       uint8   // Sector 2 whole minute part
	DeltaToCarInFrontMSPart      uint16  // Time delta to car in front milliseconds part
	DeltaToCarInFrontMinutesPart uint8   // Time delta to car in front whole minute part
	DeltaToRaceLeaderMSPart      uint16  // Time delta to race leader milliseconds part
	DeltaToRaceLeaderMinutesPart uint8   // Time delta to race leader whole minute part
	LapDistance                  float32 // Distance vehicle is around current lap in metres – could  be negative if line hasn’t been crossed yet
	TotalDistance                float32 // Total distance travelled in session in metres – could  be negative if line hasn’t been crossed yet
	SafetyCarDelta               float32 // Delta in seconds for safety car
	CarPosition                  uint8   // Car race position
	CurrentLapNum                uint8   // Current lap number
	PitStatus                    uint8   // 0 = none, 1 = pitting, 2 = in pit area
	NumPitStops                  uint8   // Number of pit stops taken in this race
	Sector                       uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid            uint8   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                    uint8   // Accumulated time penalties in seconds to be added
	TotalWarnings                uint8   // Accumulated number of warnings issued
	CornerCuttingWarnings        uint8   // Accumulated number of corner cutting warnings issued
	NumUnservedDriveThroughPens  uint8   // Num drive through pens left to serve
	NumUnservedStopGoPens        uint8   // Num stop go pens left to serve
	GridPosition                 uint8   // Grid position the vehicle started the race in
	DriverStatus                 uint8   // Status of driver - 0 = in garage, 1 = flying lap  2 = in lap, 3 = out lap, 4 = on track
	ResultStatus                 uint8   // Result status - 0 = invalid, 1 = inactive, 2 = active  3 = finished, 4 = didnotfinish, 5 = disqualified  6 = not classified, 7 = retired
	PitLaneTimerActive           uint8   // Pit lane timing, 0 = inactive, 1 = active
	PitLaneTimeInLaneInMS        uint16  // If active, the current time spent in the pit lane in ms
	PitStopTimerInMS             uint16  // Time of the actual pit stop in ms
	PitStopShouldServePen        uint8   // Whether the car should serve a penalty at this stop
	SpeedTrapFastestSpeed        float32 // Fastest speed through speed trap for this car in kmph
	SpeedTrapFastestLap          uint8   // Lap no the fastest speed was achieved, 255 = not set
}

// GetLastLapTimeInMS returns the LastLapTimeInMS of *LapData
func (data *LapData) GetLastLapTimeInMS() uint32 { return data.LastLapTimeInMS }

// SetLastLapTimeInMS stores the LastLapTimeInMS of *LapData
func (data *LapData) SetLastLapTimeInMS(v uint32) { data.LastLapTimeInMS = v }

// GetCurrentLapTimeInMS returns the CurrentLapTimeInMS of *LapData
func (data *LapData) GetCurrentLapTimeInMS() uint32 { return data.CurrentLapTimeInMS }

// SetCurrentLapTimeInMS stores the CurrentLapTimeInMS of *LapData
func (data *LapData) SetCurrentLapTimeInMS(v uint32) { data.CurrentLapTimeInMS = v }

// GetSector1TimeMSPart returns the Sector1TimeMSPart of *LapData
func (data *LapData) GetSector1TimeMSPart() uint16 { return data.Sector1TimeMSPart }

// SetSector1TimeMSPart stores the Sector1TimeMSPart of *LapData
func (data *LapData) SetSector1TimeMSPart(v uint16) { data.Sector1TimeMSPart = v }

// GetSector1TimeMinutesPart returns the Sector1TimeMinutesPart of *LapData
func (data *LapData) GetSector1TimeMinutesPart() uint8 { return data.Sector1TimeMinutesPart }

// SetSector1TimeMinutesPart stores the Sector1TimeMinutesPart of *LapData
func (data *LapData) SetSector1TimeMinutesPart(v uint8) { data.Sector1TimeMinutesPart = v }

// GetSector2TimeMSPart returns the Sector2TimeMSPart of *LapData
func (data *LapData) GetSector2TimeMSPart() uint16 { return data.Sector2TimeMSPart }

// SetSector2TimeMSPart stores the Sector2TimeMSPart of *LapData
func (data *LapData) SetSector2TimeMSPart(v uint16) { data.Sector2TimeMSPart = v }

// GetSector2TimeMinutesPart returns the Sector2TimeMinutesPart of *LapData
func (data *LapData) GetSector2TimeMinutesPart() uint8 { return data.Sector2TimeMinutesPart }

// SetSector2TimeMinutesPart stores the Sector2TimeMinutesPart of *LapData
func (data *LapData) SetSector2TimeMinutesPart(v uint8) { data.Sector2TimeMinutesPart = v }

// GetDeltaToCarInFrontMSPart returns the DeltaToCarInFrontMSPart of *LapData
func (data *LapData) GetDeltaToCarInFrontMSPart() uint16 { return data.DeltaToCarInFrontMSPart }

// SetDeltaToCarInFrontMSPart stores the DeltaToCarInFrontMSPart of *LapData
func (data *LapData) SetDeltaToCarInFrontMSPart(v uint16) { data.DeltaToCarInFrontMSPart = v }

// GetDeltaToCarInFrontMinutesPart returns the DeltaToCarInFrontMinutesPart of *LapData
func (data *LapData) GetDeltaToCarInFrontMinutesPart() uint8 {
	return data.DeltaToCarInFrontMinutesPart
}

// SetDeltaToCarInFrontMinutesPart stores the DeltaToCarInFrontMinutesPart of *LapData
func (data *LapData) SetDeltaToCarInFrontMinutesPart(v uint8) { data.DeltaToCarInFrontMinutesPart = v }

// GetDeltaToRaceLeaderMSPart returns the DeltaToRaceLeaderMSPart of *LapData
func (data *LapData) GetDeltaToRaceLeaderMSPart() uint16 { return data.DeltaToRaceLeaderMSPart }

// SetDeltaToRaceLeaderMSPart stores the DeltaToRaceLeaderMSPart of *LapData
func (data *LapData) SetDeltaToRaceLeaderMSPart(v uint16) { data.DeltaToRaceLeaderMSPart = v }

// GetDeltaToRaceLeaderMinutesPart returns the DeltaToRaceLeaderMinutesPart of *LapData
func (data *LapData) GetDeltaToRaceLeaderMinutesPart() uint8 {
	return data.DeltaToRaceLeaderMinutesPart
}

// SetDeltaToRaceLeaderMinutesPart stores the DeltaToRaceLeaderMinutesPart of *LapData
func (data *LapData) SetDeltaToRaceLeaderMinutesPart(v uint8) { data.DeltaToRaceLeaderMinutesPart = v }

// GetLapDistance returns the LapDistance of *LapData
func (data *LapData) GetLapDistance() float32 { return data.LapDistance }

// SetLapDistance stores the LapDistance of *LapData
func (data *LapData) SetLapDistance(v float32) { data.LapDistance = v }

// GetTotalDistance returns the TotalDistance of *LapData
func (data *LapData) GetTotalDistance() float32 { return data.TotalDistance }

// SetTotalDistance stores the TotalDistance of *LapData
func (data *LapData) SetTotalDistance(v float32) { data.TotalDistance = v }

// GetSafetyCarDelta returns the SafetyCarDelta of *LapData
func (data *LapData) GetSafetyCarDelta() float32 { return data.SafetyCarDelta }

// SetSafetyCarDelta stores the SafetyCarDelta of *LapData
func (data *LapData) SetSafetyCarDelta(v float32) { data.SafetyCarDelta = v }

// GetCarPosition returns the CarPosition of *LapData
func (data *LapData) GetCarPosition() uint8 { return data.CarPosition }

// SetCarPosition stores the CarPosition of *LapData
func (data *LapData) SetCarPosition(v uint8) { data.CarPosition = v }

// GetCurrentLapNum returns the CurrentLapNum of *LapData
func (data *LapData) GetCurrentLapNum() uint8 { return data.CurrentLapNum }

// SetCurrentLapNum stores the CurrentLapNum of *LapData
func (data *LapData) SetCurrentLapNum(v uint8) { data.CurrentLapNum = v }

// GetPitStatus returns the PitStatus of *LapData
func (data *LapData) GetPitStatus() uint8 { return data.PitStatus }

// SetPitStatus stores the PitStatus of *LapData
func (data *LapData) SetPitStatus(v uint8) { data.PitStatus = v }

// GetNumPitStops returns the NumPitStops of *LapData
func (data *LapData) GetNumPitStops() uint8 { return data.NumPitStops }

// SetNumPitStops stores the NumPitStops of *LapData
func (data *LapData) SetNumPitStops(v uint8) { data.NumPitStops = v }

// GetSector returns the Sector of *LapData
func (data *LapData) GetSector() uint8 { return data.Sector }

// SetSector stores the Sector of *LapData
func (data *LapData) SetSector(v uint8) { data.Sector = v }

// GetCurrentLapInvalid returns the CurrentLapInvalid of *LapData
func (data *LapData) GetCurrentLapInvalid() uint8 { return data.CurrentLapInvalid }

// SetCurrentLapInvalid stores the CurrentLapInvalid of *LapData
func (data *LapData) SetCurrentLapInvalid(v uint8) { data.CurrentLapInvalid = v }

// GetPenalties returns the Penalties of *LapData
func (data *LapData) GetPenalties() uint8 { return data.Penalties }

// SetPenalties stores the Penalties of *LapData
func (data *LapData) SetPenalties(v uint8) { data.Penalties = v }

// GetTotalWarnings returns the TotalWarnings of *LapData
func (data *LapData) GetTotalWarnings() uint8 { return data.TotalWarnings }

// SetTotalWarnings stores the TotalWarnings of *LapData
func (data *LapData) SetTotalWarnings(v uint8) { data.TotalWarnings = v }

// GetCornerCuttingWarnings returns the CornerCuttingWarnings of *LapData
func (data *LapData) GetCornerCuttingWarnings() uint8 { return data.CornerCuttingWarnings }

// SetCornerCuttingWarnings stores the CornerCuttingWarnings of *LapData
func (data *LapData) SetCornerCuttingWarnings(v uint8) { data.CornerCuttingWarnings = v }

// GetNumUnservedDriveThroughPens returns the NumUnservedDriveThroughPens of *LapData
func (data *LapData) GetNumUnservedDriveThroughPens() uint8 { return data.NumUnservedDriveThroughPens }

// SetNumUnservedDriveThroughPens stores the NumUnservedDriveThroughPens of *LapData
func (data *LapData) SetNumUnservedDriveThroughPens(v uint8) { data.NumUnservedDriveThroughPens = v }

// GetNumUnservedStopGoPens returns the NumUnservedStopGoPens of *LapData
func (data *LapData) GetNumUnservedStopGoPens() uint8 { return data.NumUnservedStopGoPens }

// SetNumUnservedStopGoPens stores the NumUnservedStopGoPens of *LapData
func (data *LapData) SetNumUnservedStopGoPens(v uint8) { data.NumUnservedStopGoPens = v }

// GetGridPosition returns the GridPosition of *LapData
func (data *LapData) GetGridPosition() uint8 { return data.GridPosition }

// SetGridPosition stores the GridPosition of *LapData
func (data *LapData) SetGridPosition(v uint8) { data.GridPosition = v }

// GetDriverStatus returns the DriverStatus of *LapData
func (data *LapData) GetDriverStatus() uint8 { return data.DriverStatus }

// SetDriverStatus stores the DriverStatus of *LapData
func (data *LapData) SetDriverStatus(v uint8) { data.DriverStatus = v }

// GetResultStatus returns the ResultStatus of *LapData
func (data *LapData) GetResultStatus() uint8 { return data.ResultStatus }

// SetResultStatus stores the ResultStatus of *LapData
func (data *LapData) SetResultStatus(v uint8) { data.ResultStatus = v }

// GetPitLaneTimerActive returns the PitLaneTimerActive of *LapData
func (data *LapData) GetPitLaneTimerActive() uint8 { return data.PitLaneTimerActive }

// SetPitLaneTimerActive stores the PitLaneTimerActive of *LapData
func (data *LapData) SetPitLaneTimerActive(v uint8) { data.PitLaneTimerActive = v }

// GetPitLaneTimeInLaneInMS returns the PitLaneTimeInLaneInMS of *LapData
func (data *LapData) GetPitLaneTimeInLaneInMS() uint16 { return data.PitLaneTimeInLaneInMS }

// SetPitLaneTimeInLaneInMS stores the PitLaneTimeInLaneInMS of *LapData
func (data *LapData) SetPitLaneTimeInLaneInMS(v uint16) { data.PitLaneTimeInLaneInMS = v }

// GetPitStopTimerInMS returns the PitStopTimerInMS of *LapData
func (data *LapData) GetPitStopTimerInMS() uint16 { return data.PitStopTimerInMS }

// SetPitStopTimerInMS stores the PitStopTimerInMS of *LapData
func (data *LapData) SetPitStopTimerInMS(v uint16) { data.PitStopTimerInMS = v }

// GetPitStopShouldServePen returns the PitStopShouldServePen of *LapData
func (data *LapData) GetPitStopShouldServePen() uint8 { return data.PitStopShouldServePen }

// SetPitStopShouldServePen stores the PitStopShouldServePen of *LapData
func (data *LapData) SetPitStopShouldServePen(v uint8) { data.PitStopShouldServePen = v }

// GetSpeedTrapFastestSpeed returns the SpeedTrapFastestSpeed of *LapData
func (data *LapData) GetSpeedTrapFastestSpeed() float32 { return data.SpeedTrapFastestSpeed }

// SetSpeedTrapFastestSpeed stores the SpeedTrapFastestSpeed of *LapData
func (data *LapData) SetSpeedTrapFastestSpeed(v float32) { data.SpeedTrapFastestSpeed = v }

// GetSpeedTrapFastestLap returns the SpeedTrapFastestLap of *LapData
func (data *LapData) GetSpeedTrapFastestLap() uint8 { return data.SpeedTrapFastestLap }

// SetSpeedTrapFastestLap stores the SpeedTrapFastestLap of *LapData
func (data *LapData) SetSpeedTrapFastestLap(v uint8) { data.SpeedTrapFastestLap = v }

func (data *LapData) Parse(reader *xbinary.LittleEndianReader) {
	data.LastLapTimeInMS = reader.ReadUint32()
	data.CurrentLapTimeInMS = reader.ReadUint32()
	data.Sector1TimeMSPart = reader.ReadUint16()
	data.Sector1TimeMinutesPart = reader.ReadUint8()
	data.Sector2TimeMSPart = reader.ReadUint16()
	data.Sector2TimeMinutesPart = reader.ReadUint8()
	data.DeltaToCarInFrontMSPart = reader.ReadUint16()
	data.DeltaToCarInFrontMinutesPart = reader.ReadUint8()
	data.DeltaToRaceLeaderMSPart = reader.ReadUint16()
	data.DeltaToRaceLeaderMinutesPart = reader.ReadUint8()
	data.LapDistance = reader.ReadFloat32()
	data.TotalDistance = reader.ReadFloat32()
	data.SafetyCarDelta = reader.ReadFloat32()
	data.CarPosition = reader.ReadUint8()
	data.CurrentLapNum = reader.ReadUint8()
	data.PitStatus = reader.ReadUint8()
	data.NumPitStops = reader.ReadUint8()
	data.Sector = reader.ReadUint8()
	data.CurrentLapInvalid = reader.ReadUint8()
	data.Penalties = reader.ReadUint8()
	data.TotalWarnings = reader.ReadUint8()
	data.CornerCuttingWarnings = reader.ReadUint8()
	data.NumUnservedDriveThroughPens = reader.ReadUint8()
	data.NumUnservedStopGoPens = reader.ReadUint8()
	data.GridPosition = reader.ReadUint8()
	data.DriverStatus = reader.ReadUint8()
	data.ResultStatus = reader.ReadUint8()
	data.PitLaneTimerActive = reader.ReadUint8()
	data.PitLaneTimeInLaneInMS = reader.ReadUint16()
	data.PitStopTimerInMS = reader.ReadUint16()
	data.PitStopShouldServePen = reader.ReadUint8()
	data.SpeedTrapFastestSpeed = reader.ReadFloat32()
	data.SpeedTrapFastestLap = reader.ReadUint8()

}
