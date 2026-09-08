package f12026s

import xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"

type Packet interface {
	Parse(header *PacketHeader, reader *xbinary.LittleEndianReader)
	// GetHeader returns the Header of packet
	GetHeader() PacketHeader
	// GetPacketID returns the identification code of this packet
	GetPacketID() PacketID
}
