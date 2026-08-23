package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketTimeTrialData struct {
	Header                   PacketHeader     // Header
	PlayerSessionBestDataSet TimeTrialDataSet // Player session best data set
	PersonalBestDataSet      TimeTrialDataSet // Personal best data set
	RivalDataSet             TimeTrialDataSet // Rival data set
}

// GetHeader returns the Header of *PacketTimeTrialData
func (data *PacketTimeTrialData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketTimeTrialData
func (data *PacketTimeTrialData) SetHeader(v PacketHeader) { data.Header = v }

// GetPlayerSessionBestDataSet returns the PlayerSessionBestDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) GetPlayerSessionBestDataSet() TimeTrialDataSet {
	return data.PlayerSessionBestDataSet
}

// SetPlayerSessionBestDataSet stores the PlayerSessionBestDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) SetPlayerSessionBestDataSet(v TimeTrialDataSet) {
	data.PlayerSessionBestDataSet = v
}

// GetPersonalBestDataSet returns the PersonalBestDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) GetPersonalBestDataSet() TimeTrialDataSet {
	return data.PersonalBestDataSet
}

// SetPersonalBestDataSet stores the PersonalBestDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) SetPersonalBestDataSet(v TimeTrialDataSet) {
	data.PersonalBestDataSet = v
}

// GetRivalDataSet returns the RivalDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) GetRivalDataSet() TimeTrialDataSet { return data.RivalDataSet }

// SetRivalDataSet stores the RivalDataSet of *PacketTimeTrialData
func (data *PacketTimeTrialData) SetRivalDataSet(v TimeTrialDataSet) { data.RivalDataSet = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketTimeTrialData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.PlayerSessionBestDataSet.Parse(reader)
	data.PersonalBestDataSet.Parse(reader)
	data.RivalDataSet.Parse(reader)

}
