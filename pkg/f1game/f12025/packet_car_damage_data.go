package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketCarDamageData -ignore-field CarDamageData

type PacketCarDamageData struct {
	Header        PacketHeader                   // Header
	CarDamageData [CS_MAX_NUM_CARS]CarDamageData //
}

// GetPacketID returns the identification of this packet
func (data *PacketCarDamageData) GetPacketID() PacketID { return PACKET_ID_CAR_DAMAGE }

// GetCarDamageData returns the CarDamageData of *PacketCarDamageData
func (data *PacketCarDamageData) GetCarDamageData(car int) CarDamageData {
	return data.CarDamageData[car]
}

// SetCarDamageData stores the CarDamageData of *PacketCarDamageData
func (data *PacketCarDamageData) SetCarDamageData(car int, v CarDamageData) {
	data.CarDamageData[car] = v
}

func (data *PacketCarDamageData) GetPlayerData() CarDamageData {
	return data.CarDamageData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketCarDamageData) GetSecondPlayerData() CarDamageData {
	return data.CarDamageData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarDamageData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarDamageData {
		data.CarDamageData[i].Parse(reader)
	}
}
