package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetActiveAeroMode returns the ActiveAeroMode of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetActiveAeroMode() uint8 { return data.ActiveAeroMode }

// SetActiveAeroMode stores the ActiveAeroMode of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetActiveAeroMode(v uint8) { data.ActiveAeroMode = v }

// GetActiveAeroAvailable returns the ActiveAeroAvailable of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetActiveAeroAvailable() uint8 { return data.ActiveAeroAvailable }

// SetActiveAeroAvailable stores the ActiveAeroAvailable of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetActiveAeroAvailable(v uint8) { data.ActiveAeroAvailable = v }

// GetActiveAeroActivationDistance returns the ActiveAeroActivationDistance of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetActiveAeroActivationDistance() uint16 {
	return data.ActiveAeroActivationDistance
}

// SetActiveAeroActivationDistance stores the ActiveAeroActivationDistance of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetActiveAeroActivationDistance(v uint16) {
	data.ActiveAeroActivationDistance = v
}

// GetOvertakeAvailable returns the OvertakeAvailable of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetOvertakeAvailable() uint8 { return data.OvertakeAvailable }

// SetOvertakeAvailable stores the OvertakeAvailable of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetOvertakeAvailable(v uint8) { data.OvertakeAvailable = v }

// GetOvertakeActive returns the OvertakeActive of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetOvertakeActive() uint8 { return data.OvertakeActive }

// SetOvertakeActive stores the OvertakeActive of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetOvertakeActive(v uint8) { data.OvertakeActive = v }

// GetOvertakeActivationDistance returns the OvertakeActivationDistance of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetOvertakeActivationDistance() uint16 {
	return data.OvertakeActivationDistance
}

// SetOvertakeActivationDistance stores the OvertakeActivationDistance of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetOvertakeActivationDistance(v uint16) {
	data.OvertakeActivationDistance = v
}

// Get2026Regulations returns the 2026Regulations of *CarTelemetry2Data
func (data *CarTelemetry2Data) Get2026Regulations() uint8 { return data.Regulations2026 }

// Set2026Regulations stores the 2026Regulations of *CarTelemetry2Data
func (data *CarTelemetry2Data) Set2026Regulations(v uint8) { data.Regulations2026 = v }

// GetDrivingWrongWay returns the DrivingWrongWay of *CarTelemetry2Data
func (data *CarTelemetry2Data) GetDrivingWrongWay() uint8 { return data.DrivingWrongWay }

// SetDrivingWrongWay stores the DrivingWrongWay of *CarTelemetry2Data
func (data *CarTelemetry2Data) SetDrivingWrongWay(v uint8) { data.DrivingWrongWay = v }

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
