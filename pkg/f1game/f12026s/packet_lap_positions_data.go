package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketLapPositionsData struct {
	Header                PacketHeader                                                            // Header  Packet specific data
	NumLaps               uint8                                                                   // Number of laps in the data
	LapStart              uint8                                                                   // Index of the lap where the data starts, 0 indexed  Array holding the position of the car in a given lap, 0 if no record
	PositionForVehicleIdx [CS_MAX_NUM_LAPS_IN_LAP_POSITIONS_HISTORY_PACKET][CS_MAX_NUM_CARS]uint8 //
}

// GetHeader returns the Header of *PacketLapPositionsData
func (data *PacketLapPositionsData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketLapPositionsData
func (data *PacketLapPositionsData) SetHeader(v PacketHeader) { data.Header = v }

// GetNumLaps returns the NumLaps of *PacketLapPositionsData
func (data *PacketLapPositionsData) GetNumLaps() uint8 { return data.NumLaps }

// SetNumLaps stores the NumLaps of *PacketLapPositionsData
func (data *PacketLapPositionsData) SetNumLaps(v uint8) { data.NumLaps = v }

// GetLapStart returns the LapStart of *PacketLapPositionsData
func (data *PacketLapPositionsData) GetLapStart() uint8 { return data.LapStart }

// SetLapStart stores the LapStart of *PacketLapPositionsData
func (data *PacketLapPositionsData) SetLapStart(v uint8) { data.LapStart = v }

// GetPositionForVehicleIdx returns the PositionForVehicleIdx of *PacketLapPositionsData
func (data *PacketLapPositionsData) GetPositionForVehicleIdx(lap int) [CS_MAX_NUM_CARS]uint8 {
	return data.PositionForVehicleIdx[lap]
}

// SetPositionForVehicleIdx stores the PositionForVehicleIdx of *PacketLapPositionsData
func (data *PacketLapPositionsData) SetPositionForVehicleIdx(lap int, v [CS_MAX_NUM_CARS]uint8) {
	data.PositionForVehicleIdx[lap] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketLapPositionsData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumLaps = reader.ReadUint8()
	data.LapStart = reader.ReadUint8()

	for i := range data.PositionForVehicleIdx {
		var buf [CS_MAX_NUM_CARS]byte
		reader.Read(buf[:])

		data.PositionForVehicleIdx[i] = buf
	}
}
