package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketParticipantsData struct {
	Header        PacketHeader        // Header
	NumActiveCars uint8               // Number of active cars in the data – should match number of  cars on HUD
	Participants  [24]ParticipantData //
}

// GetHeader returns the Header of *PacketParticipantsData
func (data *PacketParticipantsData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketParticipantsData
func (data *PacketParticipantsData) SetHeader(v PacketHeader) { data.Header = v }

// GetNumActiveCars returns the NumActiveCars of *PacketParticipantsData
func (data *PacketParticipantsData) GetNumActiveCars() uint8 { return data.NumActiveCars }

// SetNumActiveCars stores the NumActiveCars of *PacketParticipantsData
func (data *PacketParticipantsData) SetNumActiveCars(v uint8) { data.NumActiveCars = v }

// GetParticipants returns the Participants of *PacketParticipantsData
func (data *PacketParticipantsData) GetParticipants(participant int) ParticipantData {
	return data.Participants[participant]
}

// SetParticipants stores the Participants of *PacketParticipantsData
func (data *PacketParticipantsData) SetParticipants(participant int, v ParticipantData) {
	data.Participants[participant] = v
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketParticipantsData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.NumActiveCars = reader.ReadUint8()

	for i := range data.Participants {
		data.Participants[i].Parse(reader)
	}
}
