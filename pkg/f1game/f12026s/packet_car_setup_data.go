package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketCarSetupData struct {
	Header             PacketHeader     // Header
	CarSetups          [CS_MAX_NUM_CARS]CarSetupData //
	NextFrontWingValue float32          // Value of front wing after next pit stop - player only
}

// GetHeader returns the Header of *PacketCarSetupData
func (data *PacketCarSetupData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketCarSetupData
func (data *PacketCarSetupData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarSetups returns the CarSetups of *PacketCarSetupData
func (data *PacketCarSetupData) GetCarSetups() [CS_MAX_NUM_CARS]CarSetupData { return data.CarSetups }

// SetCarSetups stores the CarSetups of *PacketCarSetupData
func (data *PacketCarSetupData) SetCarSetups(v [CS_MAX_NUM_CARS]CarSetupData) { data.CarSetups = v }

// GetNextFrontWingValue returns the NextFrontWingValue of *PacketCarSetupData
func (data *PacketCarSetupData) GetNextFrontWingValue() float32 { return data.NextFrontWingValue }

// SetNextFrontWingValue stores the NextFrontWingValue of *PacketCarSetupData
func (data *PacketCarSetupData) SetNextFrontWingValue(v float32) { data.NextFrontWingValue = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketCarSetupData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.CarSetups {
		data.CarSetups[i].Parse(reader)
	}

	data.NextFrontWingValue = reader.ReadFloat32()
}
