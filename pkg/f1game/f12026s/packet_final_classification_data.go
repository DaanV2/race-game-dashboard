package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketFinalClassificationData struct {
	Header             PacketHeader                             // Header
	NumCars            uint8                                    // Number of cars in the final classification
	ClassificationData [CS_MAX_NUM_CARS]FinalClassificationData //
}

// GetHeader returns the Header of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) SetHeader(v PacketHeader) { data.Header = v }

// GetNumCars returns the NumCars of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) GetNumCars() uint8 { return data.NumCars }

// SetNumCars stores the NumCars of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) SetNumCars(v uint8) { data.NumCars = v }

// GetClassificationData returns the ClassificationData of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) GetClassificationData(car int) FinalClassificationData {
	return data.ClassificationData[car]
}

// SetClassificationData stores the ClassificationData of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) SetClassificationData(car int, v FinalClassificationData) {
	data.ClassificationData[car] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketFinalClassificationData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumCars = reader.ReadUint8()

	for i := range data.ClassificationData {
		data.ClassificationData[i].Parse(reader)
	}

}
