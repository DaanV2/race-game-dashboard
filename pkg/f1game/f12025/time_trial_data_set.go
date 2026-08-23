package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type TimeTrialDataSet struct {
	CarIdx              uint8  // Index of the car this data relates to
	TeamId              uint16 // Team id - see appendix
	LapTimeInMS         uint32 // Lap time in milliseconds
	Sector1TimeInMS     uint32 // Sector 1 time in milliseconds
	Sector2TimeInMS     uint32 // Sector 2 time in milliseconds
	Sector3TimeInMS     uint32 // Sector 3 time in milliseconds
	TractionControl     uint8  // 0 = assist off, 1 = assist on
	GearboxAssist       uint8  // 0 = assist off, 1 = assist on
	AntiLockBrakes      uint8  // 0 = assist off, 1 = assist on
	EqualCarPerformance uint8  // 0 = Realistic, 1 = Equal
	CustomSetup         uint8  // 0 = No, 1 = Yes
	Valid               uint8  // 0 = invalid, 1 = valid
}

// GetCarIdx returns the CarIdx of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetCarIdx() uint8 { return data.CarIdx }

// SetCarIdx stores the CarIdx of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetCarIdx(v uint8) { data.CarIdx = v }

// GetTeamId returns the TeamId of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetTeamId() uint16 { return data.TeamId }

// SetTeamId stores the TeamId of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetTeamId(v uint16) { data.TeamId = v }

// GetLapTimeInMS returns the LapTimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetLapTimeInMS() uint32 { return data.LapTimeInMS }

// SetLapTimeInMS stores the LapTimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetLapTimeInMS(v uint32) { data.LapTimeInMS = v }

// GetSector1TimeInMS returns the Sector1TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetSector1TimeInMS() uint32 { return data.Sector1TimeInMS }

// SetSector1TimeInMS stores the Sector1TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetSector1TimeInMS(v uint32) { data.Sector1TimeInMS = v }

// GetSector2TimeInMS returns the Sector2TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetSector2TimeInMS() uint32 { return data.Sector2TimeInMS }

// SetSector2TimeInMS stores the Sector2TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetSector2TimeInMS(v uint32) { data.Sector2TimeInMS = v }

// GetSector3TimeInMS returns the Sector3TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetSector3TimeInMS() uint32 { return data.Sector3TimeInMS }

// SetSector3TimeInMS stores the Sector3TimeInMS of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetSector3TimeInMS(v uint32) { data.Sector3TimeInMS = v }

// GetTractionControl returns the TractionControl of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetTractionControl() uint8 { return data.TractionControl }

// SetTractionControl stores the TractionControl of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetTractionControl(v uint8) { data.TractionControl = v }

// GetGearboxAssist returns the GearboxAssist of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetGearboxAssist() uint8 { return data.GearboxAssist }

// SetGearboxAssist stores the GearboxAssist of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetGearboxAssist(v uint8) { data.GearboxAssist = v }

// GetAntiLockBrakes returns the AntiLockBrakes of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetAntiLockBrakes() uint8 { return data.AntiLockBrakes }

// SetAntiLockBrakes stores the AntiLockBrakes of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetAntiLockBrakes(v uint8) { data.AntiLockBrakes = v }

// GetEqualCarPerformance returns the EqualCarPerformance of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetEqualCarPerformance() uint8 { return data.EqualCarPerformance }

// SetEqualCarPerformance stores the EqualCarPerformance of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetEqualCarPerformance(v uint8) { data.EqualCarPerformance = v }

// GetCustomSetup returns the CustomSetup of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetCustomSetup() uint8 { return data.CustomSetup }

// SetCustomSetup stores the CustomSetup of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetCustomSetup(v uint8) { data.CustomSetup = v }

// GetValid returns the Valid of *TimeTrialDataSet
func (data *TimeTrialDataSet) GetValid() uint8 { return data.Valid }

// SetValid stores the Valid of *TimeTrialDataSet
func (data *TimeTrialDataSet) SetValid(v uint8) { data.Valid = v }

func (data *TimeTrialDataSet) Parse(reader *xbinary.LittleEndianReader) {
	data.CarIdx = reader.ReadUint8()
	data.TeamId = reader.ReadUint16()
	data.LapTimeInMS = reader.ReadUint32()
	data.Sector1TimeInMS = reader.ReadUint32()
	data.Sector2TimeInMS = reader.ReadUint32()
	data.Sector3TimeInMS = reader.ReadUint32()
	data.TractionControl = reader.ReadUint8()
	data.GearboxAssist = reader.ReadUint8()
	data.AntiLockBrakes = reader.ReadUint8()
	data.EqualCarPerformance = reader.ReadUint8()
	data.CustomSetup = reader.ReadUint8()
	data.Valid = reader.ReadUint8()

}
