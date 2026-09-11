package udp

import xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"

var _ Handler = &GameRouter{}

type GameRouter struct {
	F12025  Handler
	F12026s Handler
}

type f1PacketHeader struct {
	PacketFormat     uint16 // 2026 / 2025
	GameYear         uint8  // Game year - last two digits e.g. 25
	GameMajorVersion uint8  // Game major version - "X.00"
	GameMinorVersion uint8  // Game minor version - "1.XX"
	PacketVersion    uint8  // Version of this packet type, all start from 1
}

func (data *f1PacketHeader) Parse(reader *xbinary.LittleEndianReader) {
	data.PacketFormat = reader.ReadUint16()
	data.GameYear = reader.ReadUint8()
	data.GameMajorVersion = reader.ReadUint8()
	data.GameMinorVersion = reader.ReadUint8()
	data.PacketVersion = reader.ReadUint8()
}

// HandleMsg implements [Handler].
func (router *GameRouter) HandleMsg(data []byte) {
	// F1?
	var overlapPacket f1PacketHeader
	reader := xbinary.NewLittleEndianReader(data)
	overlapPacket.Parse(reader)

	switch {
	// f12026s game?
	case overlapPacket.PacketFormat == 2026 && overlapPacket.PacketVersion == 1:
		if router.F12026s != nil {
			router.F12026s.HandleMsg(data)
		}
	// f12025 game?
	case overlapPacket.PacketFormat == 2025 && overlapPacket.PacketVersion == 1:
		if router.F12026s != nil {
			router.F12025.HandleMsg(data)
		}
	}
}
