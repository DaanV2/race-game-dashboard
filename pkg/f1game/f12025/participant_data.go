package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	xstrings "github.com/daanv2/race-game-dashboard/pkg/extensions/strings"
)

type ParticipantData struct {
	AiControlled    uint8                             // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId        uint16                            // Driver id - see appendix, 65535 if network human
	NetworkId       uint16                            // Network id – unique identifier for network players
	TeamId          uint16                            // Team id - see appendix
	MyTeam          uint8                             // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber      uint8                             // Race number of the car
	Nationality     uint8                             // Nationality of the driver
	Name            [CS_MAX_PARTICIPANT_NAME_LEN]byte // Name of participant in UTF-8 format – null terminated  Will be truncated with … (U+2026) if too long
	YourTelemetry   uint8                             // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8                             // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16                            // F1 World tech level
	Platform        uint8                             // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	NumColours      uint8                             // Number of colours valid for this car
	LiveryColours   [4]LiveryColour                   // Colours for the car
}

// GetAiControlled returns the AiControlled of *ParticipantData
func (data *ParticipantData) GetAiControlled() uint8 { return data.AiControlled }

// SetAiControlled stores the AiControlled of *ParticipantData
func (data *ParticipantData) SetAiControlled(v uint8) { data.AiControlled = v }

// GetDriverId returns the DriverId of *ParticipantData
func (data *ParticipantData) GetDriverId() uint16 { return data.DriverId }

// SetDriverId stores the DriverId of *ParticipantData
func (data *ParticipantData) SetDriverId(v uint16) { data.DriverId = v }

// GetNetworkId returns the NetworkId of *ParticipantData
func (data *ParticipantData) GetNetworkId() uint16 { return data.NetworkId }

// SetNetworkId stores the NetworkId of *ParticipantData
func (data *ParticipantData) SetNetworkId(v uint16) { data.NetworkId = v }

// GetTeamId returns the TeamId of *ParticipantData
func (data *ParticipantData) GetTeamId() uint16 { return data.TeamId }

// SetTeamId stores the TeamId of *ParticipantData
func (data *ParticipantData) SetTeamId(v uint16) { data.TeamId = v }

// GetMyTeam returns the MyTeam of *ParticipantData
func (data *ParticipantData) GetMyTeam() uint8 { return data.MyTeam }

// SetMyTeam stores the MyTeam of *ParticipantData
func (data *ParticipantData) SetMyTeam(v uint8) { data.MyTeam = v }

// GetRaceNumber returns the RaceNumber of *ParticipantData
func (data *ParticipantData) GetRaceNumber() uint8 { return data.RaceNumber }

// SetRaceNumber stores the RaceNumber of *ParticipantData
func (data *ParticipantData) SetRaceNumber(v uint8) { data.RaceNumber = v }

// GetNationality returns the Nationality of *ParticipantData
func (data *ParticipantData) GetNationality() uint8 { return data.Nationality }

// SetNationality stores the Nationality of *ParticipantData
func (data *ParticipantData) SetNationality(v uint8) { data.Nationality = v }

// GetName returns the Name of *LobbyInfoData
func (data *ParticipantData) GetName() string { return xstrings.NullTerminated(data.Name[:]) }

// SetName stores the Name of *LobbyInfoData
func (data *ParticipantData) SetName(v string) {
	var result [CS_MAX_PARTICIPANT_NAME_LEN]byte
	b := []byte(v)

	copy(result[:], b)
	data.Name = result
}

// GetYourTelemetry returns the YourTelemetry of *ParticipantData
func (data *ParticipantData) GetYourTelemetry() uint8 { return data.YourTelemetry }

// SetYourTelemetry stores the YourTelemetry of *ParticipantData
func (data *ParticipantData) SetYourTelemetry(v uint8) { data.YourTelemetry = v }

// GetShowOnlineNames returns the ShowOnlineNames of *ParticipantData
func (data *ParticipantData) GetShowOnlineNames() uint8 { return data.ShowOnlineNames }

// SetShowOnlineNames stores the ShowOnlineNames of *ParticipantData
func (data *ParticipantData) SetShowOnlineNames(v uint8) { data.ShowOnlineNames = v }

// GetTechLevel returns the TechLevel of *ParticipantData
func (data *ParticipantData) GetTechLevel() uint16 { return data.TechLevel }

// SetTechLevel stores the TechLevel of *ParticipantData
func (data *ParticipantData) SetTechLevel(v uint16) { data.TechLevel = v }

// GetPlatform returns the Platform of *ParticipantData
func (data *ParticipantData) GetPlatform() uint8 { return data.Platform }

// SetPlatform stores the Platform of *ParticipantData
func (data *ParticipantData) SetPlatform(v uint8) { data.Platform = v }

// GetNumColours returns the NumColours of *ParticipantData
func (data *ParticipantData) GetNumColours() uint8 { return data.NumColours }

// SetNumColours stores the NumColours of *ParticipantData
func (data *ParticipantData) SetNumColours(v uint8) { data.NumColours = v }

// GetLiveryColours returns the LiveryColours of *ParticipantData
func (data *ParticipantData) GetLiveryColours() [4]LiveryColour { return data.LiveryColours }

// SetLiveryColours stores the LiveryColours of *ParticipantData
func (data *ParticipantData) SetLiveryColours(v [4]LiveryColour) { data.LiveryColours = v }

func (data *ParticipantData) Parse(reader *xbinary.LittleEndianReader) {
	data.AiControlled = reader.ReadUint8()
	data.DriverId = reader.ReadUint16()
	data.NetworkId = reader.ReadUint16()
	data.TeamId = reader.ReadUint16()
	data.MyTeam = reader.ReadUint8()
	data.RaceNumber = reader.ReadUint8()
	data.Nationality = reader.ReadUint8()
	data.Name = xbinary.Readx32(reader.ReadByte)
	data.YourTelemetry = reader.ReadUint8()
	data.ShowOnlineNames = reader.ReadUint8()
	data.TechLevel = reader.ReadUint16()
	data.Platform = reader.ReadUint8()
	data.NumColours = reader.ReadUint8()

	for i := range data.LiveryColours {
		data.LiveryColours[i].Parse(reader)
	}
}
