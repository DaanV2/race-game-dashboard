package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarTelemetry2Data struct {
	Header            PacketHeader                       // Header
	CarTelemetry2Data [CS_MAX_NUM_CARS]CarTelemetry2Data //
}

// GetPacketID returns the identification of this packet
func (data *PacketCarTelemetry2Data) GetPacketID() PacketID { return PACKET_ID_CAR_TELEMETRY_2 }

// GetHeader returns the Header of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) SetHeader(v PacketHeader) { data.Header = v }

// GetCarTelemetry2Data returns the CarTelemetry2Data of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) GetCarTelemetry2Data(car int) CarTelemetry2Data {
	return data.CarTelemetry2Data[car]
}

// SetCarTelemetry2Data stores the CarTelemetry2Data of *PacketCarTelemetry2Data
func (data *PacketCarTelemetry2Data) SetCarTelemetry2Data(car int, v CarTelemetry2Data) {
	data.CarTelemetry2Data[car] = v
}

func (data *PacketCarTelemetry2Data) GetPlayerData() CarTelemetry2Data {
	carIndex := data.Header.GetPlayerCarIndex()

	return data.CarTelemetry2Data[carIndex]
}

func (data *PacketCarTelemetry2Data) GetSecondPlayerData() CarTelemetry2Data {
	carIndex := data.Header.GetSecondaryPlayerCarIndex()

	return data.CarTelemetry2Data[carIndex]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarTelemetry2Data) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarTelemetry2Data {
		data.CarTelemetry2Data[i].Parse(reader)
	}
}
