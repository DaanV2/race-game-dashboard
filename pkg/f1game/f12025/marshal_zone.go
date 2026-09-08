package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type MarshalZone

type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
}

func (data *MarshalZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneFlag = reader.ReadInt8()
}
