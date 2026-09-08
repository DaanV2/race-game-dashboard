package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketCarTelemetryData -ignore-field CarTelemetryData

type PacketCarTelemetryData struct {
	Header                       PacketHeader                      // Header
	CarTelemetryData             [CS_MAX_NUM_CARS]CarTelemetryData //
	MfdPanelIndex                uint8                             // Index of MFD panel open - 255 = MFD closed  Single player, race – 0 = Car setup, 1 = Pits  2 = Damage, 3 =  Engine, 4 = Temperatures  May vary depending on game mode
	MfdPanelIndexSecondaryPlayer uint8                             // See above
	SuggestedGear                int8                              // Suggested gear for the player (1-8)  0 if no gear suggested
}

// GetPacketID returns the identification of this packet
func (data *PacketCarTelemetryData) GetPacketID() PacketID { return PACKET_ID_CAR_TELEMETRY }

// GetCarTelemetryData returns the CarTelemetryData of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetCarTelemetryData(car int) CarTelemetryData {
	return data.CarTelemetryData[car]
}

// SetCarTelemetryData stores the CarTelemetryData of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetCarTelemetryData(car int, v CarTelemetryData) {
	data.CarTelemetryData[car] = v
}

func (data *PacketCarTelemetryData) GetPlayerData() CarTelemetryData {
	return data.CarTelemetryData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketCarTelemetryData) GetSecondPlayerData() CarTelemetryData {
	return data.CarTelemetryData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarTelemetryData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarTelemetryData {
		data.CarTelemetryData[i].Parse(reader)
	}

	data.MfdPanelIndex = reader.ReadUint8()
	data.MfdPanelIndexSecondaryPlayer = reader.ReadUint8()
	data.SuggestedGear = reader.ReadInt8()
}
