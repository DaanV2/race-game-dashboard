package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarDamageData struct {
	Header        PacketHeader      // Header
	CarDamageData [24]CarDamageData //
}

// GetHeader returns the Header of *PacketCarDamageData
func (data *PacketCarDamageData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarDamageData
func (data *PacketCarDamageData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarDamageData returns the CarDamageData of *PacketCarDamageData
func (data *PacketCarDamageData) GetCarDamageData() [24]CarDamageData { return data.CarDamageData }

// SetCarDamageData stores the CarDamageData of *PacketCarDamageData
func (data *PacketCarDamageData) SetCarDamageData(v [24]CarDamageData) { data.CarDamageData = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarDamageData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarDamageData {
		data.CarDamageData[i].Parse(reader)
	}
}
