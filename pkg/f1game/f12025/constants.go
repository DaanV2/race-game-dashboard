package f12025

const (
	CS_MAX_NUM_CARS             uint32 = 22
	CS_MAX_PARTICIPANT_NAME_LEN uint32 = 32
	CS_MAX_TYRE_STINTS          uint32 = 8
	CS_MAX_NUM_TYRE_SETS        uint32 = 13 + 7 // 13 slick and 7 wet weather

	CS_MAX_MARSHALS_ZONE_PER_LAP    uint32 = 21
	CS_MAX_WEATHER_FORECAST_SAMPLES uint32 = 64
	CS_MAX_SESSIONS_IN_WEEKEND      uint32 = 12

	CS_EVENT_STRING_CODE_LEN                        = 4
	CS_MAX_NUM_LAPS_IN_HISTORY                      = 100
	CS_MAX_NUM_LAPS_IN_LAP_POSITIONS_HISTORY_PACKET = 50

	// PACKET_FORMAT can be used to compare against the [PacketHeader.PacketFormat]
	PACKET_FORMAT uint16 = 2025
)
