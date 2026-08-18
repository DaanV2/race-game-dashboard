package f12026s

type LapHistoryData struct {
	m_lapTimeInMS            uint32 // Lap time in milliseconds
	m_sector1TimeMSPart      uint16 // Sector 1 milliseconds part
	m_sector1TimeMinutesPart uint8  // Sector 1 whole minute part
	m_sector2TimeMSPart      uint16 // Sector 2 time milliseconds part
	m_sector2TimeMinutesPart uint8  // Sector 2 whole minute part
	m_sector3TimeMSPart      uint16 // Sector 3 time milliseconds part
	m_sector3TimeMinutesPart uint8  // Sector 3 whole minute part
	m_lapValidBitFlags       uint8  // 0x01 bit set-lap valid,      0x02 bit set-sector 1 valid
	// 0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
}

type TyreStintHistoryData struct {
	m_endLap             uint8 // Lap the tyre usage ends on (255 of current tyre)
	m_tyreActualCompound uint8 // Actual tyres used by this driver
	m_tyreVisualCompound uint8 // Visual tyres used by this driver
}

type PacketSessionHistoryData struct {
	m_header PacketHeader // Header

	m_carIdx            uint8 // Index of the car this lap data relates to
	m_numLaps           uint8 // Num laps in the data (including current partial lap)
	m_numTyreStints     uint8 // Number of tyre stints in the data
	m_bestLapTimeLapNum uint8 // Lap the best lap time was achieved on
	m_bestSector1LapNum uint8 // Lap the best Sector 1 time was achieved on
	m_bestSector2LapNum uint8 // Lap the best Sector 2 time was achieved on
	m_bestSector3LapNum uint8 // Lap the best Sector 3 time was achieved on

	m_lapHistoryData        [100]LapHistoryData // 100 laps of data max
	m_tyreStintsHistoryData [8]TyreStintHistoryData
}
