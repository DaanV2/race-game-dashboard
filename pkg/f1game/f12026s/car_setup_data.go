package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetFrontWing returns the FrontWing of *CarSetupData
func (data *CarSetupData) GetFrontWing() uint8 { return data.FrontWing }

// SetFrontWing stores the FrontWing of *CarSetupData
func (data *CarSetupData) SetFrontWing(v uint8) { data.FrontWing = v }

// GetRearWing returns the RearWing of *CarSetupData
func (data *CarSetupData) GetRearWing() uint8 { return data.RearWing }

// SetRearWing stores the RearWing of *CarSetupData
func (data *CarSetupData) SetRearWing(v uint8) { data.RearWing = v }

// GetOnThrottle returns the OnThrottle of *CarSetupData
func (data *CarSetupData) GetOnThrottle() uint8 { return data.OnThrottle }

// SetOnThrottle stores the OnThrottle of *CarSetupData
func (data *CarSetupData) SetOnThrottle(v uint8) { data.OnThrottle = v }

// GetOffThrottle returns the OffThrottle of *CarSetupData
func (data *CarSetupData) GetOffThrottle() uint8 { return data.OffThrottle }

// SetOffThrottle stores the OffThrottle of *CarSetupData
func (data *CarSetupData) SetOffThrottle(v uint8) { data.OffThrottle = v }

// GetFrontCamber returns the FrontCamber of *CarSetupData
func (data *CarSetupData) GetFrontCamber() float32 { return data.FrontCamber }

// SetFrontCamber stores the FrontCamber of *CarSetupData
func (data *CarSetupData) SetFrontCamber(v float32) { data.FrontCamber = v }

// GetRearCamber returns the RearCamber of *CarSetupData
func (data *CarSetupData) GetRearCamber() float32 { return data.RearCamber }

// SetRearCamber stores the RearCamber of *CarSetupData
func (data *CarSetupData) SetRearCamber(v float32) { data.RearCamber = v }

// GetFrontToe returns the FrontToe of *CarSetupData
func (data *CarSetupData) GetFrontToe() float32 { return data.FrontToe }

// SetFrontToe stores the FrontToe of *CarSetupData
func (data *CarSetupData) SetFrontToe(v float32) { data.FrontToe = v }

// GetRearToe returns the RearToe of *CarSetupData
func (data *CarSetupData) GetRearToe() float32 { return data.RearToe }

// SetRearToe stores the RearToe of *CarSetupData
func (data *CarSetupData) SetRearToe(v float32) { data.RearToe = v }

// GetFrontSuspension returns the FrontSuspension of *CarSetupData
func (data *CarSetupData) GetFrontSuspension() uint8 { return data.FrontSuspension }

// SetFrontSuspension stores the FrontSuspension of *CarSetupData
func (data *CarSetupData) SetFrontSuspension(v uint8) { data.FrontSuspension = v }

// GetRearSuspension returns the RearSuspension of *CarSetupData
func (data *CarSetupData) GetRearSuspension() uint8 { return data.RearSuspension }

// SetRearSuspension stores the RearSuspension of *CarSetupData
func (data *CarSetupData) SetRearSuspension(v uint8) { data.RearSuspension = v }

// GetFrontAntiRollBar returns the FrontAntiRollBar of *CarSetupData
func (data *CarSetupData) GetFrontAntiRollBar() uint8 { return data.FrontAntiRollBar }

// SetFrontAntiRollBar stores the FrontAntiRollBar of *CarSetupData
func (data *CarSetupData) SetFrontAntiRollBar(v uint8) { data.FrontAntiRollBar = v }

// GetRearAntiRollBar returns the RearAntiRollBar of *CarSetupData
func (data *CarSetupData) GetRearAntiRollBar() uint8 { return data.RearAntiRollBar }

// SetRearAntiRollBar stores the RearAntiRollBar of *CarSetupData
func (data *CarSetupData) SetRearAntiRollBar(v uint8) { data.RearAntiRollBar = v }

// GetFrontSuspensionHeight returns the FrontSuspensionHeight of *CarSetupData
func (data *CarSetupData) GetFrontSuspensionHeight() uint8 { return data.FrontSuspensionHeight }

// SetFrontSuspensionHeight stores the FrontSuspensionHeight of *CarSetupData
func (data *CarSetupData) SetFrontSuspensionHeight(v uint8) { data.FrontSuspensionHeight = v }

// GetRearSuspensionHeight returns the RearSuspensionHeight of *CarSetupData
func (data *CarSetupData) GetRearSuspensionHeight() uint8 { return data.RearSuspensionHeight }

// SetRearSuspensionHeight stores the RearSuspensionHeight of *CarSetupData
func (data *CarSetupData) SetRearSuspensionHeight(v uint8) { data.RearSuspensionHeight = v }

// GetBrakePressure returns the BrakePressure of *CarSetupData
func (data *CarSetupData) GetBrakePressure() uint8 { return data.BrakePressure }

// SetBrakePressure stores the BrakePressure of *CarSetupData
func (data *CarSetupData) SetBrakePressure(v uint8) { data.BrakePressure = v }

// GetBrakeBias returns the BrakeBias of *CarSetupData
func (data *CarSetupData) GetBrakeBias() uint8 { return data.BrakeBias }

// SetBrakeBias stores the BrakeBias of *CarSetupData
func (data *CarSetupData) SetBrakeBias(v uint8) { data.BrakeBias = v }

// GetEngineBraking returns the EngineBraking of *CarSetupData
func (data *CarSetupData) GetEngineBraking() uint8 { return data.EngineBraking }

// SetEngineBraking stores the EngineBraking of *CarSetupData
func (data *CarSetupData) SetEngineBraking(v uint8) { data.EngineBraking = v }

// GetRearLeftTyrePressure returns the RearLeftTyrePressure of *CarSetupData
func (data *CarSetupData) GetRearLeftTyrePressure() float32 { return data.RearLeftTyrePressure }

// SetRearLeftTyrePressure stores the RearLeftTyrePressure of *CarSetupData
func (data *CarSetupData) SetRearLeftTyrePressure(v float32) { data.RearLeftTyrePressure = v }

// GetRearRightTyrePressure returns the RearRightTyrePressure of *CarSetupData
func (data *CarSetupData) GetRearRightTyrePressure() float32 { return data.RearRightTyrePressure }

// SetRearRightTyrePressure stores the RearRightTyrePressure of *CarSetupData
func (data *CarSetupData) SetRearRightTyrePressure(v float32) { data.RearRightTyrePressure = v }

// GetFrontLeftTyrePressure returns the FrontLeftTyrePressure of *CarSetupData
func (data *CarSetupData) GetFrontLeftTyrePressure() float32 { return data.FrontLeftTyrePressure }

// SetFrontLeftTyrePressure stores the FrontLeftTyrePressure of *CarSetupData
func (data *CarSetupData) SetFrontLeftTyrePressure(v float32) { data.FrontLeftTyrePressure = v }

// GetFrontRightTyrePressure returns the FrontRightTyrePressure of *CarSetupData
func (data *CarSetupData) GetFrontRightTyrePressure() float32 { return data.FrontRightTyrePressure }

// SetFrontRightTyrePressure stores the FrontRightTyrePressure of *CarSetupData
func (data *CarSetupData) SetFrontRightTyrePressure(v float32) { data.FrontRightTyrePressure = v }

// GetBallast returns the Ballast of *CarSetupData
func (data *CarSetupData) GetBallast() uint8 { return data.Ballast }

// SetBallast stores the Ballast of *CarSetupData
func (data *CarSetupData) SetBallast(v uint8) { data.Ballast = v }

// GetFuelLoad returns the FuelLoad of *CarSetupData
func (data *CarSetupData) GetFuelLoad() float32 { return data.FuelLoad }

// SetFuelLoad stores the FuelLoad of *CarSetupData
func (data *CarSetupData) SetFuelLoad(v float32) { data.FuelLoad = v }

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
