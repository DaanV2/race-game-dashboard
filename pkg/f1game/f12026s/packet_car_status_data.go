package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarStatusData struct {
	Header        PacketHeader                   // Header
	CarStatusData [CS_MAX_NUM_CARS]CarStatusData //
}

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

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarStatusData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarStatusData {
		data.CarStatusData[i].Parse(reader)
	}
}
