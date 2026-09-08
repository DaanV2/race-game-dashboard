package f12024 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type LapHistoryData

type LapHistoryData struct {
	LapTimeInMS            uint32 // Lap time in milliseconds
	Sector1TimeMSPart      uint16 // Sector 1 milliseconds part
	Sector1TimeMinutesPart uint8  // Sector 1 whole minute part
	Sector2TimeMSPart      uint16 // Sector 2 time milliseconds part
	Sector2TimeMinutesPart uint8  // Sector 2 whole minute part
	Sector3TimeMSPart      uint16 // Sector 3 time milliseconds part
	Sector3TimeMinutesPart uint8  // Sector 3 whole minute part
	LapValidBitFlags       uint8  // 0x01 bit set-lap valid,      0x02 bit set-sector 1 valid  0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
}

func (data *LapHistoryData) Parse(reader *xbinary.LittleEndianReader) {
	data.LapTimeInMS = reader.ReadUint32()
	data.Sector1TimeMSPart = reader.ReadUint16()
	data.Sector1TimeMinutesPart = reader.ReadUint8()
	data.Sector2TimeMSPart = reader.ReadUint16()
	data.Sector2TimeMinutesPart = reader.ReadUint8()
	data.Sector3TimeMSPart = reader.ReadUint16()
	data.Sector3TimeMinutesPart = reader.ReadUint8()
	data.LapValidBitFlags = reader.ReadUint8()
}
