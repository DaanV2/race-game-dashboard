package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketTimeTrialData

type PacketTimeTrialData struct {
	Header                   PacketHeader     // Header
	PlayerSessionBestDataSet TimeTrialDataSet // Player session best data set
	PersonalBestDataSet      TimeTrialDataSet // Personal best data set
	RivalDataSet             TimeTrialDataSet // Rival data set
}

// GetPacketID returns the identification of this packet
func (data *PacketTimeTrialData) GetPacketID() PacketID { return PACKET_ID_TIME_TRIAL }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketTimeTrialData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.PlayerSessionBestDataSet.Parse(reader)
	data.PersonalBestDataSet.Parse(reader)
	data.RivalDataSet.Parse(reader)
}
