package f12026s

type PacketID uint8

const (
	PACKET_ID_MOTION               PacketID = 0  // Contains all motion data for player’s car – only sent while player = is in control
	PACKET_ID_SESSION              PacketID = 1  // Data about the session – track, time left
	PACKET_ID_LAP_DATA             PacketID = 2  // Data about all the lap times of cars in the session
	PACKET_ID_EVENT                PacketID = 3  // Various notable events that happen during a session
	PACKET_ID_PARTICIPANTS         PacketID = 4  // List of participants in the session, mostly relevant for multiplayer
	PACKET_ID_CAR_SETUPS           PacketID = 5  // Packet detailing car setups for cars in the race
	PACKET_ID_CAR_TELEMETRY        PacketID = 6  // Telemetry data for all cars
	PACKET_ID_CAR_STATUS           PacketID = 7  // Status data for all cars
	PACKET_ID_FINAL_CLASSIFICATION PacketID = 8  // Final classification confirmation at the end of a race
	PACKET_ID_LOBBY_INFO           PacketID = 9  // Information about players in a multiplayer lobby
	PACKET_ID_CAR_DAMAGE           PacketID = 10 // Damage status for all cars
	PACKET_ID_SESSION_HISTORY      PacketID = 11 // Lap and tyre data for session
	PACKET_ID_TYRE_SETS            PacketID = 12 // Extended tyre set data
	PACKET_ID_MOTION_EX            PacketID = 13 // Extended motion data for player car
	PACKET_ID_TIME_TRIAL           PacketID = 14 // Time Trial specific data
	PACKET_ID_LAP_POSITIONS        PacketID = 15 // Lap positions on each lap so a chart can be constructed
	PACKET_ID_CAR_TELEMETRY2       PacketID = 16 // Additional telemetry data for all cars
)
