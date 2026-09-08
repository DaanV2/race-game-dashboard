package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketCarStatusData -ignore-field CarStatusData

type PacketCarStatusData struct {
	Header        PacketHeader                   // Header
	CarStatusData [CS_MAX_NUM_CARS]CarStatusData //
}

// GetPacketID returns the identification of this packet
func (data *PacketCarStatusData) GetPacketID() PacketID { return PACKET_ID_CAR_STATUS }

func (data *PacketCarStatusData) GetPlayerData() CarStatusData {
	return data.CarStatusData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketCarStatusData) GetSecondPlayerData() CarStatusData {
	return data.CarStatusData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarStatusData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarStatusData {
		data.CarStatusData[i].Parse(reader)
	}
}
