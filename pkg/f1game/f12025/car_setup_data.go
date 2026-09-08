package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarSetupData

type CarSetupData struct {
	FrontWing              uint8   // Front wing aero
	RearWing               uint8   // Rear wing aero
	OnThrottle             uint8   // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   // Differential adjustment off throttle (percentage)
	FrontCamber            float32 // Front camber angle (suspension geometry)
	RearCamber             float32 // Rear camber angle (suspension geometry)
	FrontToe               float32 // Front toe angle (suspension geometry)
	RearToe                float32 // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   // Front suspension
	RearSuspension         uint8   // Rear suspension
	FrontAntiRollBar       uint8   // Front anti-roll bar
	RearAntiRollBar        uint8   // Front anti-roll bar
	FrontSuspensionHeight  uint8   // Front ride height
	RearSuspensionHeight   uint8   // Rear ride height
	BrakePressure          uint8   // Brake pressure (percentage)
	BrakeBias              uint8   // Brake bias (percentage)
	EngineBraking          uint8   // Engine braking (percentage)
	RearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 // Front right tyre pressure (PSI)
	Ballast                uint8   // Ballast
	FuelLoad               float32 // Fuel load
}

func (data *CarSetupData) TyrePressure() WheelMap[float32] {
	return NewWheelMap(data.RearLeftTyrePressure, data.RearRightTyrePressure, data.FrontLeftTyrePressure, data.FrontRightTyrePressure)
}

func (data *CarSetupData) Parse(reader *xbinary.LittleEndianReader) {
	data.FrontWing = reader.ReadUint8()
	data.RearWing = reader.ReadUint8()
	data.OnThrottle = reader.ReadUint8()
	data.OffThrottle = reader.ReadUint8()
	data.FrontCamber = reader.ReadFloat32()
	data.RearCamber = reader.ReadFloat32()
	data.FrontToe = reader.ReadFloat32()
	data.RearToe = reader.ReadFloat32()
	data.FrontSuspension = reader.ReadUint8()
	data.RearSuspension = reader.ReadUint8()
	data.FrontAntiRollBar = reader.ReadUint8()
	data.RearAntiRollBar = reader.ReadUint8()
	data.FrontSuspensionHeight = reader.ReadUint8()
	data.RearSuspensionHeight = reader.ReadUint8()
	data.BrakePressure = reader.ReadUint8()
	data.BrakeBias = reader.ReadUint8()
	data.EngineBraking = reader.ReadUint8()
	data.RearLeftTyrePressure = reader.ReadFloat32()
	data.RearRightTyrePressure = reader.ReadFloat32()
	data.FrontLeftTyrePressure = reader.ReadFloat32()
	data.FrontRightTyrePressure = reader.ReadFloat32()
	data.Ballast = reader.ReadUint8()
	data.FuelLoad = reader.ReadFloat32()

}
