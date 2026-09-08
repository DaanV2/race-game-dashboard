package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketParticipantsData -ignore-field Participants

type PacketParticipantsData struct {
	Header        PacketHeader                     // Header
	NumActiveCars uint8                            // Number of active cars in the data – should match number of  cars on HUD
	Participants  [CS_MAX_NUM_CARS]ParticipantData //
}

// GetPacketID returns the identification of this packet
func (data *PacketParticipantsData) GetPacketID() PacketID { return PACKET_ID_PARTICIPANTS }

// GetParticipants returns the Participants of *PacketParticipantsData
func (data *PacketParticipantsData) GetParticipants(participant int) ParticipantData {
	return data.Participants[participant]
}

// SetParticipants stores the Participants of *PacketParticipantsData
func (data *PacketParticipantsData) SetParticipants(participant int, v ParticipantData) {
	data.Participants[participant] = v
}

func (data *PacketParticipantsData) GetPlayerData() ParticipantData {

	return data.Participants[data.Header.GetPlayerCarIndex()]
}

func (data *PacketParticipantsData) GetSecondPlayerData() ParticipantData {

	return data.Participants[data.Header.GetSecondaryPlayerCarIndex()]
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketParticipantsData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumActiveCars = reader.ReadUint8()

	for i := range data.Participants {
		data.Participants[i].Parse(reader)
	}
}
