package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketLobbyInfoData struct {
	Header       PacketHeader                   // Header  Packet specific data
	NumPlayers   uint8                          // Number of players in the lobby data
	LobbyPlayers [CS_MAX_NUM_CARS]LobbyInfoData //
}

// GetHeader returns the Header of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) SetHeader(v PacketHeader) { data.Header = v }

// GetNumPlayers returns the NumPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) GetNumPlayers() uint8 { return data.NumPlayers }

// SetNumPlayers stores the NumPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) SetNumPlayers(v uint8) { data.NumPlayers = v }

// GetLobbyPlayers returns the LobbyPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) GetLobbyPlayers(participant int) LobbyInfoData {
	return data.LobbyPlayers[participant]
}

// SetLobbyPlayers stores the LobbyPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) SetLobbyPlayers(participant int, v LobbyInfoData) {
	data.LobbyPlayers[participant] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketLobbyInfoData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumPlayers = reader.ReadUint8()

	for i := range data.LobbyPlayers {
		data.LobbyPlayers[i].Parse(reader)
	}

}
