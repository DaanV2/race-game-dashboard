package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type DRSZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the DRS zone starts
	ZoneEnd   float32 // Fraction (0..1) of way through the lap the DRS zone ends
}

// GetZoneStart returns the ZoneStart of *DRSZone
func (data *DRSZone) GetZoneStart() float32 { return data.ZoneStart }

// SetZoneStart stores the ZoneStart of *DRSZone
func (data *DRSZone) SetZoneStart(v float32) { data.ZoneStart = v }

// GetZoneEnd returns the ZoneEnd of *DRSZone
func (data *DRSZone) GetZoneEnd() float32 { return data.ZoneEnd }

// SetZoneEnd stores the ZoneEnd of *DRSZone
func (data *DRSZone) SetZoneEnd(v float32) { data.ZoneEnd = v }

func (data *DRSZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneEnd = reader.ReadFloat32()

}
