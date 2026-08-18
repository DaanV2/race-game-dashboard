package f12026s

type FinalClassificationData struct {
	m_position     uint8 // Finishing position
	m_numLaps      uint8 // Number of laps completed
	m_gridPosition uint8 // Grid position of the car
	m_points       uint8 // Number of points scored
	m_numPitStops  uint8 // Number of pit stops made
	// Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = didnotfinish, 5 = disqualified
	// 6 = not classified, 7 = retired
	m_resultStatus uint8
	// Result reason - 0 = invalid, 1 = retired, 2 = finished
	// 3 = terminal damage, 4 = inactive, 5 = not enough laps completed
	// 6 = black flagged, 7 = red flagged, 8 = mechanical failure
	// 9 = session skipped, 10 = session simulated
	m_resultReason      uint8
	m_bestLapTimeInMS   uint32   // Best lap time of the session in milliseconds
	m_totalRaceTime     float64  // Total race time in seconds without penalties
	m_penaltiesTime     uint8    // Total penalties accumulated in seconds
	m_numPenalties      uint8    // Number of penalties applied to this driver
	m_numTyreStints     uint8    // Number of tyres stints up to maximum
	m_tyreStintsActual  [8]uint8 // Actual tyres used by this driver
	m_tyreStintsVisual  [8]uint8 // Visual tyres used by this driver
	m_tyreStintsEndLaps [8]uint8 // The lap number stints end on
}

type PacketFinalClassificationData struct {
	m_header PacketHeader // Header

	m_numCars            uint8 // Number of cars in the final classification
	m_classificationData [24]FinalClassificationData
}
