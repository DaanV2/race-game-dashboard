package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketMotionData -ignore-field CarMotionData

type PacketMotionData struct {
	Header        PacketHeader                   // Header
	CarMotionData [CS_MAX_NUM_CARS]CarMotionData // Data for all cars on track
}

func (data *PacketMotionData) GetPacketID() PacketID { return PACKET_ID_MOTION }

// GetCarMotionData returns the CarMotionData of *PacketMotionData
func (data *PacketMotionData) GetCarMotionData(car int) CarMotionData { return data.CarMotionData[car] }

// SetCarMotionData stores the CarMotionData of *PacketMotionData
func (data *PacketMotionData) SetCarMotionData(car int, v CarMotionData) { data.CarMotionData[car] = v }

func (data *PacketMotionData) GetPlayerData() CarMotionData {
	return data.CarMotionData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketMotionData) GetSecondPlayerData() CarMotionData {
	return data.CarMotionData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketMotionData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarMotionData {
		data.CarMotionData[i].Parse(reader)
	}
}
