package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataStartLights struct {
	NumLights uint8 // Number of lights showing
}

// GetNumLights returns the NumLights of *StartLights
func (data *EventDataStartLights) GetNumLights() uint8 { return data.NumLights }

// SetNumLights stores the NumLights of *StartLights
func (data *EventDataStartLights) SetNumLights(v uint8) { data.NumLights = v }

func (data *EventDataStartLights) Parse(reader *xbinary.LittleEndianReader) {
	data.NumLights = reader.ReadUint8()

}
