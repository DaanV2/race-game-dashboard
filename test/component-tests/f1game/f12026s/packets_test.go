package f12026s_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/daanv2/race-game-dashboard/pkg/f1game/f12026s"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Laps(t *testing.T) {
	files := []string{
		"lap_records/timetrail-lap-suzuka.txt",
		"lap_records/gp-zandvoort.txt",
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			var checks = map[f12026s.PacketID]int{}

			pckids := f12026s.PacketIds()
			handler := &f12026s.PacketHandler{}

			// Record each package coming in, if it came in, its record 2 (or second bit)
			handler.Motion.Register(recordExist[*f12026s.PacketMotionData](checks))
			handler.Session.Register(recordExist[*f12026s.PacketSessionData](checks))
			handler.LapData.Register(recordExist[*f12026s.PacketLapData](checks))
			handler.Event.Register(recordExist[*f12026s.PacketEventData](checks))
			handler.Participants.Register(recordExist[*f12026s.PacketParticipantsData](checks))
			handler.CarSetups.Register(recordExist[*f12026s.PacketCarSetupData](checks))
			handler.CarTelemetry.Register(recordExist[*f12026s.PacketCarTelemetryData](checks))
			handler.CarTelemetry2.Register(recordExist[*f12026s.PacketCarTelemetry2Data](checks))
			handler.CarStatus.Register(recordExist[*f12026s.PacketCarStatusData](checks))
			handler.FinalClassification.Register(recordExist[*f12026s.PacketFinalClassificationData](checks))
			handler.LobbyInfo.Register(recordExist[*f12026s.PacketLobbyInfoData](checks))
			handler.CarDamage.Register(recordExist[*f12026s.PacketCarDamageData](checks))
			handler.SessionHistory.Register(recordExist[*f12026s.PacketSessionHistoryData](checks))
			handler.TyreSets.Register(recordExist[*f12026s.PacketTyreSetsData](checks))
			handler.MotionEx.Register(recordExist[*f12026s.PacketMotionExData](checks))
			handler.TimeTrial.Register(recordExist[*f12026s.PacketTimeTrialData](checks))
			handler.LapPositions.Register(recordExist[*f12026s.PacketLapPositionsData](checks))

			body, err := os.ReadFile(filepath.Clean(file))
			require.NoError(t, err)

			// Process file, each line is a base64 encoded packet
			// If packetId is valid, and parsing when well, check the header and record first bit if found
			var index int
			for line := range bytes.SplitSeq(body, []byte{'\n'}) {
				if len(line) == 0 {
					index++

					continue
				}

				data, err := base64.RawStdEncoding.AppendDecode(nil, line)
				require.NoError(t, err, "at index: %v", index)

				p := handler.ParsePacket(data)
				assert.NotNil(t, p, "at index: %v", index)
				assert.Contains(t, pckids, p.GetPacketID(), "at index: %v", index)

				checks[p.GetPacketID()] |= 1

				v, _ := json.Marshal(p)
				fmt.Println(string(v))

				handler.Ingest(data)
				validateHeader(t, p.GetHeader(), index)

				// NEXT
				index++
			}

			for pid, v := range checks {
				assert.Equal(t, 1|2, v, "no valid packets for: %s", pid)
			}
		})
	}

}

func validateHeader(t *testing.T, header f12026s.PacketHeader, line int) {
	assert.Equal(t, f12026s.PACKET_FORMAT, header.PacketFormat, "at index: %v", line)
	assert.EqualValues(t, 25, header.GameYear, "at index: %v", line)
	assert.EqualValues(t, 1, header.GameMajorVersion, "at index: %v", line)
	assert.Positive(t, header.GameMinorVersion, "at index: %v", line)
	assert.EqualValues(t, 1, header.PacketVersion, "at index: %v", line)

	assert.GreaterOrEqual(t, header.PacketId, f12026s.PacketID(0), "at index: %v", line)
	assert.LessOrEqual(t, header.PacketId, f12026s.PacketID(16), "at index: %v", line)

	// No need to test if above 0, its a uint 0 is the lowest

	assert.LessOrEqual(t, header.PlayerCarIndex, uint8(f12026s.CS_MAX_NUM_CARS), "at index: %v", line)

	if header.SecondaryPlayerCarIndex != 255 {
		assert.LessOrEqual(t, header.SecondaryPlayerCarIndex, uint8(f12026s.CS_MAX_NUM_CARS), "at index: %v", line)
	}

	// assert.Equal(t, 0, header.SessionTime, "at index: %v", line)
	// assert.Equal(t, 0, header.SessionUID, "at index: %v", line)
	// assert.Equal(t, 0, header.OverallFrameIdentifier, "at index: %v", line)
	// assert.Equal(t, 0, header.FrameIdentifier, "at index: %v", line)

}

func recordExist[T f12026s.Packet](receiver map[f12026s.PacketID]int) func(data T) {
	return func(data T) {
		receiver[data.GetPacketID()] |= 2
	}
}
