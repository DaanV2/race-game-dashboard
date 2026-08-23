package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarTelemetryData struct {
	Header                       PacketHeader                      // Header
	CarTelemetryData             [CS_MAX_NUM_CARS]CarTelemetryData //
	MfdPanelIndex                uint8                             // Index of MFD panel open - 255 = MFD closed  Single player, race – 0 = Car setup, 1 = Pits  2 = Damage, 3 =  Engine, 4 = Temperatures  May vary depending on game mode
	MfdPanelIndexSecondaryPlayer uint8                             // See above
	SuggestedGear                int8                              // Suggested gear for the player (1-8)  0 if no gear suggested
}

// GetHeader returns the Header of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarTelemetryData returns the CarTelemetryData of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetCarTelemetryData() [CS_MAX_NUM_CARS]CarTelemetryData {
	return data.CarTelemetryData
}

// SetCarTelemetryData stores the CarTelemetryData of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetCarTelemetryData(v [CS_MAX_NUM_CARS]CarTelemetryData) {
	data.CarTelemetryData = v
}

// GetMfdPanelIndex returns the MfdPanelIndex of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetMfdPanelIndex() uint8 { return data.MfdPanelIndex }

// SetMfdPanelIndex stores the MfdPanelIndex of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetMfdPanelIndex(v uint8) { data.MfdPanelIndex = v }

// GetMfdPanelIndexSecondaryPlayer returns the MfdPanelIndexSecondaryPlayer of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetMfdPanelIndexSecondaryPlayer() uint8 {
	return data.MfdPanelIndexSecondaryPlayer
}

// SetMfdPanelIndexSecondaryPlayer stores the MfdPanelIndexSecondaryPlayer of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetMfdPanelIndexSecondaryPlayer(v uint8) {
	data.MfdPanelIndexSecondaryPlayer = v
}

// GetSuggestedGear returns the SuggestedGear of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) GetSuggestedGear() int8 { return data.SuggestedGear }

// SetSuggestedGear stores the SuggestedGear of *PacketCarTelemetryData
func (data *PacketCarTelemetryData) SetSuggestedGear(v int8) { data.SuggestedGear = v }

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
