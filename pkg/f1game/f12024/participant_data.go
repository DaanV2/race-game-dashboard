package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	xstrings "github.com/daanv2/race-game-dashboard/pkg/extensions/strings"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type ParticipantData -ignore-field Name

type ParticipantData struct {
	AiControlled    uint8                             // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId        f1common.DriverID                 // Driver id (uint8) - see appendix, 255 if network human
	NetworkId       uint8                             // Network id – unique identifier for network players
	TeamId          f1common.TeamID                   // Team id (uint8) - see appendix
	MyTeam          uint8                             // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber      uint8                             // Race number of the car
	Nationality     uint8                             // Nationality of the driver
	Name            [CS_MAX_PARTICIPANT_NAME_LEN]byte // Name of participant in UTF-8 format – null terminated  Will be truncated with … (U+2026) if too long
	YourTelemetry   uint8                             // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8                             // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16                            // F1 World tech level
	Platform        uint8                             // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
}

// GetName returns the Name of *LobbyInfoData
func (data *ParticipantData) GetName() string { return xstrings.NullTerminated(data.Name[:]) }

// SetName stores the Name of *LobbyInfoData
func (data *ParticipantData) SetName(v string) {
	var result [CS_MAX_PARTICIPANT_NAME_LEN]byte
	b := []byte(v)

	copy(result[:], b)
	data.Name = result
}

func (data *ParticipantData) Parse(reader *xbinary.LittleEndianReader) {
	data.AiControlled = reader.ReadUint8()
	data.DriverId = f1common.DriverID(reader.ReadUint8())
	data.NetworkId = reader.ReadUint8()
	data.TeamId = f1common.TeamID(reader.ReadUint8())
	data.MyTeam = reader.ReadUint8()
	data.RaceNumber = reader.ReadUint8()
	data.Nationality = reader.ReadUint8()
	reader.Read(data.Name[:])
	data.YourTelemetry = reader.ReadUint8()
	data.ShowOnlineNames = reader.ReadUint8()
	data.TechLevel = reader.ReadUint16()
	data.Platform = reader.ReadUint8()
}
