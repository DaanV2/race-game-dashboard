package f12025 // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketLobbyInfoData -ignore-field LobbyPlayers

type PacketLobbyInfoData struct {
	Header       PacketHeader                   // Header  Packet specific data
	NumPlayers   uint8                          // Number of players in the lobby data
	LobbyPlayers [CS_MAX_NUM_CARS]LobbyInfoData //
}

func (data *PacketLobbyInfoData) GetPacketID() PacketID { return PACKET_ID_LOBBY_INFO }

// GetLobbyPlayers returns the LobbyPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) GetLobbyPlayers(participant int) LobbyInfoData {
	return data.LobbyPlayers[participant]
}

// SetLobbyPlayers stores the LobbyPlayers of *PacketLobbyInfoData
func (data *PacketLobbyInfoData) SetLobbyPlayers(participant int, v LobbyInfoData) {
	data.LobbyPlayers[participant] = v
}

func (data *PacketLobbyInfoData) GetPlayerData() LobbyInfoData {
	return data.LobbyPlayers[data.Header.GetPlayerCarIndex()]
}

func (data *PacketLobbyInfoData) GetSecondPlayerData() LobbyInfoData {
	return data.LobbyPlayers[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketLobbyInfoData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumPlayers = reader.ReadUint8()

	for i := range data.LobbyPlayers {
		data.LobbyPlayers[i].Parse(reader)
	}

}
