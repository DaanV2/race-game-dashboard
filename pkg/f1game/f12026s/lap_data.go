package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type LapData

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
