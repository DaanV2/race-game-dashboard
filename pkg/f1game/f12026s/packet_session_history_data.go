package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetHeader returns the Header of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetHeader(v PacketHeader) { data.Header = v }

// GetCarIdx returns the CarIdx of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetCarIdx() uint8 { return data.CarIdx }

// SetCarIdx stores the CarIdx of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetCarIdx(v uint8) { data.CarIdx = v }

// GetNumLaps returns the NumLaps of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetNumLaps() uint8 { return data.NumLaps }

// SetNumLaps stores the NumLaps of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetNumLaps(v uint8) { data.NumLaps = v }

// GetNumTyreStints returns the NumTyreStints of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetNumTyreStints() uint8 { return data.NumTyreStints }

// SetNumTyreStints stores the NumTyreStints of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetNumTyreStints(v uint8) { data.NumTyreStints = v }

// GetBestLapTimeLapNum returns the BestLapTimeLapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetBestLapTimeLapNum() uint8 { return data.BestLapTimeLapNum }

// SetBestLapTimeLapNum stores the BestLapTimeLapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetBestLapTimeLapNum(v uint8) { data.BestLapTimeLapNum = v }

// GetBestSector1LapNum returns the BestSector1LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetBestSector1LapNum() uint8 { return data.BestSector1LapNum }

// SetBestSector1LapNum stores the BestSector1LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetBestSector1LapNum(v uint8) { data.BestSector1LapNum = v }

// GetBestSector2LapNum returns the BestSector2LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetBestSector2LapNum() uint8 { return data.BestSector2LapNum }

// SetBestSector2LapNum stores the BestSector2LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetBestSector2LapNum(v uint8) { data.BestSector2LapNum = v }

// GetBestSector3LapNum returns the BestSector3LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetBestSector3LapNum() uint8 { return data.BestSector3LapNum }

// SetBestSector3LapNum stores the BestSector3LapNum of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetBestSector3LapNum(v uint8) { data.BestSector3LapNum = v }

// GetLapHistoryData returns the LapHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetLapHistoryData() [100]LapHistoryData {
	return data.LapHistoryData
}

// SetLapHistoryData stores the LapHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetLapHistoryData(v [100]LapHistoryData) {
	data.LapHistoryData = v
}

// GetTyreStintsHistoryData returns the TyreStintsHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) GetTyreStintsHistoryData() [8]TyreStintHistoryData {
	return data.TyreStintsHistoryData
}

// SetTyreStintsHistoryData stores the TyreStintsHistoryData of *PacketSessionHistoryData
func (data *PacketSessionHistoryData) SetTyreStintsHistoryData(v [8]TyreStintHistoryData) {
	data.TyreStintsHistoryData = v
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
