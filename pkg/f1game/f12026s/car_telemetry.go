package f12026s

type CarTelemetryData struct {
	m_speed                   uint16     // Speed of car in kilometres per hour
	m_throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	m_steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	m_brake                   float32    // Amount of brake applied (0.0 to 1.0)
	m_clutch                  uint8      // Amount of clutch applied (0 to 100)
	m_gear                    int8       // Gear selected (1-8, N=0, R=-1)
	m_engineRPM               uint16     // Engine RPM
	m_drs                     uint8      // 0 = off, 1 = on
	m_revLightsPercent        uint8      // Rev lights indicator (percentage)
	m_revLightsBitValue       uint16     // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	m_brakesTemperature       [4]uint16  // Brakes temperature (celsius)
	m_tyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	m_tyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	m_engineTemperature       uint8      // Engine temperature (celsius)
	m_tyresPressure           [4]float32 // Tyres pressure (PSI)
	m_surfaceType             [4]uint8   // Driving surface, see appendices
}

type PacketCarTelemetryData struct {
	m_header PacketHeader // Header

	m_carTelemetryData [24]CarTelemetryData
	// Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 =  Engine, 4 = Temperatures
	// May vary depending on game mode
	m_mfdPanelIndex                uint8
	m_mfdPanelIndexSecondaryPlayer uint8 // See above
	m_suggestedGear                int8  // Suggested gear for the player (1-8), 0 if no gear suggested
}

type CarTelemetry2Data struct {
	m_activeAeroMode               uint8  // 0 = Corner mode, 1 = Straight mode
	m_activeAeroAvailable          uint8  // 0 = not available, 1 = available
	m_activeAeroActivationDistance uint16 // 0 = Active aero not available, non-zero – Active aero will be available in [X] metres
	m_overtakeAvailable            uint8  // 0 = not available, 1 = available
	m_overtakeActive               uint8  // 0 = not active, 1 = active
	m_overtakeActivationDistance   uint16 // 0 = Overtake Mode not available, non-zero –
	m_2026Regulations              uint8  // 0 = vehicle conforms to pre-2026, 1 = 2026 regulations applicable
	m_drivingWrongWay              uint8  // Whether the car is driving the wrong way
}

type PacketCarTelemetry2Data struct {
	m_header            PacketHeader
	m_carTelemetry2Data [24]CarTelemetry2Data
}
