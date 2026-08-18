package f12026s

type LobbyInfoData struct {
	m_aiControlled    uint8     // Whether the vehicle is AI (1) or Human (0) controlled
	m_teamId          uint16    // Team id - see appendix (65535 if no team currently selected)
	m_nationality     uint8     // Nationality of the driver
	m_platform        uint8     // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	m_name            [32]uint8 // Name of participant in UTF-8 format – null terminated, Will be truncated with ... (U+2026) if too long
	m_carNumber       uint8     // Car number of the player
	m_yourTelemetry   uint8     // The player's UDP setting, 0 = restricted, 1 = public
	m_showOnlineNames uint8     // The player's show online names setting, 0 = off, 1 = on
	m_techLevel       uint16    // F1 World tech level
	m_readyStatus     uint8     // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketLobbyInfoData struct {
	m_header PacketHeader // Header

	// Packet specific data
	m_numPlayers   uint8 // Number of players in the lobby data
	m_lobbyPlayers [24]LobbyInfoData
}
