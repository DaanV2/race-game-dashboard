package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketLapData -ignore-field LapData

type PacketLapData struct {
	Header               PacketHeader             // Header
	LapData              [CS_MAX_NUM_CARS]LapData // Lap data for all cars on track
	TimeTrialPBCarIdx    uint8                    // Index of Personal Best car in time trial (255 if invalid)
	TimeTrialRivalCarIdx uint8                    // Index of Rival car in time trial (255 if invalid)
}

// GetPacketID returns the identification of this packet
func (data *PacketLapData) GetPacketID() PacketID { return PACKET_ID_LAP_DATA }

// GetLapData returns the LapData of *PacketLapData
func (data *PacketLapData) GetLapData(lap int) LapData { return data.LapData[lap] }

// SetLapData stores the LapData of *PacketLapData
func (data *PacketLapData) SetLapData(lap int, v LapData) { data.LapData[lap] = v }

func (data *PacketLapData) GetPlayerData() LapData {

	return data.LapData[data.Header.GetPlayerCarIndex()]
}

func (data *PacketLapData) GetSecondPlayerData() LapData {
	return data.LapData[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketLapData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header

	for i := range data.LapData {
		data.LapData[i].Parse(reader)
	}

	data.TimeTrialPBCarIdx = reader.ReadUint8()
	data.TimeTrialRivalCarIdx = reader.ReadUint8()
}
