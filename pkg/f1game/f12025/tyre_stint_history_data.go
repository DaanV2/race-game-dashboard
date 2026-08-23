package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type TyreStintHistoryData struct {
	EndLap             uint8 // Lap the tyre usage ends on (255 of current tyre)
	TyreActualCompound uint8 // Actual tyres used by this driver
	TyreVisualCompound uint8 // Visual tyres used by this driver
}

// GetEndLap returns the EndLap of *TyreStintHistoryData
func (data *TyreStintHistoryData) GetEndLap() uint8 { return data.EndLap }

// SetEndLap stores the EndLap of *TyreStintHistoryData
func (data *TyreStintHistoryData) SetEndLap(v uint8) { data.EndLap = v }

// GetTyreActualCompound returns the TyreActualCompound of *TyreStintHistoryData
func (data *TyreStintHistoryData) GetTyreActualCompound() uint8 { return data.TyreActualCompound }

// SetTyreActualCompound stores the TyreActualCompound of *TyreStintHistoryData
func (data *TyreStintHistoryData) SetTyreActualCompound(v uint8) { data.TyreActualCompound = v }

// GetTyreVisualCompound returns the TyreVisualCompound of *TyreStintHistoryData
func (data *TyreStintHistoryData) GetTyreVisualCompound() uint8 { return data.TyreVisualCompound }

// SetTyreVisualCompound stores the TyreVisualCompound of *TyreStintHistoryData
func (data *TyreStintHistoryData) SetTyreVisualCompound(v uint8) { data.TyreVisualCompound = v }

func (data *TyreStintHistoryData) Parse(reader *xbinary.LittleEndianReader) {
	data.EndLap = reader.ReadUint8()
	data.TyreActualCompound = reader.ReadUint8()
	data.TyreVisualCompound = reader.ReadUint8()

}
