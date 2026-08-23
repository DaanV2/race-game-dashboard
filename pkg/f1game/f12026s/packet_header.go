package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketHeader struct {
	PacketFormat            uint16   // 2026
	GameYear                uint8    // Game year - last two digits e.g. 25
	GameMajorVersion        uint8    // Game major version - "X.00"
	GameMinorVersion        uint8    // Game minor version - "1.XX"
	PacketVersion           uint8    // Version of this packet type, all start from 1
	PacketId                PacketID // Identifier for the packet type, see below
	SessionUID              uint64   // Unique identifier for the session
	SessionTime             float32  // Session timestamp
	FrameIdentifier         uint32   // Identifier for the frame the data was retrieved on
	OverallFrameIdentifier  uint32   // Overall identifier for the frame the data was retrieved  on, doesn't go back after flashbacks
	PlayerCarIndex          uint8    // Index of player's car in the array
	SecondaryPlayerCarIndex uint8    // Index of secondary player's car in the array (splitscreen)  255 if no second player
}

// GetPacketFormat returns the PacketFormat of *PacketHeader
func (data *PacketHeader) GetPacketFormat() uint16 { return data.PacketFormat }

// SetPacketFormat stores the PacketFormat of *PacketHeader
func (data *PacketHeader) SetPacketFormat(v uint16) { data.PacketFormat = v }

// GetGameYear returns the GameYear of *PacketHeader
func (data *PacketHeader) GetGameYear() uint8 { return data.GameYear }

// SetGameYear stores the GameYear of *PacketHeader
func (data *PacketHeader) SetGameYear(v uint8) { data.GameYear = v }

// GetGameMajorVersion returns the GameMajorVersion of *PacketHeader
func (data *PacketHeader) GetGameMajorVersion() uint8 { return data.GameMajorVersion }

// SetGameMajorVersion stores the GameMajorVersion of *PacketHeader
func (data *PacketHeader) SetGameMajorVersion(v uint8) { data.GameMajorVersion = v }

// GetGameMinorVersion returns the GameMinorVersion of *PacketHeader
func (data *PacketHeader) GetGameMinorVersion() uint8 { return data.GameMinorVersion }

// SetGameMinorVersion stores the GameMinorVersion of *PacketHeader
func (data *PacketHeader) SetGameMinorVersion(v uint8) { data.GameMinorVersion = v }

// GetPacketVersion returns the PacketVersion of *PacketHeader
func (data *PacketHeader) GetPacketVersion() uint8 { return data.PacketVersion }

// SetPacketVersion stores the PacketVersion of *PacketHeader
func (data *PacketHeader) SetPacketVersion(v uint8) { data.PacketVersion = v }

// GetPacketId returns the PacketId of *PacketHeader
func (data *PacketHeader) GetPacketId() PacketID { return data.PacketId }

// SetPacketId stores the PacketId of *PacketHeader
func (data *PacketHeader) SetPacketId(v PacketID) { data.PacketId = v }

// GetSessionUID returns the SessionUID of *PacketHeader
func (data *PacketHeader) GetSessionUID() uint64 { return data.SessionUID }

// SetSessionUID stores the SessionUID of *PacketHeader
func (data *PacketHeader) SetSessionUID(v uint64) { data.SessionUID = v }

// GetSessionTime returns the SessionTime of *PacketHeader
func (data *PacketHeader) GetSessionTime() float32 { return data.SessionTime }

// SetSessionTime stores the SessionTime of *PacketHeader
func (data *PacketHeader) SetSessionTime(v float32) { data.SessionTime = v }

// GetFrameIdentifier returns the FrameIdentifier of *PacketHeader
func (data *PacketHeader) GetFrameIdentifier() uint32 { return data.FrameIdentifier }

// SetFrameIdentifier stores the FrameIdentifier of *PacketHeader
func (data *PacketHeader) SetFrameIdentifier(v uint32) { data.FrameIdentifier = v }

// GetOverallFrameIdentifier returns the OverallFrameIdentifier of *PacketHeader
func (data *PacketHeader) GetOverallFrameIdentifier() uint32 { return data.OverallFrameIdentifier }

// SetOverallFrameIdentifier stores the OverallFrameIdentifier of *PacketHeader
func (data *PacketHeader) SetOverallFrameIdentifier(v uint32) { data.OverallFrameIdentifier = v }

// GetPlayerCarIndex returns the PlayerCarIndex of *PacketHeader
func (data *PacketHeader) GetPlayerCarIndex() uint8 { return data.PlayerCarIndex }

// SetPlayerCarIndex stores the PlayerCarIndex of *PacketHeader
func (data *PacketHeader) SetPlayerCarIndex(v uint8) { data.PlayerCarIndex = v }

// GetSecondaryPlayerCarIndex returns the SecondaryPlayerCarIndex of *PacketHeader
func (data *PacketHeader) GetSecondaryPlayerCarIndex() uint8 { return data.SecondaryPlayerCarIndex }

// SetSecondaryPlayerCarIndex stores the SecondaryPlayerCarIndex of *PacketHeader
func (data *PacketHeader) SetSecondaryPlayerCarIndex(v uint8) { data.SecondaryPlayerCarIndex = v }

func (data *PacketHeader) Parse(reader *xbinary.LittleEndianReader) {
	data.PacketFormat = reader.ReadUint16()
	data.GameYear = reader.ReadUint8()
	data.GameMajorVersion = reader.ReadUint8()
	data.GameMinorVersion = reader.ReadUint8()
	data.PacketVersion = reader.ReadUint8()
	data.PacketId = PacketID(reader.ReadUint8())
	data.SessionUID = reader.ReadUint64()
	data.SessionTime = reader.ReadFloat32()
	data.FrameIdentifier = reader.ReadUint32()
	data.OverallFrameIdentifier = reader.ReadUint32()
	data.PlayerCarIndex = reader.ReadUint8()
	data.SecondaryPlayerCarIndex = reader.ReadUint8()
}
