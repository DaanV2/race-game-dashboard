package f12026s

type PacketID uint8

const (
	PACKET_ID_MOTION               PacketID = 0  // Contains all motion data for player’s car – only sent while player is in control
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
	PACKET_ID_CAR_TELEMETRY_2      PacketID = 16 // Additional telemetry data for all cars
)

func (p PacketID) String() string {
	switch p {
	case PACKET_ID_MOTION:
		return "PacketMotion"
	case PACKET_ID_SESSION:
		return "PacketSession"
	case PACKET_ID_LAP_DATA:
		return "PacketLapData"
	case PACKET_ID_EVENT:
		return "PacketEvent"
	case PACKET_ID_PARTICIPANTS:
		return "PacketParticipants"
	case PACKET_ID_CAR_SETUPS:
		return "PacketCarSetups"
	case PACKET_ID_CAR_TELEMETRY:
		return "PacketCarTelemetry"
	case PACKET_ID_CAR_STATUS:
		return "PacketCaStatus"
	case PACKET_ID_FINAL_CLASSIFICATION:
		return "PacketFinalClassification"
	case PACKET_ID_LOBBY_INFO:
		return "PacketLobbyInfo"
	case PACKET_ID_CAR_DAMAGE:
		return "PacketCarDamage"
	case PACKET_ID_SESSION_HISTORY:
		return "PacketSessionHistory"
	case PACKET_ID_TYRE_SETS:
		return "PacketTyreSets"
	case PACKET_ID_MOTION_EX:
		return "PacketMotionEx"
	case PACKET_ID_TIME_TRIAL:
		return "PacketTimeTrial"
	case PACKET_ID_LAP_POSITIONS:
		return "PacketLapPositions"
	case PACKET_ID_CAR_TELEMETRY_2:
		return "PacketCarTelemetry2"
	}

	return "unknown"
}

func PacketIds() []PacketID {
	return []PacketID{
		PACKET_ID_MOTION,
		PACKET_ID_SESSION,
		PACKET_ID_LAP_DATA,
		PACKET_ID_EVENT,
		PACKET_ID_PARTICIPANTS,
		PACKET_ID_CAR_SETUPS,
		PACKET_ID_CAR_TELEMETRY,
		PACKET_ID_CAR_STATUS,
		PACKET_ID_FINAL_CLASSIFICATION,
		PACKET_ID_LOBBY_INFO,
		PACKET_ID_CAR_DAMAGE,
		PACKET_ID_SESSION_HISTORY,
		PACKET_ID_TYRE_SETS,
		PACKET_ID_MOTION_EX,
		PACKET_ID_TIME_TRIAL,
		PACKET_ID_LAP_POSITIONS,
		PACKET_ID_CAR_TELEMETRY_2,
	}
}
