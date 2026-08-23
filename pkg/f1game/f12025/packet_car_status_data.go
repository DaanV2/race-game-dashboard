package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarStatusData struct {
	Header        PacketHeader      // Header
	CarStatusData [24]CarStatusData //
}

// GetHeader returns the Header of *PacketCarStatusData
func (data *PacketCarStatusData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarStatusData
func (data *PacketCarStatusData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarStatusData returns the CarStatusData of *PacketCarStatusData
func (data *PacketCarStatusData) GetCarStatusData() [24]CarStatusData { return data.CarStatusData }

// SetCarStatusData stores the CarStatusData of *PacketCarStatusData
func (data *PacketCarStatusData) SetCarStatusData(v [24]CarStatusData) { data.CarStatusData = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarStatusData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarStatusData {
		data.CarStatusData[i].Parse(reader)
	}
}
