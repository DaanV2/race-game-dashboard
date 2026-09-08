package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketFinalClassificationData -ignore-field ClassificationData

type PacketFinalClassificationData struct {
	Header             PacketHeader                             // Header
	NumCars            uint8                                    // Number of cars in the final classification
	ClassificationData [CS_MAX_NUM_CARS]FinalClassificationData //
}

// GetPacketID returns the identification of this packet
func (data *PacketFinalClassificationData) GetPacketID() PacketID {
	return PACKET_ID_FINAL_CLASSIFICATION
}

// GetClassificationData returns the ClassificationData of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) GetClassificationData(car int) FinalClassificationData {
	return data.ClassificationData[car]
}

// SetClassificationData stores the ClassificationData of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) SetClassificationData(car int, v FinalClassificationData) {
	data.ClassificationData[car] = v
}

func (data *PacketFinalClassificationData) GetPlayerData() FinalClassificationData {
	return data.ClassificationData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketFinalClassificationData) GetSecondPlayerData() FinalClassificationData {
	return data.ClassificationData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketFinalClassificationData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumCars = reader.ReadUint8()

	for i := range data.ClassificationData {
		data.ClassificationData[i].Parse(reader)
	}
}
