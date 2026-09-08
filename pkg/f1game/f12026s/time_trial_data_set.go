package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type TimeTrialDataSet

type TimeTrialDataSet struct {
	CarIdx              uint8           // Index of the car this data relates to
	TeamId              f1common.TeamID // Team id (uint16) - see appendix
	LapTimeInMS         uint32          // Lap time in milliseconds
	Sector1TimeInMS     uint32          // Sector 1 time in milliseconds
	Sector2TimeInMS     uint32          // Sector 2 time in milliseconds
	Sector3TimeInMS     uint32          // Sector 3 time in milliseconds
	TractionControl     uint8           // 0 = assist off, 1 = assist on
	GearboxAssist       uint8           // 0 = assist off, 1 = assist on
	AntiLockBrakes      uint8           // 0 = assist off, 1 = assist on
	EqualCarPerformance uint8           // 0 = Realistic, 1 = Equal
	CustomSetup         uint8           // 0 = No, 1 = Yes
	Valid               uint8           // 0 = invalid, 1 = valid
}

func (data *TimeTrialDataSet) Parse(reader *xbinary.LittleEndianReader) {
	data.CarIdx = reader.ReadUint8()
	data.TeamId = f1common.TeamID(reader.ReadUint16())
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
