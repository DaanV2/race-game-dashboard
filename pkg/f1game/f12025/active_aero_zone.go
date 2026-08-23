package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type ActiveAeroZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the Active Aero zone starts
	ZoneEnd   float32 // Fraction (0..1) of way through the lap the Active Aero zone ends
}

// GetZoneStart returns the ZoneStart of *ActiveAeroZone
func (data *ActiveAeroZone) GetZoneStart() float32 { return data.ZoneStart }

// SetZoneStart stores the ZoneStart of *ActiveAeroZone
func (data *ActiveAeroZone) SetZoneStart(v float32) { data.ZoneStart = v }

// GetZoneEnd returns the ZoneEnd of *ActiveAeroZone
func (data *ActiveAeroZone) GetZoneEnd() float32 { return data.ZoneEnd }

// SetZoneEnd stores the ZoneEnd of *ActiveAeroZone
func (data *ActiveAeroZone) SetZoneEnd(v float32) { data.ZoneEnd = v }

func (data *ActiveAeroZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneEnd = reader.ReadFloat32()

}
