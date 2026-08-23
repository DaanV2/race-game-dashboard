package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type PacketMotionExData struct {
	Header                 PacketHeader      // Header  Extra player car ONLY data
	SuspensionPosition     WheelMap[float32] // Note: All wheel arrays have the following order:
	SuspensionVelocity     WheelMap[float32] // RL, RR, FL, FR
	SuspensionAcceleration WheelMap[float32] // RL, RR, FL, FR
	WheelSpeed             WheelMap[float32] // Speed of each wheel
	WheelSlipRatio         WheelMap[float32] // Slip ratio for each wheel
	WheelSlipAngle         WheelMap[float32] // Slip angles for each wheel
	WheelLatForce          WheelMap[float32] // Lateral forces for each wheel
	WheelLongForce         WheelMap[float32] // Longitudinal forces for each wheel
	HeightOfCOGAboveGround float32           // Height of centre of gravity above ground
	LocalVelocityX         float32           // Velocity in local space – metres/s
	LocalVelocityY         float32           // Velocity in local space
	LocalVelocityZ         float32           // Velocity in local space
	AngularVelocityX       float32           // Angular velocity x-component – radians/s
	AngularVelocityY       float32           // Angular velocity y-component
	AngularVelocityZ       float32           // Angular velocity z-component
	AngularAccelerationX   float32           // Angular acceleration x-component – radians/s/s
	AngularAccelerationY   float32           // Angular acceleration y-component
	AngularAccelerationZ   float32           // Angular acceleration z-component
	FrontWheelsAngle       float32           // Current front wheels angle in radians
	WheelVertForce         WheelMap[float32] // Vertical forces for each wheel
	FrontAeroHeight        float32           // Front plank edge height above road surface
	RearAeroHeight         float32           // Rear plank edge height above road surface
	FrontRollAngle         float32           // Roll angle of the front suspension
	RearRollAngle          float32           // Roll angle of the rear suspension
	ChassisYaw             float32           // Yaw angle of the chassis relative to the direction  of motion - radians
	ChassisPitch           float32           // Pitch angle of the chassis relative to the  direction of motion – radians
	WheelCamber            WheelMap[float32] // Camber of each wheel in radians
	WheelCamberGain        WheelMap[float32] // Camber gain for each wheel in radians, difference  between active camber and dynamic camber
}

// GetHeader returns the Header of *PacketMotionExData
func (data *PacketMotionExData) GetHeader() PacketHeader { return data.Header }

// SetHeader stores the Header of *PacketMotionExData
func (data *PacketMotionExData) SetHeader(v PacketHeader) { data.Header = v }

// GetSuspensionPosition returns the SuspensionPosition of *PacketMotionExData
func (data *PacketMotionExData) GetSuspensionPosition() WheelMap[float32] {
	return data.SuspensionPosition
}

// SetSuspensionPosition stores the SuspensionPosition of *PacketMotionExData
func (data *PacketMotionExData) SetSuspensionPosition(v WheelMap[float32]) {
	data.SuspensionPosition = v
}

// GetSuspensionVelocity returns the SuspensionVelocity of *PacketMotionExData
func (data *PacketMotionExData) GetSuspensionVelocity() WheelMap[float32] {
	return data.SuspensionVelocity
}

// SetSuspensionVelocity stores the SuspensionVelocity of *PacketMotionExData
func (data *PacketMotionExData) SetSuspensionVelocity(v WheelMap[float32]) {
	data.SuspensionVelocity = v
}

// GetSuspensionAcceleration returns the SuspensionAcceleration of *PacketMotionExData
func (data *PacketMotionExData) GetSuspensionAcceleration() WheelMap[float32] {
	return data.SuspensionAcceleration
}

// SetSuspensionAcceleration stores the SuspensionAcceleration of *PacketMotionExData
func (data *PacketMotionExData) SetSuspensionAcceleration(v WheelMap[float32]) {
	data.SuspensionAcceleration = v
}

// GetWheelSpeed returns the WheelSpeed of *PacketMotionExData
func (data *PacketMotionExData) GetWheelSpeed() WheelMap[float32] { return data.WheelSpeed }

// SetWheelSpeed stores the WheelSpeed of *PacketMotionExData
func (data *PacketMotionExData) SetWheelSpeed(v WheelMap[float32]) { data.WheelSpeed = v }

// GetWheelSlipRatio returns the WheelSlipRatio of *PacketMotionExData
func (data *PacketMotionExData) GetWheelSlipRatio() WheelMap[float32] { return data.WheelSlipRatio }

// SetWheelSlipRatio stores the WheelSlipRatio of *PacketMotionExData
func (data *PacketMotionExData) SetWheelSlipRatio(v WheelMap[float32]) { data.WheelSlipRatio = v }

// GetWheelSlipAngle returns the WheelSlipAngle of *PacketMotionExData
func (data *PacketMotionExData) GetWheelSlipAngle() WheelMap[float32] { return data.WheelSlipAngle }

// SetWheelSlipAngle stores the WheelSlipAngle of *PacketMotionExData
func (data *PacketMotionExData) SetWheelSlipAngle(v WheelMap[float32]) { data.WheelSlipAngle = v }

// GetWheelLatForce returns the WheelLatForce of *PacketMotionExData
func (data *PacketMotionExData) GetWheelLatForce() WheelMap[float32] { return data.WheelLatForce }

// SetWheelLatForce stores the WheelLatForce of *PacketMotionExData
func (data *PacketMotionExData) SetWheelLatForce(v WheelMap[float32]) { data.WheelLatForce = v }

// GetWheelLongForce returns the WheelLongForce of *PacketMotionExData
func (data *PacketMotionExData) GetWheelLongForce() WheelMap[float32] { return data.WheelLongForce }

// SetWheelLongForce stores the WheelLongForce of *PacketMotionExData
func (data *PacketMotionExData) SetWheelLongForce(v WheelMap[float32]) { data.WheelLongForce = v }

// GetHeightOfCOGAboveGround returns the HeightOfCOGAboveGround of *PacketMotionExData
func (data *PacketMotionExData) GetHeightOfCOGAboveGround() float32 {
	return data.HeightOfCOGAboveGround
}

// SetHeightOfCOGAboveGround stores the HeightOfCOGAboveGround of *PacketMotionExData
func (data *PacketMotionExData) SetHeightOfCOGAboveGround(v float32) { data.HeightOfCOGAboveGround = v }

// GetLocalVelocityX returns the LocalVelocityX of *PacketMotionExData
func (data *PacketMotionExData) GetLocalVelocityX() float32 { return data.LocalVelocityX }

// SetLocalVelocityX stores the LocalVelocityX of *PacketMotionExData
func (data *PacketMotionExData) SetLocalVelocityX(v float32) { data.LocalVelocityX = v }

// GetLocalVelocityY returns the LocalVelocityY of *PacketMotionExData
func (data *PacketMotionExData) GetLocalVelocityY() float32 { return data.LocalVelocityY }

// SetLocalVelocityY stores the LocalVelocityY of *PacketMotionExData
func (data *PacketMotionExData) SetLocalVelocityY(v float32) { data.LocalVelocityY = v }

// GetLocalVelocityZ returns the LocalVelocityZ of *PacketMotionExData
func (data *PacketMotionExData) GetLocalVelocityZ() float32 { return data.LocalVelocityZ }

// SetLocalVelocityZ stores the LocalVelocityZ of *PacketMotionExData
func (data *PacketMotionExData) SetLocalVelocityZ(v float32) { data.LocalVelocityZ = v }

// GetAngularVelocityX returns the AngularVelocityX of *PacketMotionExData
func (data *PacketMotionExData) GetAngularVelocityX() float32 { return data.AngularVelocityX }

// SetAngularVelocityX stores the AngularVelocityX of *PacketMotionExData
func (data *PacketMotionExData) SetAngularVelocityX(v float32) { data.AngularVelocityX = v }

// GetAngularVelocityY returns the AngularVelocityY of *PacketMotionExData
func (data *PacketMotionExData) GetAngularVelocityY() float32 { return data.AngularVelocityY }

// SetAngularVelocityY stores the AngularVelocityY of *PacketMotionExData
func (data *PacketMotionExData) SetAngularVelocityY(v float32) { data.AngularVelocityY = v }

// GetAngularVelocityZ returns the AngularVelocityZ of *PacketMotionExData
func (data *PacketMotionExData) GetAngularVelocityZ() float32 { return data.AngularVelocityZ }

// SetAngularVelocityZ stores the AngularVelocityZ of *PacketMotionExData
func (data *PacketMotionExData) SetAngularVelocityZ(v float32) { data.AngularVelocityZ = v }

// GetAngularAccelerationX returns the AngularAccelerationX of *PacketMotionExData
func (data *PacketMotionExData) GetAngularAccelerationX() float32 { return data.AngularAccelerationX }

// SetAngularAccelerationX stores the AngularAccelerationX of *PacketMotionExData
func (data *PacketMotionExData) SetAngularAccelerationX(v float32) { data.AngularAccelerationX = v }

// GetAngularAccelerationY returns the AngularAccelerationY of *PacketMotionExData
func (data *PacketMotionExData) GetAngularAccelerationY() float32 { return data.AngularAccelerationY }

// SetAngularAccelerationY stores the AngularAccelerationY of *PacketMotionExData
func (data *PacketMotionExData) SetAngularAccelerationY(v float32) { data.AngularAccelerationY = v }

// GetAngularAccelerationZ returns the AngularAccelerationZ of *PacketMotionExData
func (data *PacketMotionExData) GetAngularAccelerationZ() float32 { return data.AngularAccelerationZ }

// SetAngularAccelerationZ stores the AngularAccelerationZ of *PacketMotionExData
func (data *PacketMotionExData) SetAngularAccelerationZ(v float32) { data.AngularAccelerationZ = v }

// GetFrontWheelsAngle returns the FrontWheelsAngle of *PacketMotionExData
func (data *PacketMotionExData) GetFrontWheelsAngle() float32 { return data.FrontWheelsAngle }

// SetFrontWheelsAngle stores the FrontWheelsAngle of *PacketMotionExData
func (data *PacketMotionExData) SetFrontWheelsAngle(v float32) { data.FrontWheelsAngle = v }

// GetWheelVertForce returns the WheelVertForce of *PacketMotionExData
func (data *PacketMotionExData) GetWheelVertForce() WheelMap[float32] { return data.WheelVertForce }

// SetWheelVertForce stores the WheelVertForce of *PacketMotionExData
func (data *PacketMotionExData) SetWheelVertForce(v WheelMap[float32]) { data.WheelVertForce = v }

// GetFrontAeroHeight returns the FrontAeroHeight of *PacketMotionExData
func (data *PacketMotionExData) GetFrontAeroHeight() float32 { return data.FrontAeroHeight }

// SetFrontAeroHeight stores the FrontAeroHeight of *PacketMotionExData
func (data *PacketMotionExData) SetFrontAeroHeight(v float32) { data.FrontAeroHeight = v }

// GetRearAeroHeight returns the RearAeroHeight of *PacketMotionExData
func (data *PacketMotionExData) GetRearAeroHeight() float32 { return data.RearAeroHeight }

// SetRearAeroHeight stores the RearAeroHeight of *PacketMotionExData
func (data *PacketMotionExData) SetRearAeroHeight(v float32) { data.RearAeroHeight = v }

// GetFrontRollAngle returns the FrontRollAngle of *PacketMotionExData
func (data *PacketMotionExData) GetFrontRollAngle() float32 { return data.FrontRollAngle }

// SetFrontRollAngle stores the FrontRollAngle of *PacketMotionExData
func (data *PacketMotionExData) SetFrontRollAngle(v float32) { data.FrontRollAngle = v }

// GetRearRollAngle returns the RearRollAngle of *PacketMotionExData
func (data *PacketMotionExData) GetRearRollAngle() float32 { return data.RearRollAngle }

// SetRearRollAngle stores the RearRollAngle of *PacketMotionExData
func (data *PacketMotionExData) SetRearRollAngle(v float32) { data.RearRollAngle = v }

// GetChassisYaw returns the ChassisYaw of *PacketMotionExData
func (data *PacketMotionExData) GetChassisYaw() float32 { return data.ChassisYaw }

// SetChassisYaw stores the ChassisYaw of *PacketMotionExData
func (data *PacketMotionExData) SetChassisYaw(v float32) { data.ChassisYaw = v }

// GetChassisPitch returns the ChassisPitch of *PacketMotionExData
func (data *PacketMotionExData) GetChassisPitch() float32 { return data.ChassisPitch }

// SetChassisPitch stores the ChassisPitch of *PacketMotionExData
func (data *PacketMotionExData) SetChassisPitch(v float32) { data.ChassisPitch = v }

// GetWheelCamber returns the WheelCamber of *PacketMotionExData
func (data *PacketMotionExData) GetWheelCamber() WheelMap[float32] { return data.WheelCamber }

// SetWheelCamber stores the WheelCamber of *PacketMotionExData
func (data *PacketMotionExData) SetWheelCamber(v WheelMap[float32]) { data.WheelCamber = v }

// GetWheelCamberGain returns the WheelCamberGain of *PacketMotionExData
func (data *PacketMotionExData) GetWheelCamberGain() WheelMap[float32] { return data.WheelCamberGain }

// SetWheelCamberGain stores the WheelCamberGain of *PacketMotionExData
func (data *PacketMotionExData) SetWheelCamberGain(v WheelMap[float32]) { data.WheelCamberGain = v }

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketMotionExData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.SuspensionPosition = xbinary.Readx4(reader.ReadFloat32)
	data.SuspensionVelocity = xbinary.Readx4(reader.ReadFloat32)
	data.SuspensionAcceleration = xbinary.Readx4(reader.ReadFloat32)
	data.WheelSpeed = xbinary.Readx4(reader.ReadFloat32)
	data.WheelSlipRatio = xbinary.Readx4(reader.ReadFloat32)
	data.WheelSlipAngle = xbinary.Readx4(reader.ReadFloat32)
	data.WheelLatForce = xbinary.Readx4(reader.ReadFloat32)
	data.WheelLongForce = xbinary.Readx4(reader.ReadFloat32)
	data.HeightOfCOGAboveGround = reader.ReadFloat32()
	data.LocalVelocityX = reader.ReadFloat32()
	data.LocalVelocityY = reader.ReadFloat32()
	data.LocalVelocityZ = reader.ReadFloat32()
	data.AngularVelocityX = reader.ReadFloat32()
	data.AngularVelocityY = reader.ReadFloat32()
	data.AngularVelocityZ = reader.ReadFloat32()
	data.AngularAccelerationX = reader.ReadFloat32()
	data.AngularAccelerationY = reader.ReadFloat32()
	data.AngularAccelerationZ = reader.ReadFloat32()
	data.FrontWheelsAngle = reader.ReadFloat32()
	data.WheelVertForce = xbinary.Readx4(reader.ReadFloat32)
	data.FrontAeroHeight = reader.ReadFloat32()
	data.RearAeroHeight = reader.ReadFloat32()
	data.FrontRollAngle = reader.ReadFloat32()
	data.RearRollAngle = reader.ReadFloat32()
	data.ChassisYaw = reader.ReadFloat32()
	data.ChassisPitch = reader.ReadFloat32()
	data.WheelCamber = xbinary.Readx4(reader.ReadFloat32)
	data.WheelCamberGain = xbinary.Readx4(reader.ReadFloat32)

}
