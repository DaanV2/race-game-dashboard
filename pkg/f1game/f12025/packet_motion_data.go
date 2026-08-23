package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketMotionData struct {
	Header        PacketHeader                   // Header
	CarMotionData [CS_MAX_NUM_CARS]CarMotionData // Data for all cars on track
}

// GetHeader returns the Header of *PacketMotionData
func (data *PacketMotionData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketMotionData
func (data *PacketMotionData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarMotionData returns the CarMotionData of *PacketMotionData
func (data *PacketMotionData) GetCarMotionData(car int) CarMotionData { return data.CarMotionData[car] }

// SetCarMotionData stores the CarMotionData of *PacketMotionData
func (data *PacketMotionData) SetCarMotionData(car int, v CarMotionData) { data.CarMotionData[car] = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketMotionData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarMotionData {
		data.CarMotionData[i].Parse(reader)
	}
}
