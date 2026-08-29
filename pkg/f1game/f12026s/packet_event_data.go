package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketEventData struct {
	Header          PacketHeader     // Header
	EventStringCode EventCode        // Event string code, see below
	EventDetails    EventDataDetails // Event details - should be interpreted differently  for each type
}

type EventDataDetails interface {
	Parse(reader *xbinary.LittleEndianReader)
}

// GetHeader returns the Header of *PacketEventData
func (data *PacketEventData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketEventData
func (data *PacketEventData) SetHeader(v PacketHeader) { data.Header = v }

// GetEventStringCode returns the EventStringCode of *PacketEventData
func (data *PacketEventData) GetEventStringCode() EventCode { return data.EventStringCode }

// SetEventStringCode stores the EventStringCode of *PacketEventData
func (data *PacketEventData) SetEventStringCode(v EventCode) { data.EventStringCode = v }

// GetEventDetails returns the EventDetails of *PacketEventData, returns any of the EvenData* structs
func (data *PacketEventData) GetEventDetails() EventDataDetails { return data.EventDetails }

// SetEventDetails stores the EventDetails of *PacketEventData
func (data *PacketEventData) SetEventDetails(v EventDataDetails) { data.EventDetails = v }

type parsable interface {
	Parse(reader *xbinary.LittleEndianReader)
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketEventData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) { // nolint:cyclop // Not needed here
	data.Header = *header

	var buf [CS_EVENT_STRING_CODE_LEN]byte
	reader.Read(buf[:])
	escStr := string(buf[:])
	data.EventStringCode = EventCode(escStr)

	var event parsable = nil

	switch data.EventStringCode {
	case EVENT_CODE_BUTTON_STATUS:
		event = &EventDataButtons{}
	case EVENT_CODE_CHEQUERED_FLAG:
		event = nil
	case EVENT_CODE_COLLISION:
		event = &EventDataCollision{}
	case EVENT_CODE_DRIVE_THROUGH_SERVED:
		event = &EventDataDriveThroughPenaltyServed{}
	case EVENT_CODE_DRS_DISABLED:
		event = &EventDataDRSDisabled{}
	case EVENT_CODE_DRS_ENABLED:
		event = nil
	case EVENT_CODE_FASTEST_LAP:
		event = &EventDataFastestLap{}
	case EVENT_CODE_FLASHBACK:
		event = &EventDataFlashback{}
	case EVENT_CODE_LIGHTS_OUT:
		event = nil
	case EVENT_CODE_OVERTAKE:
		event = &EventDataOvertake{}
	case EVENT_CODE_PENALTY_ISSUED:
		event = &EventDataPenalty{}
	case EVENT_CODE_RACE_WINNER:
		event = &EventDataRaceWinner{}
	case EVENT_CODE_RED_FLAG:
		event = nil
	case EVENT_CODE_RETIREMENT:
		event = &EventDataRetirement{}
	case EVENT_CODE_SAFETY_CAR:
		event = &EventDataSafetyCar{}
	case EVENT_CODE_SESSION_ENDED:
		event = nil
	case EVENT_CODE_SESSION_STARTED:
		event = nil
	case EVENT_CODE_SPEED_TRAP_TRIGGERED:
		event = &EventDataSpeedTrap{}
	case EVENT_CODE_START_LIGHTS:
		event = &EventDataStartLights{}
	case EVENT_CODE_STOP_GO_SERVED:
		event = &EventDataStopGoPenaltyServed{}
	case EVENT_CODE_TEAM_MATE_IN_PITS:
		event = &EventDataTeamMateInPits{}
	}

	if event != nil {
		data.EventDetails = event
		event.Parse(reader)
	}
}
