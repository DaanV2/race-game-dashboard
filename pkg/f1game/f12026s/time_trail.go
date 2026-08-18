package f12026s

type TimeTrialDataSet struct {
	m_carIdx              uint8  // Index of the car this data relates to
	m_teamId              uint16 // Team id - see appendix
	m_lapTimeInMS         uint32 // Lap time in milliseconds
	m_sector1TimeInMS     uint32 // Sector 1 time in milliseconds
	m_sector2TimeInMS     uint32 // Sector 2 time in milliseconds
	m_sector3TimeInMS     uint32 // Sector 3 time in milliseconds
	m_tractionControl     uint8  // 0 = assist off, 1 = assist on
	m_gearboxAssist       uint8  // 0 = assist off, 1 = assist on
	m_antiLockBrakes      uint8  // 0 = assist off, 1 = assist on
	m_equalCarPerformance uint8  // 0 = Realistic, 1 = Equal
	m_customSetup         uint8  // 0 = No, 1 = Yes
	m_valid               uint8  // 0 = invalid, 1 = valid
}

type PacketTimeTrialData struct {
	m_header                   PacketHeader     // Header
	m_playerSessionBestDataSet TimeTrialDataSet // Player session best data set
	m_personalBestDataSet      TimeTrialDataSet // Personal best data set
	m_rivalDataSet             TimeTrialDataSet // Rival data set
}
