package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type LobbyInfoData struct {
	AiControlled    uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId          uint16   // Team id - see appendix (65535 if no team currently selected)
	Nationality     uint8    // Nationality of the driver
	Platform        uint8    // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	Name            [32]byte // Name of participant in UTF-8 format – null terminated  Will be truncated with ... (U+2026) if too long
	CarNumber       uint8    // Car number of the player
	YourTelemetry   uint8    // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8    // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16   // F1 World tech level
	ReadyStatus     uint8    // 0 = not ready, 1 = ready, 2 = spectating
}

// GetAiControlled returns the AiControlled of *LobbyInfoData
func (data *LobbyInfoData) GetAiControlled() uint8 { return data.AiControlled }

// SetAiControlled stores the AiControlled of *LobbyInfoData
func (data *LobbyInfoData) SetAiControlled(v uint8) { data.AiControlled = v }

// GetTeamId returns the TeamId of *LobbyInfoData
func (data *LobbyInfoData) GetTeamId() uint16 { return data.TeamId }

// SetTeamId stores the TeamId of *LobbyInfoData
func (data *LobbyInfoData) SetTeamId(v uint16) { data.TeamId = v }

// GetNationality returns the Nationality of *LobbyInfoData
func (data *LobbyInfoData) GetNationality() uint8 { return data.Nationality }

// SetNationality stores the Nationality of *LobbyInfoData
func (data *LobbyInfoData) SetNationality(v uint8) { data.Nationality = v }

// GetPlatform returns the Platform of *LobbyInfoData
func (data *LobbyInfoData) GetPlatform() uint8 { return data.Platform }

// SetPlatform stores the Platform of *LobbyInfoData
func (data *LobbyInfoData) SetPlatform(v uint8) { data.Platform = v }

// GetName returns the Name of *LobbyInfoData
func (data *LobbyInfoData) GetName() [32]byte { return data.Name }

// SetName stores the Name of *LobbyInfoData
func (data *LobbyInfoData) SetName(v [32]byte) { data.Name = v }

// GetCarNumber returns the CarNumber of *LobbyInfoData
func (data *LobbyInfoData) GetCarNumber() uint8 { return data.CarNumber }

// SetCarNumber stores the CarNumber of *LobbyInfoData
func (data *LobbyInfoData) SetCarNumber(v uint8) { data.CarNumber = v }

// GetYourTelemetry returns the YourTelemetry of *LobbyInfoData
func (data *LobbyInfoData) GetYourTelemetry() uint8 { return data.YourTelemetry }

// SetYourTelemetry stores the YourTelemetry of *LobbyInfoData
func (data *LobbyInfoData) SetYourTelemetry(v uint8) { data.YourTelemetry = v }

// GetShowOnlineNames returns the ShowOnlineNames of *LobbyInfoData
func (data *LobbyInfoData) GetShowOnlineNames() uint8 { return data.ShowOnlineNames }

// SetShowOnlineNames stores the ShowOnlineNames of *LobbyInfoData
func (data *LobbyInfoData) SetShowOnlineNames(v uint8) { data.ShowOnlineNames = v }

// GetTechLevel returns the TechLevel of *LobbyInfoData
func (data *LobbyInfoData) GetTechLevel() uint16 { return data.TechLevel }

// SetTechLevel stores the TechLevel of *LobbyInfoData
func (data *LobbyInfoData) SetTechLevel(v uint16) { data.TechLevel = v }

// GetReadyStatus returns the ReadyStatus of *LobbyInfoData
func (data *LobbyInfoData) GetReadyStatus() uint8 { return data.ReadyStatus }

// SetReadyStatus stores the ReadyStatus of *LobbyInfoData
func (data *LobbyInfoData) SetReadyStatus(v uint8) { data.ReadyStatus = v }

func (data *LobbyInfoData) Parse(reader *xbinary.LittleEndianReader) {
	data.AiControlled = reader.ReadUint8()
	data.TeamId = reader.ReadUint16()
	data.Nationality = reader.ReadUint8()
	data.Platform = reader.ReadUint8()
	data.Name = xbinary.Readx32(reader.ReadByte)
	data.CarNumber = reader.ReadUint8()
	data.YourTelemetry = reader.ReadUint8()
	data.ShowOnlineNames = reader.ReadUint8()
	data.TechLevel = reader.ReadUint16()
	data.ReadyStatus = reader.ReadUint8()

}
