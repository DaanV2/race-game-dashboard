package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketCarSetupData -ignore-field CarSetups

type PacketCarSetupData struct {
	Header             PacketHeader                  // Header
	CarSetups          [CS_MAX_NUM_CARS]CarSetupData //
	NextFrontWingValue float32                       // Value of front wing after next pit stop - player only
}

// GetPacketID returns the identification of this packet
func (data *PacketCarSetupData) GetPacketID() PacketID { return PACKET_ID_CAR_SETUPS }

// GetCarSetups returns the CarSetups of *PacketCarSetupData
func (data *PacketCarSetupData) GetCarSetups(car int) CarSetupData { return data.CarSetups[car] }

// SetCarSetups stores the CarSetups of *PacketCarSetupData
func (data *PacketCarSetupData) SetCarSetups(car int, v CarSetupData) { data.CarSetups[car] = v }

func (data *PacketCarSetupData) GetPlayerData() CarSetupData {
	return data.CarSetups[data.Header.GetPlayerCarIndex()]
}

func (data *PacketCarSetupData) GetSecondPlayerData() CarSetupData {
	return data.CarSetups[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarSetupData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarSetups {
		data.CarSetups[i].Parse(reader)
	}

	data.NextFrontWingValue = reader.ReadFloat32()
}
