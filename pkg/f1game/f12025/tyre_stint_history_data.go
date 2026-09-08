package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type TyreStintHistoryData

type TyreStintHistoryData struct {
	EndLap             uint8 // Lap the tyre usage ends on (255 of current tyre)
	TyreActualCompound uint8 // Actual tyres used by this driver
	TyreVisualCompound uint8 // Visual tyres used by this driver
}

func (data *TyreStintHistoryData) Parse(reader *xbinary.LittleEndianReader) {
	data.EndLap = reader.ReadUint8()
	data.TyreActualCompound = reader.ReadUint8()
	data.TyreVisualCompound = reader.ReadUint8()

}
