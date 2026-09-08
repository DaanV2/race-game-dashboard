package f12026s // nolint:dupl // Don't care about dupl here

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarTelemetry2Data

type CarTelemetry2Data struct {
	ActiveAeroMode               uint8  // 0 = Corner mode, 1 = Straight mode
	ActiveAeroAvailable          uint8  // 0 = not available, 1 = available
	ActiveAeroActivationDistance uint16 // 0 = Active aero not available, non-zero – Active  aero will be available in [X] metres
	OvertakeAvailable            uint8  // 0 = not available, 1 = available
	OvertakeActive               uint8  // 0 = not active, 1 = active
	OvertakeActivationDistance   uint16 // 0 = Overtake Mode not available, non-zero –
	Regulations2026              uint8  // 0 = vehicle conforms to pre-2026, 1 = 2026  regulations applicable
	DrivingWrongWay              uint8  // Whether the car is driving the wrong way
}

func (data *CarTelemetry2Data) Parse(reader *xbinary.LittleEndianReader) {
	data.ActiveAeroMode = reader.ReadUint8()
	data.ActiveAeroAvailable = reader.ReadUint8()
	data.ActiveAeroActivationDistance = reader.ReadUint16()
	data.OvertakeAvailable = reader.ReadUint8()
	data.OvertakeActive = reader.ReadUint8()
	data.OvertakeActivationDistance = reader.ReadUint16()
	data.Regulations2026 = reader.ReadUint8()
	data.DrivingWrongWay = reader.ReadUint8()

}
