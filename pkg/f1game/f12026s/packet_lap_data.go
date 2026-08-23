package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketLapData struct {
	Header               PacketHeader // Header
	LapData              [CS_MAX_NUM_CARS]LapData  // Lap data for all cars on track
	TimeTrialPBCarIdx    uint8        // Index of Personal Best car in time trial (255 if invalid)
	TimeTrialRivalCarIdx uint8        // Index of Rival car in time trial (255 if invalid)
}

// GetHeader returns the Header of *PacketLapData
func (data *PacketLapData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketLapData
func (data *PacketLapData) SetHeader(v PacketHeader) { data.Header = v }

// GetLapData returns the LapData of *PacketLapData
func (data *PacketLapData) GetLapData(lap int) LapData { return data.LapData[lap] }

// SetLapData stores the LapData of *PacketLapData
func (data *PacketLapData) SetLapData(lap int, v LapData) { data.LapData[lap] = v }

// GetTimeTrialPBCarIdx returns the TimeTrialPBCarIdx of *PacketLapData
func (data *PacketLapData) GetTimeTrialPBCarIdx() uint8 { return data.TimeTrialPBCarIdx }

// SetTimeTrialPBCarIdx stores the TimeTrialPBCarIdx of *PacketLapData
func (data *PacketLapData) SetTimeTrialPBCarIdx(v uint8) { data.TimeTrialPBCarIdx = v }

// GetTimeTrialRivalCarIdx returns the TimeTrialRivalCarIdx of *PacketLapData
func (data *PacketLapData) GetTimeTrialRivalCarIdx() uint8 { return data.TimeTrialRivalCarIdx }

// SetTimeTrialRivalCarIdx stores the TimeTrialRivalCarIdx of *PacketLapData
func (data *PacketLapData) SetTimeTrialRivalCarIdx(v uint8) { data.TimeTrialRivalCarIdx = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketLapData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.LapData {
		data.LapData[i].Parse(reader)
	}

	data.TimeTrialPBCarIdx = reader.ReadUint8()
	data.TimeTrialRivalCarIdx = reader.ReadUint8()

}
