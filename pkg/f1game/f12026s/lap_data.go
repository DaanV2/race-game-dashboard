package f12026s

type LapData struct {
	m_lastLapTimeInMS              uint32  // Last lap time in milliseconds
	m_currentLapTimeInMS           uint32  // Current time around the lap in milliseconds
	m_sector1TimeMSPart            uint16  // Sector 1 time milliseconds part
	m_sector1TimeMinutesPart       uint8   // Sector 1 whole minute part
	m_sector2TimeMSPart            uint16  // Sector 2 time milliseconds part
	m_sector2TimeMinutesPart       uint8   // Sector 2 whole minute part
	m_deltaToCarInFrontMSPart      uint16  // Time delta to car in front milliseconds part
	m_deltaToCarInFrontMinutesPart uint8   // Time delta to car in front whole minute part
	m_deltaToRaceLeaderMSPart      uint16  // Time delta to race leader milliseconds part
	m_deltaToRaceLeaderMinutesPart uint8   // Time delta to race leader whole minute part
	m_lapDistance                  float32 // Distance vehicle is around current lap in metres – could, be negative if line hasn’t been crossed yet
	m_totalDistance                float32 // Total distance travelled in session in metres – could, be negative if line hasn’t been crossed yet
	m_safetyCarDelta               float32 // Delta in seconds for safety car
	m_carPosition                  uint8   // Car race position
	m_currentLapNum                uint8   // Current lap number
	m_pitStatus                    uint8   // 0 = none, 1 = pitting, 2 = in pit area
	m_numPitStops                  uint8   // Number of pit stops taken in this race
	m_sector                       uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	m_currentLapInvalid            uint8   // Current lap invalid - 0 = valid, 1 = invalid
	m_penalties                    uint8   // Accumulated time penalties in seconds to be added
	m_totalWarnings                uint8   // Accumulated number of warnings issued
	m_cornerCuttingWarnings        uint8   // Accumulated number of corner cutting warnings issued
	m_numUnservedDriveThroughPens  uint8   // Num drive through pens left to serve
	m_numUnservedStopGoPens        uint8   // Num stop go pens left to serve
	m_gridPosition                 uint8   // Grid position the vehicle started the race in
	m_driverStatus                 uint8   // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	m_resultStatus                 uint8   // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified,6 = not classified, 7 = retired
	m_pitLaneTimerActive           uint8   // Pit lane timing, 0 = inactive, 1 = active
	m_pitLaneTimeInLaneInMS        uint16  // If active, the current time spent in the pit lane in ms
	m_pitStopTimerInMS             uint16  // Time of the actual pit stop in ms
	m_pitStopShouldServePen        uint8   // Whether the car should serve a penalty at this stop
	m_speedTrapFastestSpeed        float32 // Fastest speed through speed trap for this car in kmph
	m_speedTrapFastestLap          uint8   // Lap no the fastest speed was achieved, 255 = not set
}

type PacketLapData struct {
	m_header               PacketHeader // Header
	m_lapData              [24]LapData  // Lap data for all cars on track
	m_timeTrialPBCarIdx    uint8        // Index of Personal Best car in time trial (255 if invalid)
	m_timeTrialRivalCarIdx uint8        // Index of Rival car in time trial (255 if invalid)
}
