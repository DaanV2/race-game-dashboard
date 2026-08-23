package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetLapTimeInMS returns the LapTimeInMS of *LapHistoryData
func (data *LapHistoryData) GetLapTimeInMS() uint32 { return data.LapTimeInMS }

// SetLapTimeInMS stores the LapTimeInMS of *LapHistoryData
func (data *LapHistoryData) SetLapTimeInMS(v uint32) { data.LapTimeInMS = v }

// GetSector1TimeMSPart returns the Sector1TimeMSPart of *LapHistoryData
func (data *LapHistoryData) GetSector1TimeMSPart() uint16 { return data.Sector1TimeMSPart }

// SetSector1TimeMSPart stores the Sector1TimeMSPart of *LapHistoryData
func (data *LapHistoryData) SetSector1TimeMSPart(v uint16) { data.Sector1TimeMSPart = v }

// GetSector1TimeMinutesPart returns the Sector1TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) GetSector1TimeMinutesPart() uint8 { return data.Sector1TimeMinutesPart }

// SetSector1TimeMinutesPart stores the Sector1TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) SetSector1TimeMinutesPart(v uint8) { data.Sector1TimeMinutesPart = v }

// GetSector2TimeMSPart returns the Sector2TimeMSPart of *LapHistoryData
func (data *LapHistoryData) GetSector2TimeMSPart() uint16 { return data.Sector2TimeMSPart }

// SetSector2TimeMSPart stores the Sector2TimeMSPart of *LapHistoryData
func (data *LapHistoryData) SetSector2TimeMSPart(v uint16) { data.Sector2TimeMSPart = v }

// GetSector2TimeMinutesPart returns the Sector2TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) GetSector2TimeMinutesPart() uint8 { return data.Sector2TimeMinutesPart }

// SetSector2TimeMinutesPart stores the Sector2TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) SetSector2TimeMinutesPart(v uint8) { data.Sector2TimeMinutesPart = v }

// GetSector3TimeMSPart returns the Sector3TimeMSPart of *LapHistoryData
func (data *LapHistoryData) GetSector3TimeMSPart() uint16 { return data.Sector3TimeMSPart }

// SetSector3TimeMSPart stores the Sector3TimeMSPart of *LapHistoryData
func (data *LapHistoryData) SetSector3TimeMSPart(v uint16) { data.Sector3TimeMSPart = v }

// GetSector3TimeMinutesPart returns the Sector3TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) GetSector3TimeMinutesPart() uint8 { return data.Sector3TimeMinutesPart }

// SetSector3TimeMinutesPart stores the Sector3TimeMinutesPart of *LapHistoryData
func (data *LapHistoryData) SetSector3TimeMinutesPart(v uint8) { data.Sector3TimeMinutesPart = v }

// GetLapValidBitFlags returns the LapValidBitFlags of *LapHistoryData
func (data *LapHistoryData) GetLapValidBitFlags() uint8 { return data.LapValidBitFlags }

// SetLapValidBitFlags stores the LapValidBitFlags of *LapHistoryData
func (data *LapHistoryData) SetLapValidBitFlags(v uint8) { data.LapValidBitFlags = v }

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
