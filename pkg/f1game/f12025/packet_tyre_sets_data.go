package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketTyreSetsData struct {
	Header      PacketHeader                      // Header
	CarIdx      uint8                             // Index of the car this data relates to
	TyreSetData [CS_MAX_NUM_TYRE_SETS]TyreSetData // 13 (dry) + 7 (wet)
	FittedIdx   uint8                             // Index into array of fitted tyre
}

// GetHeader returns the Header of *PacketTyreSetsData
func (data *PacketTyreSetsData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketTyreSetsData
func (data *PacketTyreSetsData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarIdx returns the CarIdx of *PacketTyreSetsData
func (data *PacketTyreSetsData) GetCarIdx() uint8 { return data.CarIdx }

// SetCarIdx stores the CarIdx of *PacketTyreSetsData
func (data *PacketTyreSetsData) SetCarIdx(v uint8) { data.CarIdx = v }

// GetTyreSetData returns the TyreSetData of *PacketTyreSetsData
func (data *PacketTyreSetsData) GetTyreSetData() [CS_MAX_NUM_TYRE_SETS]TyreSetData {
	return data.TyreSetData
}

// SetTyreSetData stores the TyreSetData of *PacketTyreSetsData
func (data *PacketTyreSetsData) SetTyreSetData(v [CS_MAX_NUM_TYRE_SETS]TyreSetData) {
	data.TyreSetData = v
}

// GetFittedIdx returns the FittedIdx of *PacketTyreSetsData
func (data *PacketTyreSetsData) GetFittedIdx() uint8 { return data.FittedIdx }

// SetFittedIdx stores the FittedIdx of *PacketTyreSetsData
func (data *PacketTyreSetsData) SetFittedIdx(v uint8) { data.FittedIdx = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketTyreSetsData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.CarIdx = reader.ReadUint8()

	for i := range data.TyreSetData {
		data.TyreSetData[i].Parse(reader)
	}

	data.FittedIdx = reader.ReadUint8()
}
