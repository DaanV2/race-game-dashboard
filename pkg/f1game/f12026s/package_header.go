package f12026s

type PacketHeader struct {
	m_packetFormat            uint16  // 2026
	m_gameYear                uint8   // Game year - last two digits e.g. 25
	m_gameMajorVersion        uint8   // Game major version - "X.00"
	m_gameMinorVersion        uint8   // Game minor version - "1.XX"
	m_packetVersion           uint8   // Version of this packet type, all start from 1
	m_packetId                uint8   // Identifier for the packet type, see below
	m_sessionUID              uint64  // Unique identifier for the session
	m_sessionTime             float32 // Session timestamp
	m_frameIdentifier         uint32  // Identifier for the frame the data was retrieved on
	m_overallFrameIdentifier  uint32  // Overall identifier for the frame the data was retrieved on, doesn't go back after flashbacks
	m_playerCarIndex          uint8   // Index of player's car in the array
	m_secondaryPlayerCarIndex uint8   // Index of secondary player's car in the array (splitscreen) 255 if no second player
}
