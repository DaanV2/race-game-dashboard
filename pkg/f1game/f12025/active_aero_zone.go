package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type ActiveAeroZone

type ActiveAeroZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the Active Aero zone starts
	ZoneEnd   float32 // Fraction (0..1) of way through the lap the Active Aero zone ends
}

func (data *ActiveAeroZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneEnd = reader.ReadFloat32()
}
