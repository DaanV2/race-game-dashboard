package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/math/vectors"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type PacketMotionExData

type PacketMotionExData struct {
	Header PacketHeader // Header  Extra player car ONLY data

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

// GetPacketID returns the identification of this packet
func (data *PacketMotionExData) GetPacketID() PacketID { return PACKET_ID_MOTION_EX }

func (data *PacketMotionExData) LocalVelocity() vectors.Vec3[float32] {
	return vectors.NewVec3(data.LocalVelocityX, data.LocalVelocityY, data.LocalVelocityZ)
}
func (data *PacketMotionExData) AngularVelocity() vectors.Vec3[float32] {
	return vectors.NewVec3(data.AngularVelocityX, data.AngularVelocityY, data.AngularVelocityZ)
}
func (data *PacketMotionExData) AngularAcceleration() vectors.Vec3[float32] {
	return vectors.NewVec3(data.AngularAccelerationX, data.AngularAccelerationY, data.AngularAccelerationZ)
}

// Parse assumes the header as already been read, and only the rest needs to be done
func (data *PacketMotionExData) Parse(header *PacketHeader, reader *xbinary.LittleEndianReader) {
	data.Header = *header
	data.SuspensionPosition = reader.ReadFloat32x4()
	data.SuspensionVelocity = reader.ReadFloat32x4()
	data.SuspensionAcceleration = reader.ReadFloat32x4()
	data.WheelSpeed = reader.ReadFloat32x4()
	data.WheelSlipRatio = reader.ReadFloat32x4()
	data.WheelSlipAngle = reader.ReadFloat32x4()
	data.WheelLatForce = reader.ReadFloat32x4()
	data.WheelLongForce = reader.ReadFloat32x4()
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
	data.WheelVertForce = reader.ReadFloat32x4()
	data.FrontAeroHeight = reader.ReadFloat32()
	data.RearAeroHeight = reader.ReadFloat32()
	data.FrontRollAngle = reader.ReadFloat32()
	data.RearRollAngle = reader.ReadFloat32()
	data.ChassisYaw = reader.ReadFloat32()
	data.ChassisPitch = reader.ReadFloat32()
	data.WheelCamber = reader.ReadFloat32x4()
	data.WheelCamberGain = reader.ReadFloat32x4()

}
