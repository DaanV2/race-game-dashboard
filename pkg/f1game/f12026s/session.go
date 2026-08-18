package f12026s

type MarshalZone struct {
	m_zoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	m_zoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
}

type ActiveAeroZone struct {
	m_zoneStart float32 // Fraction (0..1) of way through the lap the Active Aero zone starts
	m_zoneEnd   float32 // Fraction (0..1) of way through the lap the Active Aero zone ends
}

type DRSZone struct {
	m_zoneStart float32 // Fraction (0..1) of way through the lap the DRS zone starts
	m_zoneEnd   float32 // Fraction (0..1) of way through the lap the DRS zone ends
}

type WeatherForecastSample struct {
	m_sessionType            uint8 // 0 = unknown, see appendix
	m_timeOffset             uint8 // Time in minutes the forecast is for
	m_weather                uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast 3 = light rain, 4 = heavy rain, 5 = storm
	m_trackTemperature       int8  // Track temp. in degrees Celsius
	m_trackTemperatureChange int8  // Track temp. change – 0 = up, 1 = down, 2 = no change
	m_airTemperature         int8  // Air temp. in degrees celsius
	m_airTemperatureChange   int8  // Air temp. change – 0 = up, 1 = down, 2 = no change
	m_rainPercentage         uint8 // Percentage chance of rain (0-100)
}

type PacketSessionData struct {
	m_header              PacketHeader    // Header
	m_weather             uint8           // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	m_trackTemperature    int8            // Track temp. in degrees celsius
	m_airTemperature      int8            // Air temp. in degrees celsius
	m_totalLaps           uint8           // Total number of laps in this race
	m_trackLength         uint16          // Track length in metres
	m_sessionType         uint8           // 0 = unknown, see appendix
	m_trackId             int8            // -1 for unknown, see appendix
	m_formula             uint8           // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,  3 = F1 Generic, 4 = Beta, 6 = Esports, 8 = F1 World, 9 = F1 Elimination, 13 = F1 26
	m_sessionTimeLeft     uint16          // Time left in session in seconds
	m_sessionDuration     uint16          // Session duration in seconds
	m_pitSpeedLimit       uint8           // Pit speed limit in kilometres per hour
	m_gamePaused          uint8           // Whether the game is paused – network game only
	m_isSpectating        uint8           // Whether the player is spectating
	m_spectatorCarIndex   uint8           // Index of the car being spectated
	m_sliProNativeSupport uint8           // SLI Pro support, 0 = inactive, 1 = active
	m_numMarshalZones     uint8           // Number of marshal zones to follow
	m_marshalZones        [21]MarshalZone // List of marshal zones – max 21
	m_safetyCarStatus     uint8           // 0 = no safety car, 1 = full
	// 2 = virtual, 3 = formation lap
	m_networkGame                     uint8                     // 0 = offline, 1 = online
	m_numWeatherForecastSamples       uint8                     // Number of weather samples to follow
	m_weatherForecastSamples          [64]WeatherForecastSample // Array of weather forecast samples
	m_forecastAccuracy                uint8                     // 0 = Perfect, 1 = Approximate
	m_aiDifficulty                    uint8                     // AI Difficulty rating – 0-110
	m_seasonLinkIdentifier            uint32                    // Identifier for season - persists across saves
	m_weekendLinkIdentifier           uint32                    // Identifier for weekend - persists across saves
	m_sessionLinkIdentifier           uint32                    // Identifier for session - persists across saves
	m_pitStopWindowIdealLap           uint8                     // Ideal lap to pit on for current strategy (player)
	m_pitStopWindowLatestLap          uint8                     // Latest lap to pit on for current strategy (player)
	m_pitStopRejoinPosition           uint8                     // Predicted position to rejoin at (player)
	m_steeringAssist                  uint8                     // 0 = off, 1 = on
	m_brakingAssist                   uint8                     // 0 = off, 1 = low, 2 = medium, 3 = high
	m_gearboxAssist                   uint8                     // 1 = manual, 2 = manual & suggested gear, 3 = auto
	m_pitAssist                       uint8                     // 0 = off, 1 = on
	m_pitReleaseAssist                uint8                     // 0 = off, 1 = on
	m_ERSAssist                       uint8                     // 0 = off, 1 = on
	m_DRSAssist                       uint8                     // 0 = off, 1 = on
	m_dynamicRacingLine               uint8                     // 0 = off, 1 = corners only, 2 = full
	m_dynamicRacingLineType           uint8                     // 0 = 2D, 1 = 3D
	m_gameMode                        uint8                     // Game mode id - see appendix
	m_ruleSet                         uint8                     // Ruleset - see appendix
	m_timeOfDay                       uint32                    // Local time of day - minutes since midnight
	m_sessionLength                   uint8                     // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium, 5 = Medium Long, 6 = Long, 7 = Full
	m_speedUnitsLeadPlayer            uint8                     // 0 = MPH, 1 = KPH
	m_temperatureUnitsLeadPlayer      uint8                     // 0 = Celsius, 1 = Fahrenheit
	m_speedUnitsSecondaryPlayer       uint8                     // 0 = MPH, 1 = KPH
	m_temperatureUnitsSecondaryPlayer uint8                     // 0 = Celsius, 1 = Fahrenheit
	m_numSafetyCarPeriods             uint8                     // Number of safety cars called during session
	m_numVirtualSafetyCarPeriods      uint8                     // Number of virtual safety cars called
	m_numRedFlagPeriods               uint8                     // Number of red flags called during session
	m_equalCarPerformance             uint8                     // 0 = Off, 1 = On
	m_recoveryMode                    uint8                     // 0 = None, 1 = Flashbacks, 2 = Auto-recovery
	m_flashbackLimit                  uint8                     // 0 = Low, 1 = Medium, 2 = High, 3 = Unlimited
	m_surfaceType                     uint8                     // 0 = Simplified, 1 = Realistic
	m_lowFuelMode                     uint8                     // 0 = Easy, 1 = Hard
	m_raceStarts                      uint8                     // 0 = Manual, 1 = Assisted
	m_tyreTemperature                 uint8                     // 0 = Surface only, 1 = Surface & Carcass
	m_pitLaneTyreSim                  uint8                     // 0 = On, 1 = Off
	m_carDamage                       uint8                     // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Simulation
	m_carDamageRate                   uint8                     // 0 = Reduced, 1 = Standard, 2 = Simulation
	m_collisions                      uint8                     // 0 = Off, 1 = Player-to-Player Off, 2 = On
	m_collisionsOffForFirstLapOnly    uint8                     // 0 = Disabled, 1 = Enabled
	m_mpUnsafePitRelease              uint8                     // 0 = On, 1 = Off (Multiplayer)
	m_mpOffForGriefing                uint8                     // 0 = Disabled, 1 = Enabled (Multiplayer)
	m_cornerCuttingStringency         uint8                     // 0 = Regular, 1 = Strict
	m_parcFermeRules                  uint8                     // 0 = Off, 1 = On
	m_pitStopExperience               uint8                     // 0 = Automatic, 1 = Broadcast, 2 = Immersive
	m_safetyCar                       uint8                     // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	m_safetyCarExperience             uint8                     // 0 = Broadcast, 1 = Immersive
	m_formationLap                    uint8                     // 0 = Off, 1 = On
	m_formationLapExperience          uint8                     // 0 = Broadcast, 1 = Immersive
	m_redFlags                        uint8                     // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	m_affectsLicenceLevelSolo         uint8                     // 0 = Off, 1 = On
	m_affectsLicenceLevelMP           uint8                     // 0 = Off, 1 = On
	m_numSessionsInWeekend            uint8                     // Number of session in following array
	m_weekendStructure                [12]uint8                 // List of session types to show weekend structure - see appendix for types
	m_sector2LapDistanceStart         float32                   // Distance in m around track where sector 2 starts
	m_sector3LapDistanceStart         float32                   // Distance in m around track where sector 3 starts

	// Aero and DRS zones
	m_activeAeroTrackStatus        uint8             // 0 = Full, 1 = Partial
	m_numActiveAeroZonesFull       uint8             // Number of Active Aero zones to follow
	m_activeAeroZonesFull          [8]ActiveAeroZone // List of Active Aero zones - max 8
	m_numActiveAeroZonesPartial    uint8             // Number of Active Aero zones to follow
	m_activeAeroZonesPartial       [8]ActiveAeroZone // List of Active Aero zones - max 8
	m_numDRSZones                  uint8             // Number of DRS zones to follow
	m_drsZones                     [4]DRSZone        // List of DRS zones - max 4
	m_startReactionTime            float32           // Driver start reaction time in seconds, 0.0f if assisted starts
	m_antiLockBrakesAssist         uint8             // 0 = Off, 1 = On
	m_tractionControlAssist        uint8             // 0 = Off, 1 = Medium, 2 = Full
	m_dynamicRacingLineHiVis       uint8             // 0 = Off, 1 = On
	m_dynamicRacingLineColourBlind uint8             // 0 = Off, 1 = Protanopia, 2 = Deuteranopia, 3 = Tritanopia
	m_recurringRewindPrompt        uint8             // 0 = Off, 1 = On
}
