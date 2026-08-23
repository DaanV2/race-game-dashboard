package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketFinalClassificationData struct {
	Header             PacketHeader                // Header
	NumCars            uint8                       // Number of cars in the final classification
	ClassificationData [24]FinalClassificationData //
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
func (data *PacketFinalClassificationData) GetClassificationData() [24]FinalClassificationData {
	return data.ClassificationData
}

// SetClassificationData stores the ClassificationData of *PacketFinalClassificationData
func (data *PacketFinalClassificationData) SetClassificationData(v [24]FinalClassificationData) {
	data.ClassificationData = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketFinalClassificationData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumCars = reader.ReadUint8()

	for i := range data.ClassificationData {
		data.ClassificationData[i].Parse(reader)
	}

}
