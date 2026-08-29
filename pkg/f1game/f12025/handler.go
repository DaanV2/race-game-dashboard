package f12025

import (
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

type PacketHandler struct {
	Motion              PacketPipeline[*PacketMotionData]
	Session             PacketPipeline[*PacketSessionData]
	LapData             PacketPipeline[*PacketLapData]
	Event               PacketPipeline[*PacketEventData]
	Participants        PacketPipeline[*PacketParticipantsData]
	CarSetups           PacketPipeline[*PacketCarSetupData]
	CarTelemetry        PacketPipeline[*PacketCarTelemetryData]
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
