package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	xstrings "github.com/daanv2/race-game-dashboard/pkg/extensions/strings"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type LobbyInfoData -ignore-field Name

type LobbyInfoData struct {
	AiControlled    uint8                             // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId          f1common.TeamID                   // Team id (uint16) - see appendix (65535 if no team currently selected)
	Nationality     f1common.NationalityID            // Nationality (uint8) of the driver
	Platform        uint8                             // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	Name            [CS_MAX_PARTICIPANT_NAME_LEN]byte // Name of participant in UTF-8 format – null terminated  Will be truncated with ... (U+2026) if too long
	CarNumber       uint8                             // Car number of the player
	YourTelemetry   uint8                             // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8                             // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16                            // F1 World tech level
	ReadyStatus     uint8                             // 0 = not ready, 1 = ready, 2 = spectating
}

// GetName returns the Name of *LobbyInfoData
func (data *LobbyInfoData) GetName() string { return xstrings.NullTerminated(data.Name[:]) }

// SetName stores the Name of *LobbyInfoData
func (data *LobbyInfoData) SetName(v string) {
	var result [CS_MAX_PARTICIPANT_NAME_LEN]byte
	b := []byte(v)

	copy(result[:], b)
	data.Name = result
}

func (data *LobbyInfoData) Parse(reader *xbinary.LittleEndianReader) {
	data.AiControlled = reader.ReadUint8()
	data.TeamId = f1common.TeamID(reader.ReadUint16())
	data.Nationality = f1common.NationalityID(reader.ReadUint8())
	data.Platform = reader.ReadUint8()
	reader.Read(data.Name[:])
	data.CarNumber = reader.ReadUint8()
	data.YourTelemetry = reader.ReadUint8()
	data.ShowOnlineNames = reader.ReadUint8()
	data.TechLevel = reader.ReadUint16()
	data.ReadyStatus = reader.ReadUint8()
}
