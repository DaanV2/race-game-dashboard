package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarTelemetry2Data struct {
	Header            PacketHeader          // Header
	CarTelemetry2Data [24]CarTelemetry2Data //
}

// GetHeader returns the Header of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) SetHeader(v PacketHeader) { data.Header = v }

// GetCarTelemetry2Data returns the CarTelemetry2Data of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) GetCarTelemetry2Data() [24]CarTelemetry2Data {
	return data.CarTelemetry2Data
}

// SetCarTelemetry2Data stores the CarTelemetry2Data of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) SetCarTelemetry2Data(v [24]CarTelemetry2Data) {
	data.CarTelemetry2Data = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarTelemetry2Data) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarTelemetry2Data {
		data.CarTelemetry2Data[i].Parse(reader)
	}
}
