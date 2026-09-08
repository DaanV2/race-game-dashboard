package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
	"github.com/daanv2/race-game-dashboard/pkg/generics"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketSessionData -ignore-field MarshalZones,WeatherForecastSamples

type PacketSessionData struct {
	Header PacketHeader // Header

	Weather                         uint8                                                  // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature                int8                                                   // Track temp. in degrees celsius
	AirTemperature                  int8                                                   // Air temp. in degrees celsius
	TotalLaps                       uint8                                                  // Total number of laps in this race
	TrackLength                     uint16                                                 // Track length in metres
	SessionType                     f1common.SessionTypeID                                 // (uint8) 0 = unknown, see appendix
	TrackId                         f1common.TrackID                                       // (int8) -1 for unknown, see appendix
	Formula                         uint8                                                  // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, 3 = F1 Generic, 4 = Beta, 6 = Esports, 8 = F1 World, 9 = F1 Elimination
	SessionTimeLeft                 uint16                                                 // Time left in session in seconds
	SessionDuration                 uint16                                                 // Session duration in seconds
	PitSpeedLimit                   uint8                                                  // Pit speed limit in kilometres per hour
	GamePaused                      uint8                                                  // Whether the game is paused - network game only
	IsSpectating                    uint8                                                  // Whether the player is spectating
	SpectatorCarIndex               uint8                                                  // Index of the car being spectated
	SliProNativeSupport             uint8                                                  // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones                 uint8                                                  // Number of marshal zones to follow
	MarshalZones                    [CS_MAX_MARSHALS_ZONE_PER_LAP]MarshalZone              // List of marshal zones - max 21
	SafetyCarStatus                 uint8                                                  // 0 = no safety car, 1 = full, 2 = virtual, 3 = formation lap
	NetworkGame                     uint8                                                  // 0 = offline, 1 = online
	NumWeatherForecastSamples       uint8                                                  // Number of weather samples to follow
	WeatherForecastSamples          [CS_MAX_WEATHER_FORECAST_SAMPLES]WeatherForecastSample // Array of weather forecast samples
	ForecastAccuracy                uint8                                                  // 0 = Perfect, 1 = Approximate
	AiDifficulty                    uint8                                                  // AI difficulty rating - 0-110
	SeasonLinkIdentifier            uint32                                                 // Identifier for season - persists across saves
	WeekendLinkIdentifier           uint32                                                 // Identifier for weekend - persists across saves
	SessionLinkIdentifier           uint32                                                 // Identifier for session - persists across saves
	PitStopWindowIdealLap           uint8                                                  // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap          uint8                                                  // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition           uint8                                                  // Predicted position to rejoin at (player)
	SteeringAssist                  uint8                                                  // 0 = off, 1 = on
	BrakingAssist                   uint8                                                  // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist                   uint8                                                  // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist                       uint8                                                  // 0 = off, 1 = on
	PitReleaseAssist                uint8                                                  // 0 = off, 1 = on
	ERSAssist                       uint8                                                  // 0 = off, 1 = on
	DRSAssist                       uint8                                                  // 0 = off, 1 = on
	DynamicRacingLine               uint8                                                  // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType           uint8                                                  // 0 = 2D, 1 = 3D
	GameMode                        f1common.GamemodeID                                    // Game mode id (uint8) - see appendix
	RuleSet                         f1common.RulesetID                                     // Ruleset (uint8) - see appendix
	TimeOfDay                       uint32                                                 // Local time of day - minutes since midnight
	SessionLength                   uint8                                                  // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium, 5 = Medium Long, 6 = Long, 7 = Full
	SpeedUnitsLeadPlayer            uint8                                                  // 0 = MPH, 1 = KPH
	TemperatureUnitsLeadPlayer      uint8                                                  // 0 = Celsius, 1 = Fahrenheit
	SpeedUnitsSecondaryPlayer       uint8                                                  // 0 = MPH, 1 = KPH
	TemperatureUnitsSecondaryPlayer uint8                                                  // 0 = Celsius, 1 = Fahrenheit
	NumSafetyCarPeriods             uint8                                                  // Number of safety cars called during session
	NumVirtualSafetyCarPeriods      uint8                                                  // Number of virtual safety cars called
	NumRedFlagPeriods               uint8                                                  // Number of red flags called during session
	EqualCarPerformance             uint8                                                  // 0 = Off, 1 = On
	RecoveryMode                    uint8                                                  // 0 = None, 1 = Flashbacks, 2 = Auto-recovery
	FlashbackLimit                  uint8                                                  // 0 = Low, 1 = Medium, 2 = High, 3 = Unlimited
	SurfaceType                     uint8                                                  // 0 = Simplified, 1 = Realistic
	LowFuelMode                     uint8                                                  // 0 = Easy, 1 = Hard
	RaceStarts                      uint8                                                  // 0 = Manual, 1 = Assisted
	TyreTemperature                 uint8                                                  // 0 = Surface only, 1 = Surface & Carcass
	PitLaneTyreSim                  uint8                                                  // 0 = On, 1 = Off
	CarDamage                       uint8                                                  // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Simulation
	CarDamageRate                   uint8                                                  // 0 = Reduced, 1 = Standard, 2 = Simulation
	Collisions                      uint8                                                  // 0 = Off, 1 = Player-to-Player Off, 2 = On
	CollisionsOffForFirstLapOnly    uint8                                                  // 0 = Disabled, 1 = Enabled
	MpUnsafePitRelease              uint8                                                  // 0 = On, 1 = Off (Multiplayer)
	MpOffForGriefing                uint8                                                  // 0 = Disabled, 1 = Enabled (Multiplayer)
	CornerCuttingStringency         uint8                                                  // 0 = Regular, 1 = Strict
	ParcFermeRules                  uint8                                                  // 0 = Off, 1 = On
	PitStopExperience               uint8                                                  // 0 = Automatic, 1 = Broadcast, 2 = Immersive
	SafetyCar                       uint8                                                  // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	SafetyCarExperience             uint8                                                  // 0 = Broadcast, 1 = Immersive
	FormationLap                    uint8                                                  // 0 = Off, 1 = On
	FormationLapExperience          uint8                                                  // 0 = Broadcast, 1 = Immersive
	RedFlags                        uint8                                                  // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	AffectsLicenceLevelSolo         uint8                                                  // 0 = Off, 1 = On
	AffectsLicenceLevelMP           uint8                                                  // 0 = Off, 1 = On
	NumSessionsInWeekend            uint8                                                  // Number of session in following array
	WeekendStructure                [CS_MAX_SESSIONS_IN_WEEKEND]uint8                      // TODO List of session types to show weekend structure - see appendix for types
	Sector2LapDistanceStart         float32                                                // Distance in m around track where sector 2 starts
	Sector3LapDistanceStart         float32                                                // Distance in m around track where sector 3 starts
}

// GetPacketID returns the identification of this packet
func (data *PacketSessionData) GetPacketID() PacketID { return PACKET_ID_SESSION }

// GetMarshalZones returns the MarshalZones of *PacketSessionData
func (data *PacketSessionData) GetMarshalZones(marshalZone int) MarshalZone {
	return data.MarshalZones[marshalZone]
}

// SetMarshalZones stores the MarshalZones of *PacketSessionData
func (data *PacketSessionData) SetMarshalZones(marshalZone int, v MarshalZone) {
	data.MarshalZones[marshalZone] = v
}

// GetWeatherForecastSamples returns the WeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) GetWeatherForecastSamples(sample int) WeatherForecastSample {
	return data.WeatherForecastSamples[sample]
}

// SetWeatherForecastSamples stores the WeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) SetWeatherForecastSamples(sample int, v WeatherForecastSample) {
	data.WeatherForecastSamples[sample] = v
}

func (data *PacketSessionData) WeekendStructureIDs() (sessions [CS_MAX_SESSIONS_IN_WEEKEND]f1common.SessionTypeID) {
	generics.CopySlice(sessions[:], data.WeekendStructure[:])

	return
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketSessionData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	data.Weather = reader.ReadUint8()
	data.TrackTemperature = reader.ReadInt8()
	data.AirTemperature = reader.ReadInt8()
	data.TotalLaps = reader.ReadUint8()
	data.TrackLength = reader.ReadUint16()
	data.SessionType = f1common.SessionTypeID(reader.ReadUint8())
	data.TrackId = f1common.TrackID(reader.ReadInt8())
	data.Formula = reader.ReadUint8()
	data.SessionTimeLeft = reader.ReadUint16()
	data.SessionDuration = reader.ReadUint16()
	data.PitSpeedLimit = reader.ReadUint8()
	data.GamePaused = reader.ReadUint8()
	data.IsSpectating = reader.ReadUint8()
	data.SpectatorCarIndex = reader.ReadUint8()
	data.SliProNativeSupport = reader.ReadUint8()
	data.NumMarshalZones = reader.ReadUint8()

	for i := range data.MarshalZones {
		data.MarshalZones[i].Parse(reader)
	}

	data.SafetyCarStatus = reader.ReadUint8()
	data.NetworkGame = reader.ReadUint8()
	data.NumWeatherForecastSamples = reader.ReadUint8()

	for i := range data.WeatherForecastSamples {
		data.WeatherForecastSamples[i].Parse(reader)
	}

	data.ForecastAccuracy = reader.ReadUint8()
	data.AiDifficulty = reader.ReadUint8()
	data.SeasonLinkIdentifier = reader.ReadUint32()
	data.WeekendLinkIdentifier = reader.ReadUint32()
	data.SessionLinkIdentifier = reader.ReadUint32()
	data.PitStopWindowIdealLap = reader.ReadUint8()
	data.PitStopWindowLatestLap = reader.ReadUint8()
	data.PitStopRejoinPosition = reader.ReadUint8()
	data.SteeringAssist = reader.ReadUint8()
	data.BrakingAssist = reader.ReadUint8()
	data.GearboxAssist = reader.ReadUint8()
	data.PitAssist = reader.ReadUint8()
	data.PitReleaseAssist = reader.ReadUint8()
	data.ERSAssist = reader.ReadUint8()
	data.DRSAssist = reader.ReadUint8()
	data.DynamicRacingLine = reader.ReadUint8()
	data.DynamicRacingLineType = reader.ReadUint8()
	data.GameMode = f1common.GamemodeID(reader.ReadUint8())
	data.RuleSet = f1common.RulesetID(reader.ReadUint8())
	data.TimeOfDay = reader.ReadUint32()
	data.SessionLength = reader.ReadUint8()
	data.SpeedUnitsLeadPlayer = reader.ReadUint8()
	data.TemperatureUnitsLeadPlayer = reader.ReadUint8()
	data.SpeedUnitsSecondaryPlayer = reader.ReadUint8()
	data.TemperatureUnitsSecondaryPlayer = reader.ReadUint8()
	data.NumSafetyCarPeriods = reader.ReadUint8()
	data.NumVirtualSafetyCarPeriods = reader.ReadUint8()
	data.NumRedFlagPeriods = reader.ReadUint8()
	data.EqualCarPerformance = reader.ReadUint8()
	data.RecoveryMode = reader.ReadUint8()
	data.FlashbackLimit = reader.ReadUint8()
	data.SurfaceType = reader.ReadUint8()
	data.LowFuelMode = reader.ReadUint8()
	data.RaceStarts = reader.ReadUint8()
	data.TyreTemperature = reader.ReadUint8()
	data.PitLaneTyreSim = reader.ReadUint8()
	data.CarDamage = reader.ReadUint8()
	data.CarDamageRate = reader.ReadUint8()
	data.Collisions = reader.ReadUint8()
	data.CollisionsOffForFirstLapOnly = reader.ReadUint8()
	data.MpUnsafePitRelease = reader.ReadUint8()
	data.MpOffForGriefing = reader.ReadUint8()
	data.CornerCuttingStringency = reader.ReadUint8()
	data.ParcFermeRules = reader.ReadUint8()
	data.PitStopExperience = reader.ReadUint8()
	data.SafetyCar = reader.ReadUint8()
	data.SafetyCarExperience = reader.ReadUint8()
	data.FormationLap = reader.ReadUint8()
	data.FormationLapExperience = reader.ReadUint8()
	data.RedFlags = reader.ReadUint8()
	data.AffectsLicenceLevelSolo = reader.ReadUint8()
	data.AffectsLicenceLevelMP = reader.ReadUint8()
	data.NumSessionsInWeekend = reader.ReadUint8()
	reader.Read(data.WeekendStructure[:])
	data.Sector2LapDistanceStart = reader.ReadFloat32()
	data.Sector3LapDistanceStart = reader.ReadFloat32()
}
