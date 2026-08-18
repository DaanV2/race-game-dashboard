package f12026s

type CarStatusData struct {
	m_tractionControl       uint8   // Traction control - 0 = off, 1 = medium, 2 = full
	m_antiLockBrakes        uint8   // 0 (off) - 1 (on)
	m_fuelMix               uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	m_frontBrakeBias        uint8   // Front brake bias (percentage)
	m_pitLimiterStatus      uint8   // Pit limiter status - 0 = off, 1 = on
	m_fuelInTank            float32 // Current fuel mass
	m_fuelCapacity          float32 // Fuel capacity
	m_fuelRemainingLaps     float32 // Fuel remaining in terms of laps (value on MFD)
	m_maxRPM                uint16  // Cars max RPM, point of rev limiter
	m_idleRPM               uint16  // Cars idle RPM
	m_maxGears              uint8   // Maximum number of gears
	m_drsAllowed            uint8   // 0 = not allowed, 1 = allowed
	m_drsActivationDistance uint16  // 0 = DRS not available, non-zero - DRS will be available in [X] metres

	// F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1
	// 21 = C0, 22 = C6, 7 = inter, 8 = wet
	// F1 Classic - 9 = dry, 10 = wet
	// F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard
	// 15 = wet
	m_actualTyreCompound uint8
	// F1 visual (can be different from actual compound)
	// 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet
	// F1 Classic – same as above
	// F2 ‘20, 15 = wet, 19 – super soft, 20 = soft
	// 21 = medium, 22 = hard
	m_visualTyreCompound uint8

	m_tyresAgeLaps            uint8   // Age in laps of the current set of tyres
	m_vehicleFiaFlags         int8    // -1 = invalid/unknown, 0 = none, 1 = green 2 = blue, 3 = yellow
	m_enginePowerICE          float32 // Engine power output of ICE (W)
	m_enginePowerMGUK         float32 // Engine power output of MGU-K (W)
	m_ersStoreEnergy          float32 // ERS energy store in Joules
	m_ersDeployMode           uint8   // ERS deployment mode, 0 = none, 1 = medium 2 = hotlap, 3 = boost
	m_ersHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	m_ersHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	m_ersHarvestedLimitPerLap float32 // ERS energy harvest limit for this lap
	m_ersDeployedThisLap      float32 // ERS energy deployed this lap
	m_networkPaused           uint8   // Whether the car is paused in a network game
}

type PacketCarStatusData struct {
	m_header PacketHeader // Header

	m_carStatusData [24]CarStatusData
}
