package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
}

// GetZoneStart returns the ZoneStart of *MarshalZone
func (data *MarshalZone) GetZoneStart() float32 { return data.ZoneStart }

// SetZoneStart stores the ZoneStart of *MarshalZone
func (data *MarshalZone) SetZoneStart(v float32) { data.ZoneStart = v }

// GetZoneFlag returns the ZoneFlag of *MarshalZone
func (data *MarshalZone) GetZoneFlag() int8 { return data.ZoneFlag }

// SetZoneFlag stores the ZoneFlag of *MarshalZone
func (data *MarshalZone) SetZoneFlag(v int8) { data.ZoneFlag = v }

func (data *MarshalZone) Parse(reader *xbinary.LittleEndianReader) {
	data.ZoneStart = reader.ReadFloat32()
	data.ZoneFlag = reader.ReadInt8()

}
