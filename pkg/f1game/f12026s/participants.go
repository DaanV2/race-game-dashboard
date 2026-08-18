package f12026s

// RGB value of a colour
type LiveryColour struct {
	red   uint8
	green uint8
	blue  uint8
}

type ParticipantData struct {
	m_aiControlled    uint8           // Whether the vehicle is AI (1) or Human (0) controlled
	m_driverId        uint16          // Driver id - see appendix, 65535 if network human
	m_networkId       uint16          // Network id – unique identifier for network players
	m_teamId          uint16          // Team id - see appendix
	m_myTeam          uint8           // My team flag – 1 = My Team, 0 = otherwise
	m_raceNumber      uint8           // Race number of the car
	m_nationality     uint8           // Nationality of the driver
	m_name            [32]uint8       // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
	m_yourTelemetry   uint8           // The player's UDP setting, 0 = restricted, 1 = public
	m_showOnlineNames uint8           // The player's show online names setting, 0 = off, 1 = on
	m_techLevel       uint16          // F1 World tech level
	m_platform        uint8           // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	m_numColours      uint8           // Number of colours valid for this car
	m_liveryColours   [4]LiveryColour // Colours for the car
}

type PacketParticipantsData struct {
	m_header        PacketHeader // Header
	m_numActiveCars uint8        // Number of active cars in the data – should match number of cars on HUD
	m_participants  [24]ParticipantData
}
