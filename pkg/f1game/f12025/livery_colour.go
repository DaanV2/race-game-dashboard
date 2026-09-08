package f12025 // nolint:dupl // Don't care about dupl here

import (
	"image/color"

	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type LiveryColour

type LiveryColour struct {
	Red   uint8 //
	Green uint8 //
	Blue  uint8 //
}

func (data *LiveryColour) ToRGB() color.RGBA {
	return color.RGBA{
		R: data.Red,
		G: data.Green,
		B: data.Blue,
		A: 255,
	}
}

func (data *LiveryColour) Parse(reader *xbinary.LittleEndianReader) {
	data.Red = reader.ReadUint8()
	data.Green = reader.ReadUint8()
	data.Blue = reader.ReadUint8()

}
