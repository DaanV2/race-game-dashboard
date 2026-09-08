package f12026s

import (
	"encoding/base64"
	"fmt"

	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/generics"
)

type PacketPipeline[T Packet] struct {
	receiver []func(data T)
}

func (pp *PacketPipeline[T]) Register(call func(data T)) {
	if call == nil {
		return
	}

	pp.receiver = append(pp.receiver, call)
}

func (pp *PacketPipeline[T]) invoke(data T) {
	for _, call := range pp.receiver {
		call(data)
	}
}

func (pp *PacketPipeline[T]) handlePacket(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	if len(pp.receiver) == 0 {
		return
	}

	data := generics.New[T]()
	data.Parse(header, reader)

	pp.invoke(data)
}

func parsePacket[T Packet](header *PacketHeader, reader *xbinary.LittleEndianReader) T {
	data := generics.New[T]()
	data.Parse(header, reader)

	return data
}

// Example:
//
//	hand := &f12026s.PacketHandler{}
type PacketHandler struct {
	Motion              PacketPipeline[*PacketMotionData]
	Session             PacketPipeline[*PacketSessionData]
	LapData             PacketPipeline[*PacketLapData]
	Event               PacketPipeline[*PacketEventData]
	Participants        PacketPipeline[*PacketParticipantsData]
	CarSetups           PacketPipeline[*PacketCarSetupData]
	CarTelemetry        PacketPipeline[*PacketCarTelemetryData]
	CarTelemetry2       PacketPipeline[*PacketCarTelemetry2Data]
	CarStatus           PacketPipeline[*PacketCarStatusData]
	FinalClassification PacketPipeline[*PacketFinalClassificationData]
	LobbyInfo           PacketPipeline[*PacketLobbyInfoData]
	CarDamage           PacketPipeline[*PacketCarDamageData]
	SessionHistory      PacketPipeline[*PacketSessionHistoryData]
	TyreSets            PacketPipeline[*PacketTyreSetsData]
	MotionEx            PacketPipeline[*PacketMotionExData]
	TimeTrial           PacketPipeline[*PacketTimeTrialData]
	LapPositions        PacketPipeline[*PacketLapPositionsData]
}

func (h *PacketHandler) Ingest(data []byte) {
	defer func() {
		if r := recover(); r != nil {
			packet := base64.RawStdEncoding.EncodeToString(data)
			if err, ok := r.(error); ok {
				fmt.Printf("error with package: %s\nerror: %v\n", packet, err)
			} else {
				fmt.Printf("unknown error with package: %s\nerror: %v\n", packet, r)
			}
		}
	}()

	reader := NewByteReader(data)

	var header PacketHeader
	header.Parse(reader)

	switch header.PacketId {
	case PACKET_ID_MOTION:
		h.Motion.handlePacket(&header, reader)
	case PACKET_ID_SESSION:
		h.Session.handlePacket(&header, reader)
	case PACKET_ID_LAP_DATA:
		h.LapData.handlePacket(&header, reader)
	case PACKET_ID_EVENT:
		h.Event.handlePacket(&header, reader)
	case PACKET_ID_PARTICIPANTS:
		h.Participants.handlePacket(&header, reader)
	case PACKET_ID_CAR_SETUPS:
		h.CarSetups.handlePacket(&header, reader)
	case PACKET_ID_CAR_TELEMETRY:
		h.CarTelemetry.handlePacket(&header, reader)
	case PACKET_ID_CAR_TELEMETRY_2:
		h.CarTelemetry2.handlePacket(&header, reader)
	case PACKET_ID_CAR_STATUS:
		h.CarStatus.handlePacket(&header, reader)
	case PACKET_ID_FINAL_CLASSIFICATION:
		h.FinalClassification.handlePacket(&header, reader)
	case PACKET_ID_LOBBY_INFO:
		h.LobbyInfo.handlePacket(&header, reader)
	case PACKET_ID_CAR_DAMAGE:
		h.CarDamage.handlePacket(&header, reader)
	case PACKET_ID_SESSION_HISTORY:
		h.SessionHistory.handlePacket(&header, reader)
	case PACKET_ID_TYRE_SETS:
		h.TyreSets.handlePacket(&header, reader)
	case PACKET_ID_MOTION_EX:
		h.MotionEx.handlePacket(&header, reader)
	case PACKET_ID_TIME_TRIAL:
		h.TimeTrial.handlePacket(&header, reader)
	case PACKET_ID_LAP_POSITIONS:
		h.LapPositions.handlePacket(&header, reader)
	}
}

func (h *PacketHandler) ParsePacket(data []byte) Packet {
	defer func() {
		if r := recover(); r != nil {
			packet := base64.RawStdEncoding.EncodeToString(data)
			if err, ok := r.(error); ok {
				fmt.Printf("error with package: %s\nerror: %v\n", packet, err)
			} else {
				fmt.Printf("unknown error with package: %s\nerror: %v\n", packet, r)
			}
		}
	}()

	reader := NewByteReader(data)

	var header PacketHeader
	header.Parse(reader)

	switch header.PacketId {
	case PACKET_ID_MOTION:
		return parsePacket[*PacketMotionData](&header, reader)
	case PACKET_ID_SESSION:
		return parsePacket[*PacketSessionData](&header, reader)
	case PACKET_ID_LAP_DATA:
		return parsePacket[*PacketLapData](&header, reader)
	case PACKET_ID_EVENT:
		return parsePacket[*PacketEventData](&header, reader)
	case PACKET_ID_PARTICIPANTS:
		return parsePacket[*PacketParticipantsData](&header, reader)
	case PACKET_ID_CAR_SETUPS:
		return parsePacket[*PacketCarSetupData](&header, reader)
	case PACKET_ID_CAR_TELEMETRY:
		return parsePacket[*PacketCarTelemetryData](&header, reader)
	case PACKET_ID_CAR_TELEMETRY_2:
		return parsePacket[*PacketCarTelemetry2Data](&header, reader)
	case PACKET_ID_CAR_STATUS:
		return parsePacket[*PacketCarStatusData](&header, reader)
	case PACKET_ID_FINAL_CLASSIFICATION:
		return parsePacket[*PacketFinalClassificationData](&header, reader)
	case PACKET_ID_LOBBY_INFO:
		return parsePacket[*PacketLobbyInfoData](&header, reader)
	case PACKET_ID_CAR_DAMAGE:
		return parsePacket[*PacketCarDamageData](&header, reader)
	case PACKET_ID_SESSION_HISTORY:
		return parsePacket[*PacketSessionHistoryData](&header, reader)
	case PACKET_ID_TYRE_SETS:
		return parsePacket[*PacketTyreSetsData](&header, reader)
	case PACKET_ID_MOTION_EX:
		return parsePacket[*PacketMotionExData](&header, reader)
	case PACKET_ID_TIME_TRIAL:
		return parsePacket[*PacketTimeTrialData](&header, reader)
	case PACKET_ID_LAP_POSITIONS:
		return parsePacket[*PacketLapPositionsData](&header, reader)
	}

	return nil
}
