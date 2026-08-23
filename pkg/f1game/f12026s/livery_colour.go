package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type LiveryColour struct {
	Red   uint8 //
	Green uint8 //
	Blue  uint8 //
}

// GetRed returns the Red of *LiveryColour
func (data *LiveryColour) GetRed() uint8 { return data.Red }

// SetRed stores the Red of *LiveryColour
func (data *LiveryColour) SetRed(v uint8) { data.Red = v }

// GetGreen returns the Green of *LiveryColour
func (data *LiveryColour) GetGreen() uint8 { return data.Green }

// SetGreen stores the Green of *LiveryColour
func (data *LiveryColour) SetGreen(v uint8) { data.Green = v }

// GetBlue returns the Blue of *LiveryColour
func (data *LiveryColour) GetBlue() uint8 { return data.Blue }

// SetBlue stores the Blue of *LiveryColour
func (data *LiveryColour) SetBlue(v uint8) { data.Blue = v }

func (data *LiveryColour) Parse(reader *xbinary.LittleEndianReader) {
	data.Red = reader.ReadUint8()
	data.Green = reader.ReadUint8()
	data.Blue = reader.ReadUint8()

}
