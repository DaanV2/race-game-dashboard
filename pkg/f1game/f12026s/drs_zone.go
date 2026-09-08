package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type DRSZone

type DRSZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the DRS zone starts
	ZoneEnd   float32 // Fraction (0..1) of way through the lap the DRS zone ends
}

func (data *DRSZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneEnd = reader.ReadFloat32()
}
