package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketSessionHistoryData -ignore-field LapHistoryData

type PacketSessionHistoryData struct {
	Header                PacketHeader            // Header
	CarIdx                uint8                   // Index of the car this lap data relates to
	NumLaps               uint8                   // Num laps in the data (including current partial lap)
	NumTyreStints         uint8                   // Number of tyre stints in the data
	BestLapTimeLapNum     uint8                   // Lap the best lap time was achieved on
	BestSector1LapNum     uint8                   // Lap the best Sector 1 time was achieved on
	BestSector2LapNum     uint8                   // Lap the best Sector 2 time was achieved on
	BestSector3LapNum     uint8                   // Lap the best Sector 3 time was achieved on
	LapHistoryData        [100]LapHistoryData     // 100 laps of data max
	TyreStintsHistoryData [8]TyreStintHistoryData //
}

// GetPacketID returns the identification of this packet
func (data *PacketSessionHistoryData) GetPacketID() PacketID { return PACKET_ID_SESSION_HISTORY }

// GetLapHistoryData returns the LapHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetLapHistoryData(lapIndex int) LapHistoryData {
	return data.LapHistoryData[lapIndex]
}

// SetLapHistoryData stores the LapHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetLapHistoryData(lapIndex int, v LapHistoryData) {
	data.LapHistoryData[lapIndex] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketSessionHistoryData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.CarIdx = reader.ReadUint8()
	data.NumLaps = reader.ReadUint8()
	data.NumTyreStints = reader.ReadUint8()
	data.BestLapTimeLapNum = reader.ReadUint8()
	data.BestSector1LapNum = reader.ReadUint8()
	data.BestSector2LapNum = reader.ReadUint8()
	data.BestSector3LapNum = reader.ReadUint8()

	for i := range data.LapHistoryData {
		data.LapHistoryData[i].Parse(reader)
	}

	for i := range data.TyreStintsHistoryData {
		data.TyreStintsHistoryData[i].Parse(reader)
	}
}
