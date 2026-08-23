package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetActualTyreCompound returns the ActualTyreCompound of *TyreSetData
func (data *TyreSetData) GetActualTyreCompound() uint8 { return data.ActualTyreCompound }

// SetActualTyreCompound stores the ActualTyreCompound of *TyreSetData
func (data *TyreSetData) SetActualTyreCompound(v uint8) { data.ActualTyreCompound = v }

// GetVisualTyreCompound returns the VisualTyreCompound of *TyreSetData
func (data *TyreSetData) GetVisualTyreCompound() uint8 { return data.VisualTyreCompound }

// SetVisualTyreCompound stores the VisualTyreCompound of *TyreSetData
func (data *TyreSetData) SetVisualTyreCompound(v uint8) { data.VisualTyreCompound = v }

// GetWear returns the Wear of *TyreSetData
func (data *TyreSetData) GetWear() uint8 { return data.Wear }

// SetWear stores the Wear of *TyreSetData
func (data *TyreSetData) SetWear(v uint8) { data.Wear = v }

// GetAvailable returns the Available of *TyreSetData
func (data *TyreSetData) GetAvailable() uint8 { return data.Available }

// SetAvailable stores the Available of *TyreSetData
func (data *TyreSetData) SetAvailable(v uint8) { data.Available = v }

// GetRecommendedSession returns the RecommendedSession of *TyreSetData
func (data *TyreSetData) GetRecommendedSession() uint8 { return data.RecommendedSession }

// SetRecommendedSession stores the RecommendedSession of *TyreSetData
func (data *TyreSetData) SetRecommendedSession(v uint8) { data.RecommendedSession = v }

// GetLifeSpan returns the LifeSpan of *TyreSetData
func (data *TyreSetData) GetLifeSpan() uint8 { return data.LifeSpan }

// SetLifeSpan stores the LifeSpan of *TyreSetData
func (data *TyreSetData) SetLifeSpan(v uint8) { data.LifeSpan = v }

// GetUsableLife returns the UsableLife of *TyreSetData
func (data *TyreSetData) GetUsableLife() uint8 { return data.UsableLife }

// SetUsableLife stores the UsableLife of *TyreSetData
func (data *TyreSetData) SetUsableLife(v uint8) { data.UsableLife = v }

// GetLapDeltaTime returns the LapDeltaTime of *TyreSetData
func (data *TyreSetData) GetLapDeltaTime() int16 { return data.LapDeltaTime }

// SetLapDeltaTime stores the LapDeltaTime of *TyreSetData
func (data *TyreSetData) SetLapDeltaTime(v int16) { data.LapDeltaTime = v }

// GetFitted returns the Fitted of *TyreSetData
func (data *TyreSetData) GetFitted() uint8 { return data.Fitted }

// SetFitted stores the Fitted of *TyreSetData
func (data *TyreSetData) SetFitted(v uint8) { data.Fitted = v }

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
