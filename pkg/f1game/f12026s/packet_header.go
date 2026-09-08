package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketHeader -ignore-field ClassificationData

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
	SecondaryPlayerCarIndex uint8    // Index of secondary player's car in the array (splitscreen), 255 if no second player
}

// HasSecondaryPlayerCar returns if there is a second player active
func (data *PacketHeader) HasSecondaryPlayerCar() bool { return data.SecondaryPlayerCarIndex != 255 }

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
