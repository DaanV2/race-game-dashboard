package f12026s

type PacketID uint8

const (
	PACKET_ID_MOTION               = 0  // Contains all motion data for player’s car – only sent while player is in control
	PACKET_ID_SESSION              = 1  // Data about the session – track, time left
	PACKET_ID_LAP_DATA             = 2  // Data about all the lap times of cars in the session
	PACKET_ID_EVENT                = 3  // Various notable events that happen during a session
	PACKET_ID_PARTICIPANTS         = 4  // List of participants in the session, mostly relevant for multiplayer
	PACKET_ID_CAR_SETUPS           = 5  // Packet detailing car setups for cars in the race
	PACKET_ID_CAR_TELEMETRY        = 6  // Telemetry data for all cars
	PACKET_ID_CAR_STATUS           = 7  // Status data for all cars
	PACKET_ID_FINAL_CLASSIFICATION = 8  // Final classification confirmation at the end of a race
	PACKET_ID_LOBBY_INFO           = 9  // Information about players in a multiplayer lobby
	PACKET_ID_CAR_DAMAGE           = 10 // Damage status for all cars
	PACKET_ID_SESSION_HISTORY      = 11 // Lap and tyre data for session
	PACKET_ID_TYRE_SETS            = 12 // Extended tyre set data
	PACKET_ID_MOTION_EX            = 13 // Extended motion data for player car
	PACKET_ID_TIME_TRIAL           = 14 // Time Trial specific data
	PACKET_ID_LAP_POSITIONS        = 15 // Lap positions on each lap so a chart can be constructed
	PACKET_ID_CAR_TELEMETRY_2      = 16 // Additional telemetry data for all cars
)
