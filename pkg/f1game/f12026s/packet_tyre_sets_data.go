package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketTyreSetsData -ignore-field TyreSetData

type PacketTyreSetsData struct {
	Header      PacketHeader                      // Header
	CarIdx      uint8                             // Index of the car this data relates to
	TyreSetData [CS_MAX_NUM_TYRE_SETS]TyreSetData // 13 (dry) + 7 (wet)
	FittedIdx   uint8                             // Index into array of fitted tyre
}

// GetPacketID returns the identification of this packet
func (data *PacketTyreSetsData) GetPacketID() PacketID { return PACKET_ID_TYRE_SETS }

// GetTyreSetData returns the TyreSetData of *PacketTyreSetsData
func (data *PacketTyreSetsData) GetTyreSetData(tyreSet int) TyreSetData {
	return data.TyreSetData[tyreSet]
}

// SetTyreSetData stores the TyreSetData of *PacketTyreSetsData
func (data *PacketTyreSetsData) SetTyreSetData(tyreSet int, v TyreSetData) {
	data.TyreSetData[tyreSet] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketTyreSetsData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.CarIdx = reader.ReadUint8()

	for i := range data.TyreSetData {
		data.TyreSetData[i].Parse(reader)
	}

	data.FittedIdx = reader.ReadUint8()
}
