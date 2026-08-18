package f12026s

type PacketLapPositionsData struct {
	m_header PacketHeader // Header

	// Packet specific data
	m_numLaps  uint8 // Number of laps in the data
	m_lapStart uint8 // Index of the lap where the data starts, 0 indexed

	// Array holding the position of the car in a given lap, 0 if no record
	m_positionForVehicleIdx [50][24]uint8
}
