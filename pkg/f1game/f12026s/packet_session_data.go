package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketSessionData struct {
	Header                          PacketHeader                                           // Header
	Weather                         uint8                                                  // Weather - 0 = clear, 1 = light cloud, 2 = overcast  3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature                int8                                                   // Track temp. in degrees celsius
	AirTemperature                  int8                                                   // Air temp. in degrees celsius
	TotalLaps                       uint8                                                  // Total number of laps in this race
	TrackLength                     uint16                                                 // Track length in metres
	SessionType                     uint8                                                  // 0 = unknown, see appendix
	TrackId                         int8                                                   // -1 for unknown, see appendix
	Formula                         uint8                                                  // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,  3 = F1 Generic, 4 = Beta, 6 = Esports  8 = F1 World, 9 = F1 Elimination, 13 = F1 26
	SessionTimeLeft                 uint16                                                 // Time left in session in seconds
	SessionDuration                 uint16                                                 // Session duration in seconds
	PitSpeedLimit                   uint8                                                  // Pit speed limit in kilometres per hour
	GamePaused                      uint8                                                  // Whether the game is paused – network game only
	IsSpectating                    uint8                                                  // Whether the player is spectating
	SpectatorCarIndex               uint8                                                  // Index of the car being spectated
	SliProNativeSupport             uint8                                                  // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones                 uint8                                                  // Number of marshal zones to follow
	MarshalZones                    [CS_MAX_MARSHALS_ZONE_PER_LAP]MarshalZone              // List of marshal zones – max 21
	SafetyCarStatus                 uint8                                                  // 0 = no safety car, 1 = full  2 = virtual, 3 = formation lap
	NetworkGame                     uint8                                                  // 0 = offline, 1 = online
	NumWeatherForecastSamples       uint8                                                  // Number of weather samples to follow
	WeatherForecastSamples          [CS_MAX_WEATHER_FORECAST_SAMPLES]WeatherForecastSample // Array of weather forecast samples
	ForecastAccuracy                uint8                                                  // 0 = Perfect, 1 = Approximate
	AiDifficulty                    uint8                                                  // AI Difficulty rating – 0-110
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
	GameMode                        uint8                                                  // Game mode id - see appendix
	RuleSet                         uint8                                                  // Ruleset - see appendix
	TimeOfDay                       uint32                                                 // Local time of day - minutes since midnight
	SessionLength                   uint8                                                  // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium  5 = Medium Long, 6 = Long, 7 = Full
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
	WeekendStructure                [12]uint8                                              // List of session types to show weekend  structure - see appendix for types
	Sector2LapDistanceStart         float32                                                // Distance in m around track where sector 2 starts
	Sector3LapDistanceStart         float32                                                // Distance in m around track where sector 3 starts  Aero and DRS zones
	ActiveAeroTrackStatus           uint8                                                  // 0 = Full, 1 = Partial
	NumActiveAeroZonesFull          uint8                                                  // Number of Active Aero zones to follow
	ActiveAeroZonesFull             [8]ActiveAeroZone                                      // List of Active Aero zones - max 8
	NumActiveAeroZonesPartial       uint8                                                  // Number of Active Aero zones to follow
	ActiveAeroZonesPartial          [8]ActiveAeroZone                                      // List of Active Aero zones - max 8
	NumDRSZones                     uint8                                                  // Number of DRS zones to follow
	DrsZones                        [4]DRSZone                                             // List of DRS zones - max 4
	StartReactionTime               float32                                                // Driver start reaction time in seconds  0.0f if assisted starts
	AntiLockBrakesAssist            uint8                                                  // 0 = Off, 1 = On
	TractionControlAssist           uint8                                                  // 0 = Off, 1 = Medium, 2 = Full
	DynamicRacingLineHiVis          uint8                                                  // 0 = Off, 1 = On
	DynamicRacingLineColourBlind    uint8                                                  // 0 = Off, 1 = Protanopia, 2 = Deuteranopia  3 = Tritanopia
	RecurringRewindPrompt           uint8                                                  // 0 = Off, 1 = On
}

// GetHeader returns the Header of *PacketSessionData
func (data *PacketSessionData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketSessionData
func (data *PacketSessionData) SetHeader(v PacketHeader) { data.Header = v }

// GetWeather returns the Weather of *PacketSessionData
func (data *PacketSessionData) GetWeather() uint8 { return data.Weather }

// SetWeather stores the Weather of *PacketSessionData
func (data *PacketSessionData) SetWeather(v uint8) { data.Weather = v }

// GetTrackTemperature returns the TrackTemperature of *PacketSessionData
func (data *PacketSessionData) GetTrackTemperature() int8 { return data.TrackTemperature }

// SetTrackTemperature stores the TrackTemperature of *PacketSessionData
func (data *PacketSessionData) SetTrackTemperature(v int8) { data.TrackTemperature = v }

// GetAirTemperature returns the AirTemperature of *PacketSessionData
func (data *PacketSessionData) GetAirTemperature() int8 { return data.AirTemperature }

// SetAirTemperature stores the AirTemperature of *PacketSessionData
func (data *PacketSessionData) SetAirTemperature(v int8) { data.AirTemperature = v }

// GetTotalLaps returns the TotalLaps of *PacketSessionData
func (data *PacketSessionData) GetTotalLaps() uint8 { return data.TotalLaps }

// SetTotalLaps stores the TotalLaps of *PacketSessionData
func (data *PacketSessionData) SetTotalLaps(v uint8) { data.TotalLaps = v }

// GetTrackLength returns the TrackLength of *PacketSessionData
func (data *PacketSessionData) GetTrackLength() uint16 { return data.TrackLength }

// SetTrackLength stores the TrackLength of *PacketSessionData
func (data *PacketSessionData) SetTrackLength(v uint16) { data.TrackLength = v }

// GetSessionType returns the SessionType of *PacketSessionData
func (data *PacketSessionData) GetSessionType() uint8 { return data.SessionType }

// SetSessionType stores the SessionType of *PacketSessionData
func (data *PacketSessionData) SetSessionType(v uint8) { data.SessionType = v }

// GetTrackId returns the TrackId of *PacketSessionData
func (data *PacketSessionData) GetTrackId() int8 { return data.TrackId }

// SetTrackId stores the TrackId of *PacketSessionData
func (data *PacketSessionData) SetTrackId(v int8) { data.TrackId = v }

// GetFormula returns the Formula of *PacketSessionData
func (data *PacketSessionData) GetFormula() uint8 { return data.Formula }

// SetFormula stores the Formula of *PacketSessionData
func (data *PacketSessionData) SetFormula(v uint8) { data.Formula = v }

// GetSessionTimeLeft returns the SessionTimeLeft of *PacketSessionData
func (data *PacketSessionData) GetSessionTimeLeft() uint16 { return data.SessionTimeLeft }

// SetSessionTimeLeft stores the SessionTimeLeft of *PacketSessionData
func (data *PacketSessionData) SetSessionTimeLeft(v uint16) { data.SessionTimeLeft = v }

// GetSessionDuration returns the SessionDuration of *PacketSessionData
func (data *PacketSessionData) GetSessionDuration() uint16 { return data.SessionDuration }

// SetSessionDuration stores the SessionDuration of *PacketSessionData
func (data *PacketSessionData) SetSessionDuration(v uint16) { data.SessionDuration = v }

// GetPitSpeedLimit returns the PitSpeedLimit of *PacketSessionData
func (data *PacketSessionData) GetPitSpeedLimit() uint8 { return data.PitSpeedLimit }

// SetPitSpeedLimit stores the PitSpeedLimit of *PacketSessionData
func (data *PacketSessionData) SetPitSpeedLimit(v uint8) { data.PitSpeedLimit = v }

// GetGamePaused returns the GamePaused of *PacketSessionData
func (data *PacketSessionData) GetGamePaused() uint8 { return data.GamePaused }

// SetGamePaused stores the GamePaused of *PacketSessionData
func (data *PacketSessionData) SetGamePaused(v uint8) { data.GamePaused = v }

// GetIsSpectating returns the IsSpectating of *PacketSessionData
func (data *PacketSessionData) GetIsSpectating() uint8 { return data.IsSpectating }

// SetIsSpectating stores the IsSpectating of *PacketSessionData
func (data *PacketSessionData) SetIsSpectating(v uint8) { data.IsSpectating = v }

// GetSpectatorCarIndex returns the SpectatorCarIndex of *PacketSessionData
func (data *PacketSessionData) GetSpectatorCarIndex() uint8 { return data.SpectatorCarIndex }

// SetSpectatorCarIndex stores the SpectatorCarIndex of *PacketSessionData
func (data *PacketSessionData) SetSpectatorCarIndex(v uint8) { data.SpectatorCarIndex = v }

// GetSliProNativeSupport returns the SliProNativeSupport of *PacketSessionData
func (data *PacketSessionData) GetSliProNativeSupport() uint8 { return data.SliProNativeSupport }

// SetSliProNativeSupport stores the SliProNativeSupport of *PacketSessionData
func (data *PacketSessionData) SetSliProNativeSupport(v uint8) { data.SliProNativeSupport = v }

// GetNumMarshalZones returns the NumMarshalZones of *PacketSessionData
func (data *PacketSessionData) GetNumMarshalZones() uint8 { return data.NumMarshalZones }

// SetNumMarshalZones stores the NumMarshalZones of *PacketSessionData
func (data *PacketSessionData) SetNumMarshalZones(v uint8) { data.NumMarshalZones = v }

// GetMarshalZones returns the MarshalZones of *PacketSessionData
func (data *PacketSessionData) GetMarshalZones(marshalZone int) MarshalZone {
	return data.MarshalZones[marshalZone]
}

// SetMarshalZones stores the MarshalZones of *PacketSessionData
func (data *PacketSessionData) SetMarshalZones(marshalZone int, v MarshalZone) {
	data.MarshalZones[marshalZone] = v
}

// GetSafetyCarStatus returns the SafetyCarStatus of *PacketSessionData
func (data *PacketSessionData) GetSafetyCarStatus() uint8 { return data.SafetyCarStatus }

// SetSafetyCarStatus stores the SafetyCarStatus of *PacketSessionData
func (data *PacketSessionData) SetSafetyCarStatus(v uint8) { data.SafetyCarStatus = v }

// GetNetworkGame returns the NetworkGame of *PacketSessionData
func (data *PacketSessionData) GetNetworkGame() uint8 { return data.NetworkGame }

// SetNetworkGame stores the NetworkGame of *PacketSessionData
func (data *PacketSessionData) SetNetworkGame(v uint8) { data.NetworkGame = v }

// GetNumWeatherForecastSamples returns the NumWeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) GetNumWeatherForecastSamples() uint8 {
	return data.NumWeatherForecastSamples
}

// SetNumWeatherForecastSamples stores the NumWeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) SetNumWeatherForecastSamples(v uint8) {
	data.NumWeatherForecastSamples = v
}

// GetWeatherForecastSamples returns the WeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) GetWeatherForecastSamples(sample int) WeatherForecastSample {
	return data.WeatherForecastSamples[sample]
}

// SetWeatherForecastSamples stores the WeatherForecastSamples of *PacketSessionData
func (data *PacketSessionData) SetWeatherForecastSamples(sample int, v WeatherForecastSample) {
	data.WeatherForecastSamples[sample] = v
}

// GetForecastAccuracy returns the ForecastAccuracy of *PacketSessionData
func (data *PacketSessionData) GetForecastAccuracy() uint8 { return data.ForecastAccuracy }

// SetForecastAccuracy stores the ForecastAccuracy of *PacketSessionData
func (data *PacketSessionData) SetForecastAccuracy(v uint8) { data.ForecastAccuracy = v }

// GetAiDifficulty returns the AiDifficulty of *PacketSessionData
func (data *PacketSessionData) GetAiDifficulty() uint8 { return data.AiDifficulty }

// SetAiDifficulty stores the AiDifficulty of *PacketSessionData
func (data *PacketSessionData) SetAiDifficulty(v uint8) { data.AiDifficulty = v }

// GetSeasonLinkIdentifier returns the SeasonLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) GetSeasonLinkIdentifier() uint32 { return data.SeasonLinkIdentifier }

// SetSeasonLinkIdentifier stores the SeasonLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) SetSeasonLinkIdentifier(v uint32) { data.SeasonLinkIdentifier = v }

// GetWeekendLinkIdentifier returns the WeekendLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) GetWeekendLinkIdentifier() uint32 { return data.WeekendLinkIdentifier }

// SetWeekendLinkIdentifier stores the WeekendLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) SetWeekendLinkIdentifier(v uint32) { data.WeekendLinkIdentifier = v }

// GetSessionLinkIdentifier returns the SessionLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) GetSessionLinkIdentifier() uint32 { return data.SessionLinkIdentifier }

// SetSessionLinkIdentifier stores the SessionLinkIdentifier of *PacketSessionData
func (data *PacketSessionData) SetSessionLinkIdentifier(v uint32) { data.SessionLinkIdentifier = v }

// GetPitStopWindowIdealLap returns the PitStopWindowIdealLap of *PacketSessionData
func (data *PacketSessionData) GetPitStopWindowIdealLap() uint8 { return data.PitStopWindowIdealLap }

// SetPitStopWindowIdealLap stores the PitStopWindowIdealLap of *PacketSessionData
func (data *PacketSessionData) SetPitStopWindowIdealLap(v uint8) { data.PitStopWindowIdealLap = v }

// GetPitStopWindowLatestLap returns the PitStopWindowLatestLap of *PacketSessionData
func (data *PacketSessionData) GetPitStopWindowLatestLap() uint8 { return data.PitStopWindowLatestLap }

// SetPitStopWindowLatestLap stores the PitStopWindowLatestLap of *PacketSessionData
func (data *PacketSessionData) SetPitStopWindowLatestLap(v uint8) { data.PitStopWindowLatestLap = v }

// GetPitStopRejoinPosition returns the PitStopRejoinPosition of *PacketSessionData
func (data *PacketSessionData) GetPitStopRejoinPosition() uint8 { return data.PitStopRejoinPosition }

// SetPitStopRejoinPosition stores the PitStopRejoinPosition of *PacketSessionData
func (data *PacketSessionData) SetPitStopRejoinPosition(v uint8) { data.PitStopRejoinPosition = v }

// GetSteeringAssist returns the SteeringAssist of *PacketSessionData
func (data *PacketSessionData) GetSteeringAssist() uint8 { return data.SteeringAssist }

// SetSteeringAssist stores the SteeringAssist of *PacketSessionData
func (data *PacketSessionData) SetSteeringAssist(v uint8) { data.SteeringAssist = v }

// GetBrakingAssist returns the BrakingAssist of *PacketSessionData
func (data *PacketSessionData) GetBrakingAssist() uint8 { return data.BrakingAssist }

// SetBrakingAssist stores the BrakingAssist of *PacketSessionData
func (data *PacketSessionData) SetBrakingAssist(v uint8) { data.BrakingAssist = v }

// GetGearboxAssist returns the GearboxAssist of *PacketSessionData
func (data *PacketSessionData) GetGearboxAssist() uint8 { return data.GearboxAssist }

// SetGearboxAssist stores the GearboxAssist of *PacketSessionData
func (data *PacketSessionData) SetGearboxAssist(v uint8) { data.GearboxAssist = v }

// GetPitAssist returns the PitAssist of *PacketSessionData
func (data *PacketSessionData) GetPitAssist() uint8 { return data.PitAssist }

// SetPitAssist stores the PitAssist of *PacketSessionData
func (data *PacketSessionData) SetPitAssist(v uint8) { data.PitAssist = v }

// GetPitReleaseAssist returns the PitReleaseAssist of *PacketSessionData
func (data *PacketSessionData) GetPitReleaseAssist() uint8 { return data.PitReleaseAssist }

// SetPitReleaseAssist stores the PitReleaseAssist of *PacketSessionData
func (data *PacketSessionData) SetPitReleaseAssist(v uint8) { data.PitReleaseAssist = v }

// GetERSAssist returns the ERSAssist of *PacketSessionData
func (data *PacketSessionData) GetERSAssist() uint8 { return data.ERSAssist }

// SetERSAssist stores the ERSAssist of *PacketSessionData
func (data *PacketSessionData) SetERSAssist(v uint8) { data.ERSAssist = v }

// GetDRSAssist returns the DRSAssist of *PacketSessionData
func (data *PacketSessionData) GetDRSAssist() uint8 { return data.DRSAssist }

// SetDRSAssist stores the DRSAssist of *PacketSessionData
func (data *PacketSessionData) SetDRSAssist(v uint8) { data.DRSAssist = v }

// GetDynamicRacingLine returns the DynamicRacingLine of *PacketSessionData
func (data *PacketSessionData) GetDynamicRacingLine() uint8 { return data.DynamicRacingLine }

// SetDynamicRacingLine stores the DynamicRacingLine of *PacketSessionData
func (data *PacketSessionData) SetDynamicRacingLine(v uint8) { data.DynamicRacingLine = v }

// GetDynamicRacingLineType returns the DynamicRacingLineType of *PacketSessionData
func (data *PacketSessionData) GetDynamicRacingLineType() uint8 { return data.DynamicRacingLineType }

// SetDynamicRacingLineType stores the DynamicRacingLineType of *PacketSessionData
func (data *PacketSessionData) SetDynamicRacingLineType(v uint8) { data.DynamicRacingLineType = v }

// GetGameMode returns the GameMode of *PacketSessionData
func (data *PacketSessionData) GetGameMode() uint8 { return data.GameMode }

// SetGameMode stores the GameMode of *PacketSessionData
func (data *PacketSessionData) SetGameMode(v uint8) { data.GameMode = v }

// GetRuleSet returns the RuleSet of *PacketSessionData
func (data *PacketSessionData) GetRuleSet() uint8 { return data.RuleSet }

// SetRuleSet stores the RuleSet of *PacketSessionData
func (data *PacketSessionData) SetRuleSet(v uint8) { data.RuleSet = v }

// GetTimeOfDay returns the TimeOfDay of *PacketSessionData
func (data *PacketSessionData) GetTimeOfDay() uint32 { return data.TimeOfDay }

// SetTimeOfDay stores the TimeOfDay of *PacketSessionData
func (data *PacketSessionData) SetTimeOfDay(v uint32) { data.TimeOfDay = v }

// GetSessionLength returns the SessionLength of *PacketSessionData
func (data *PacketSessionData) GetSessionLength() uint8 { return data.SessionLength }

// SetSessionLength stores the SessionLength of *PacketSessionData
func (data *PacketSessionData) SetSessionLength(v uint8) { data.SessionLength = v }

// GetSpeedUnitsLeadPlayer returns the SpeedUnitsLeadPlayer of *PacketSessionData
func (data *PacketSessionData) GetSpeedUnitsLeadPlayer() uint8 { return data.SpeedUnitsLeadPlayer }

// SetSpeedUnitsLeadPlayer stores the SpeedUnitsLeadPlayer of *PacketSessionData
func (data *PacketSessionData) SetSpeedUnitsLeadPlayer(v uint8) { data.SpeedUnitsLeadPlayer = v }

// GetTemperatureUnitsLeadPlayer returns the TemperatureUnitsLeadPlayer of *PacketSessionData
func (data *PacketSessionData) GetTemperatureUnitsLeadPlayer() uint8 {
	return data.TemperatureUnitsLeadPlayer
}

// SetTemperatureUnitsLeadPlayer stores the TemperatureUnitsLeadPlayer of *PacketSessionData
func (data *PacketSessionData) SetTemperatureUnitsLeadPlayer(v uint8) {
	data.TemperatureUnitsLeadPlayer = v
}

// GetSpeedUnitsSecondaryPlayer returns the SpeedUnitsSecondaryPlayer of *PacketSessionData
func (data *PacketSessionData) GetSpeedUnitsSecondaryPlayer() uint8 {
	return data.SpeedUnitsSecondaryPlayer
}

// SetSpeedUnitsSecondaryPlayer stores the SpeedUnitsSecondaryPlayer of *PacketSessionData
func (data *PacketSessionData) SetSpeedUnitsSecondaryPlayer(v uint8) {
	data.SpeedUnitsSecondaryPlayer = v
}

// GetTemperatureUnitsSecondaryPlayer returns the TemperatureUnitsSecondaryPlayer of *PacketSessionData
func (data *PacketSessionData) GetTemperatureUnitsSecondaryPlayer() uint8 {
	return data.TemperatureUnitsSecondaryPlayer
}

// SetTemperatureUnitsSecondaryPlayer stores the TemperatureUnitsSecondaryPlayer of *PacketSessionData
func (data *PacketSessionData) SetTemperatureUnitsSecondaryPlayer(v uint8) {
	data.TemperatureUnitsSecondaryPlayer = v
}

// GetNumSafetyCarPeriods returns the NumSafetyCarPeriods of *PacketSessionData
func (data *PacketSessionData) GetNumSafetyCarPeriods() uint8 { return data.NumSafetyCarPeriods }

// SetNumSafetyCarPeriods stores the NumSafetyCarPeriods of *PacketSessionData
func (data *PacketSessionData) SetNumSafetyCarPeriods(v uint8) { data.NumSafetyCarPeriods = v }

// GetNumVirtualSafetyCarPeriods returns the NumVirtualSafetyCarPeriods of *PacketSessionData
func (data *PacketSessionData) GetNumVirtualSafetyCarPeriods() uint8 {
	return data.NumVirtualSafetyCarPeriods
}

// SetNumVirtualSafetyCarPeriods stores the NumVirtualSafetyCarPeriods of *PacketSessionData
func (data *PacketSessionData) SetNumVirtualSafetyCarPeriods(v uint8) {
	data.NumVirtualSafetyCarPeriods = v
}

// GetNumRedFlagPeriods returns the NumRedFlagPeriods of *PacketSessionData
func (data *PacketSessionData) GetNumRedFlagPeriods() uint8 { return data.NumRedFlagPeriods }

// SetNumRedFlagPeriods stores the NumRedFlagPeriods of *PacketSessionData
func (data *PacketSessionData) SetNumRedFlagPeriods(v uint8) { data.NumRedFlagPeriods = v }

// GetEqualCarPerformance returns the EqualCarPerformance of *PacketSessionData
func (data *PacketSessionData) GetEqualCarPerformance() uint8 { return data.EqualCarPerformance }

// SetEqualCarPerformance stores the EqualCarPerformance of *PacketSessionData
func (data *PacketSessionData) SetEqualCarPerformance(v uint8) { data.EqualCarPerformance = v }

// GetRecoveryMode returns the RecoveryMode of *PacketSessionData
func (data *PacketSessionData) GetRecoveryMode() uint8 { return data.RecoveryMode }

// SetRecoveryMode stores the RecoveryMode of *PacketSessionData
func (data *PacketSessionData) SetRecoveryMode(v uint8) { data.RecoveryMode = v }

// GetFlashbackLimit returns the FlashbackLimit of *PacketSessionData
func (data *PacketSessionData) GetFlashbackLimit() uint8 { return data.FlashbackLimit }

// SetFlashbackLimit stores the FlashbackLimit of *PacketSessionData
func (data *PacketSessionData) SetFlashbackLimit(v uint8) { data.FlashbackLimit = v }

// GetSurfaceType returns the SurfaceType of *PacketSessionData
func (data *PacketSessionData) GetSurfaceType() uint8 { return data.SurfaceType }

// SetSurfaceType stores the SurfaceType of *PacketSessionData
func (data *PacketSessionData) SetSurfaceType(v uint8) { data.SurfaceType = v }

// GetLowFuelMode returns the LowFuelMode of *PacketSessionData
func (data *PacketSessionData) GetLowFuelMode() uint8 { return data.LowFuelMode }

// SetLowFuelMode stores the LowFuelMode of *PacketSessionData
func (data *PacketSessionData) SetLowFuelMode(v uint8) { data.LowFuelMode = v }

// GetRaceStarts returns the RaceStarts of *PacketSessionData
func (data *PacketSessionData) GetRaceStarts() uint8 { return data.RaceStarts }

// SetRaceStarts stores the RaceStarts of *PacketSessionData
func (data *PacketSessionData) SetRaceStarts(v uint8) { data.RaceStarts = v }

// GetTyreTemperature returns the TyreTemperature of *PacketSessionData
func (data *PacketSessionData) GetTyreTemperature() uint8 { return data.TyreTemperature }

// SetTyreTemperature stores the TyreTemperature of *PacketSessionData
func (data *PacketSessionData) SetTyreTemperature(v uint8) { data.TyreTemperature = v }

// GetPitLaneTyreSim returns the PitLaneTyreSim of *PacketSessionData
func (data *PacketSessionData) GetPitLaneTyreSim() uint8 { return data.PitLaneTyreSim }

// SetPitLaneTyreSim stores the PitLaneTyreSim of *PacketSessionData
func (data *PacketSessionData) SetPitLaneTyreSim(v uint8) { data.PitLaneTyreSim = v }

// GetCarDamage returns the CarDamage of *PacketSessionData
func (data *PacketSessionData) GetCarDamage() uint8 { return data.CarDamage }

// SetCarDamage stores the CarDamage of *PacketSessionData
func (data *PacketSessionData) SetCarDamage(v uint8) { data.CarDamage = v }

// GetCarDamageRate returns the CarDamageRate of *PacketSessionData
func (data *PacketSessionData) GetCarDamageRate() uint8 { return data.CarDamageRate }

// SetCarDamageRate stores the CarDamageRate of *PacketSessionData
func (data *PacketSessionData) SetCarDamageRate(v uint8) { data.CarDamageRate = v }

// GetCollisions returns the Collisions of *PacketSessionData
func (data *PacketSessionData) GetCollisions() uint8 { return data.Collisions }

// SetCollisions stores the Collisions of *PacketSessionData
func (data *PacketSessionData) SetCollisions(v uint8) { data.Collisions = v }

// GetCollisionsOffForFirstLapOnly returns the CollisionsOffForFirstLapOnly of *PacketSessionData
func (data *PacketSessionData) GetCollisionsOffForFirstLapOnly() uint8 {
	return data.CollisionsOffForFirstLapOnly
}

// SetCollisionsOffForFirstLapOnly stores the CollisionsOffForFirstLapOnly of *PacketSessionData
func (data *PacketSessionData) SetCollisionsOffForFirstLapOnly(v uint8) {
	data.CollisionsOffForFirstLapOnly = v
}

// GetMpUnsafePitRelease returns the MpUnsafePitRelease of *PacketSessionData
func (data *PacketSessionData) GetMpUnsafePitRelease() uint8 { return data.MpUnsafePitRelease }

// SetMpUnsafePitRelease stores the MpUnsafePitRelease of *PacketSessionData
func (data *PacketSessionData) SetMpUnsafePitRelease(v uint8) { data.MpUnsafePitRelease = v }

// GetMpOffForGriefing returns the MpOffForGriefing of *PacketSessionData
func (data *PacketSessionData) GetMpOffForGriefing() uint8 { return data.MpOffForGriefing }

// SetMpOffForGriefing stores the MpOffForGriefing of *PacketSessionData
func (data *PacketSessionData) SetMpOffForGriefing(v uint8) { data.MpOffForGriefing = v }

// GetCornerCuttingStringency returns the CornerCuttingStringency of *PacketSessionData
func (data *PacketSessionData) GetCornerCuttingStringency() uint8 {
	return data.CornerCuttingStringency
}

// SetCornerCuttingStringency stores the CornerCuttingStringency of *PacketSessionData
func (data *PacketSessionData) SetCornerCuttingStringency(v uint8) { data.CornerCuttingStringency = v }

// GetParcFermeRules returns the ParcFermeRules of *PacketSessionData
func (data *PacketSessionData) GetParcFermeRules() uint8 { return data.ParcFermeRules }

// SetParcFermeRules stores the ParcFermeRules of *PacketSessionData
func (data *PacketSessionData) SetParcFermeRules(v uint8) { data.ParcFermeRules = v }

// GetPitStopExperience returns the PitStopExperience of *PacketSessionData
func (data *PacketSessionData) GetPitStopExperience() uint8 { return data.PitStopExperience }

// SetPitStopExperience stores the PitStopExperience of *PacketSessionData
func (data *PacketSessionData) SetPitStopExperience(v uint8) { data.PitStopExperience = v }

// GetSafetyCar returns the SafetyCar of *PacketSessionData
func (data *PacketSessionData) GetSafetyCar() uint8 { return data.SafetyCar }

// SetSafetyCar stores the SafetyCar of *PacketSessionData
func (data *PacketSessionData) SetSafetyCar(v uint8) { data.SafetyCar = v }

// GetSafetyCarExperience returns the SafetyCarExperience of *PacketSessionData
func (data *PacketSessionData) GetSafetyCarExperience() uint8 { return data.SafetyCarExperience }

// SetSafetyCarExperience stores the SafetyCarExperience of *PacketSessionData
func (data *PacketSessionData) SetSafetyCarExperience(v uint8) { data.SafetyCarExperience = v }

// GetFormationLap returns the FormationLap of *PacketSessionData
func (data *PacketSessionData) GetFormationLap() uint8 { return data.FormationLap }

// SetFormationLap stores the FormationLap of *PacketSessionData
func (data *PacketSessionData) SetFormationLap(v uint8) { data.FormationLap = v }

// GetFormationLapExperience returns the FormationLapExperience of *PacketSessionData
func (data *PacketSessionData) GetFormationLapExperience() uint8 { return data.FormationLapExperience }

// SetFormationLapExperience stores the FormationLapExperience of *PacketSessionData
func (data *PacketSessionData) SetFormationLapExperience(v uint8) { data.FormationLapExperience = v }

// GetRedFlags returns the RedFlags of *PacketSessionData
func (data *PacketSessionData) GetRedFlags() uint8 { return data.RedFlags }

// SetRedFlags stores the RedFlags of *PacketSessionData
func (data *PacketSessionData) SetRedFlags(v uint8) { data.RedFlags = v }

// GetAffectsLicenceLevelSolo returns the AffectsLicenceLevelSolo of *PacketSessionData
func (data *PacketSessionData) GetAffectsLicenceLevelSolo() uint8 {
	return data.AffectsLicenceLevelSolo
}

// SetAffectsLicenceLevelSolo stores the AffectsLicenceLevelSolo of *PacketSessionData
func (data *PacketSessionData) SetAffectsLicenceLevelSolo(v uint8) { data.AffectsLicenceLevelSolo = v }

// GetAffectsLicenceLevelMP returns the AffectsLicenceLevelMP of *PacketSessionData
func (data *PacketSessionData) GetAffectsLicenceLevelMP() uint8 { return data.AffectsLicenceLevelMP }

// SetAffectsLicenceLevelMP stores the AffectsLicenceLevelMP of *PacketSessionData
func (data *PacketSessionData) SetAffectsLicenceLevelMP(v uint8) { data.AffectsLicenceLevelMP = v }

// GetNumSessionsInWeekend returns the NumSessionsInWeekend of *PacketSessionData
func (data *PacketSessionData) GetNumSessionsInWeekend() uint8 { return data.NumSessionsInWeekend }

// SetNumSessionsInWeekend stores the NumSessionsInWeekend of *PacketSessionData
func (data *PacketSessionData) SetNumSessionsInWeekend(v uint8) { data.NumSessionsInWeekend = v }

// GetWeekendStructure returns the WeekendStructure of *PacketSessionData
func (data *PacketSessionData) GetWeekendStructure() [12]uint8 { return data.WeekendStructure }

// SetWeekendStructure stores the WeekendStructure of *PacketSessionData
func (data *PacketSessionData) SetWeekendStructure(v [12]uint8) { data.WeekendStructure = v }

// GetSector2LapDistanceStart returns the Sector2LapDistanceStart of *PacketSessionData
func (data *PacketSessionData) GetSector2LapDistanceStart() float32 {
	return data.Sector2LapDistanceStart
}

// SetSector2LapDistanceStart stores the Sector2LapDistanceStart of *PacketSessionData
func (data *PacketSessionData) SetSector2LapDistanceStart(v float32) {
	data.Sector2LapDistanceStart = v
}

// GetSector3LapDistanceStart returns the Sector3LapDistanceStart of *PacketSessionData
func (data *PacketSessionData) GetSector3LapDistanceStart() float32 {
	return data.Sector3LapDistanceStart
}

// SetSector3LapDistanceStart stores the Sector3LapDistanceStart of *PacketSessionData
func (data *PacketSessionData) SetSector3LapDistanceStart(v float32) {
	data.Sector3LapDistanceStart = v
}

// GetActiveAeroTrackStatus returns the ActiveAeroTrackStatus of *PacketSessionData
func (data *PacketSessionData) GetActiveAeroTrackStatus() uint8 { return data.ActiveAeroTrackStatus }

// SetActiveAeroTrackStatus stores the ActiveAeroTrackStatus of *PacketSessionData
func (data *PacketSessionData) SetActiveAeroTrackStatus(v uint8) { data.ActiveAeroTrackStatus = v }

// GetNumActiveAeroZonesFull returns the NumActiveAeroZonesFull of *PacketSessionData
func (data *PacketSessionData) GetNumActiveAeroZonesFull() uint8 { return data.NumActiveAeroZonesFull }

// SetNumActiveAeroZonesFull stores the NumActiveAeroZonesFull of *PacketSessionData
func (data *PacketSessionData) SetNumActiveAeroZonesFull(v uint8) { data.NumActiveAeroZonesFull = v }

// GetActiveAeroZonesFull returns the ActiveAeroZonesFull of *PacketSessionData
func (data *PacketSessionData) GetActiveAeroZonesFull() [8]ActiveAeroZone {
	return data.ActiveAeroZonesFull
}

// SetActiveAeroZonesFull stores the ActiveAeroZonesFull of *PacketSessionData
func (data *PacketSessionData) SetActiveAeroZonesFull(v [8]ActiveAeroZone) {
	data.ActiveAeroZonesFull = v
}

// GetNumActiveAeroZonesPartial returns the NumActiveAeroZonesPartial of *PacketSessionData
func (data *PacketSessionData) GetNumActiveAeroZonesPartial() uint8 {
	return data.NumActiveAeroZonesPartial
}

// SetNumActiveAeroZonesPartial stores the NumActiveAeroZonesPartial of *PacketSessionData
func (data *PacketSessionData) SetNumActiveAeroZonesPartial(v uint8) {
	data.NumActiveAeroZonesPartial = v
}

// GetActiveAeroZonesPartial returns the ActiveAeroZonesPartial of *PacketSessionData
func (data *PacketSessionData) GetActiveAeroZonesPartial() [8]ActiveAeroZone {
	return data.ActiveAeroZonesPartial
}

// SetActiveAeroZonesPartial stores the ActiveAeroZonesPartial of *PacketSessionData
func (data *PacketSessionData) SetActiveAeroZonesPartial(v [8]ActiveAeroZone) {
	data.ActiveAeroZonesPartial = v
}

// GetNumDRSZones returns the NumDRSZones of *PacketSessionData
func (data *PacketSessionData) GetNumDRSZones() uint8 { return data.NumDRSZones }

// SetNumDRSZones stores the NumDRSZones of *PacketSessionData
func (data *PacketSessionData) SetNumDRSZones(v uint8) { data.NumDRSZones = v }

// GetDrsZones returns the DrsZones of *PacketSessionData
func (data *PacketSessionData) GetDrsZones() [4]DRSZone { return data.DrsZones }

// SetDrsZones stores the DrsZones of *PacketSessionData
func (data *PacketSessionData) SetDrsZones(v [4]DRSZone) { data.DrsZones = v }

// GetStartReactionTime returns the StartReactionTime of *PacketSessionData
func (data *PacketSessionData) GetStartReactionTime() float32 { return data.StartReactionTime }

// SetStartReactionTime stores the StartReactionTime of *PacketSessionData
func (data *PacketSessionData) SetStartReactionTime(v float32) { data.StartReactionTime = v }

// GetAntiLockBrakesAssist returns the AntiLockBrakesAssist of *PacketSessionData
func (data *PacketSessionData) GetAntiLockBrakesAssist() uint8 { return data.AntiLockBrakesAssist }

// SetAntiLockBrakesAssist stores the AntiLockBrakesAssist of *PacketSessionData
func (data *PacketSessionData) SetAntiLockBrakesAssist(v uint8) { data.AntiLockBrakesAssist = v }

// GetTractionControlAssist returns the TractionControlAssist of *PacketSessionData
func (data *PacketSessionData) GetTractionControlAssist() uint8 { return data.TractionControlAssist }

// SetTractionControlAssist stores the TractionControlAssist of *PacketSessionData
func (data *PacketSessionData) SetTractionControlAssist(v uint8) { data.TractionControlAssist = v }

// GetDynamicRacingLineHiVis returns the DynamicRacingLineHiVis of *PacketSessionData
func (data *PacketSessionData) GetDynamicRacingLineHiVis() uint8 { return data.DynamicRacingLineHiVis }

// SetDynamicRacingLineHiVis stores the DynamicRacingLineHiVis of *PacketSessionData
func (data *PacketSessionData) SetDynamicRacingLineHiVis(v uint8) { data.DynamicRacingLineHiVis = v }

// GetDynamicRacingLineColourBlind returns the DynamicRacingLineColourBlind of *PacketSessionData
func (data *PacketSessionData) GetDynamicRacingLineColourBlind() uint8 {
	return data.DynamicRacingLineColourBlind
}

// SetDynamicRacingLineColourBlind stores the DynamicRacingLineColourBlind of *PacketSessionData
func (data *PacketSessionData) SetDynamicRacingLineColourBlind(v uint8) {
	data.DynamicRacingLineColourBlind = v
}

// GetRecurringRewindPrompt returns the RecurringRewindPrompt of *PacketSessionData
func (data *PacketSessionData) GetRecurringRewindPrompt() uint8 { return data.RecurringRewindPrompt }

// SetRecurringRewindPrompt stores the RecurringRewindPrompt of *PacketSessionData
func (data *PacketSessionData) SetRecurringRewindPrompt(v uint8) { data.RecurringRewindPrompt = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketSessionData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	data.Weather = reader.ReadUint8()
	data.TrackTemperature = reader.ReadInt8()
	data.AirTemperature = reader.ReadInt8()
	data.TotalLaps = reader.ReadUint8()
	data.TrackLength = reader.ReadUint16()
	data.SessionType = reader.ReadUint8()
	data.TrackId = reader.ReadInt8()
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
	data.GameMode = reader.ReadUint8()
	data.RuleSet = reader.ReadUint8()
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
	data.ActiveAeroTrackStatus = reader.ReadUint8()
	data.NumActiveAeroZonesFull = reader.ReadUint8()

	for i := range data.ActiveAeroZonesFull {
		data.ActiveAeroZonesFull[i].Parse(reader)
	}
	data.NumActiveAeroZonesPartial = reader.ReadUint8()

	for i := range data.ActiveAeroZonesPartial {
		data.ActiveAeroZonesPartial[i].Parse(reader)
	}

	data.NumDRSZones = reader.ReadUint8()

	for i := range data.DrsZones {
		data.DrsZones[i].Parse(reader)
	}

	data.StartReactionTime = reader.ReadFloat32()
	data.AntiLockBrakesAssist = reader.ReadUint8()
	data.TractionControlAssist = reader.ReadUint8()
	data.DynamicRacingLineHiVis = reader.ReadUint8()
	data.DynamicRacingLineColourBlind = reader.ReadUint8()
	data.RecurringRewindPrompt = reader.ReadUint8()

}
