package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarStatusData struct {
	Header        PacketHeader                   // Header
	CarStatusData [CS_MAX_NUM_CARS]CarStatusData //
}

// GetPacketID returns the identification of this packet
func (data *PacketCarStatusData) GetPacketID() PacketID { return PACKET_ID_CAR_STATUS }

// GetHeader returns the Header of *PacketCarStatusData
func (data *PacketCarStatusData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarStatusData
func (data *PacketCarStatusData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarStatusData returns the CarStatusData of *PacketCarStatusData
func (data *PacketCarStatusData) GetCarStatusData(car int) CarStatusData {
	return data.CarStatusData[car]
}

// SetCarStatusData stores the CarStatusData of *PacketCarStatusData
func (data *PacketCarStatusData) SetCarStatusData(car int, v CarStatusData) {
	data.CarStatusData[car] = v
}

func (data *PacketCarStatusData) GetPlayerData() CarStatusData {
	carIndex := data.Header.GetPlayerCarIndex()

	return data.CarStatusData[carIndex]
}

func (data *PacketCarStatusData) GetSecondPlayerData() CarStatusData {
	carIndex := data.Header.GetSecondaryPlayerCarIndex()

	return data.CarStatusData[carIndex]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarStatusData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarStatusData {
		data.CarStatusData[i].Parse(reader)
	}
}
