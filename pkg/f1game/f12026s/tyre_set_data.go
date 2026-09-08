package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type TyreSetData

type TyreSetData struct {
	ActualTyreCompound uint8 // Actual tyre compound used
	VisualTyreCompound uint8 // Visual tyre compound used
	Wear               uint8 // Tyre wear (percentage)
	Available          uint8 // Whether this set is currently available
	RecommendedSession uint8 // Recommended session for tyre set, see appendix
	LifeSpan           uint8 // Laps left in this tyre set
	UsableLife         uint8 // Max number of laps recommended for this compound
	LapDeltaTime       int16 // Lap delta time in milliseconds compared to fitted set
	Fitted             uint8 // Whether the set is fitted or not
}

func (data *TyreSetData) Parse(reader *xbinary.LittleEndianReader) {
	data.ActualTyreCompound = reader.ReadUint8()
	data.VisualTyreCompound = reader.ReadUint8()
	data.Wear = reader.ReadUint8()
	data.Available = reader.ReadUint8()
	data.RecommendedSession = reader.ReadUint8()
	data.LifeSpan = reader.ReadUint8()
	data.UsableLife = reader.ReadUint8()
	data.LapDeltaTime = reader.ReadInt16()
	data.Fitted = reader.ReadUint8()

}
