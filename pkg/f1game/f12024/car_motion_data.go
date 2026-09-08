package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/math/vectors"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarMotionData

type CarMotionData struct {
	WorldPositionX     float32 // World space X position - metres
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X – metres/s
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

func (data *CarMotionData) WorldPosition() vectors.Vec3[float32] {
	return vectors.NewVec3(data.WorldPositionX, data.WorldPositionY, data.WorldPositionZ)
}
func (data *CarMotionData) WorldVelocity() vectors.Vec3[float32] {
	return vectors.NewVec3(data.WorldVelocityX, data.WorldVelocityY, data.WorldVelocityZ)
}
func (data *CarMotionData) WorldForwardDir() vectors.Vec3[int16] {
	return vectors.NewVec3(data.WorldForwardDirX, data.WorldForwardDirY, data.WorldForwardDirZ)
}
func (data *CarMotionData) WorldRightDir() vectors.Vec3[int16] {
	return vectors.NewVec3(data.WorldRightDirX, data.WorldRightDirY, data.WorldRightDirZ)
}

func (data *CarMotionData) Parse(reader *xbinary.LittleEndianReader) {
	data.WorldPositionX = reader.ReadFloat32()
	data.WorldPositionY = reader.ReadFloat32()
	data.WorldPositionZ = reader.ReadFloat32()
	data.WorldVelocityX = reader.ReadFloat32()
	data.WorldVelocityY = reader.ReadFloat32()
	data.WorldVelocityZ = reader.ReadFloat32()
	data.WorldForwardDirX = reader.ReadInt16()
	data.WorldForwardDirY = reader.ReadInt16()
	data.WorldForwardDirZ = reader.ReadInt16()
	data.WorldRightDirX = reader.ReadInt16()
	data.WorldRightDirY = reader.ReadInt16()
	data.WorldRightDirZ = reader.ReadInt16()
	data.GForceLateral = reader.ReadFloat32()
	data.GForceLongitudinal = reader.ReadFloat32()
	data.GForceVertical = reader.ReadFloat32()
	data.Yaw = reader.ReadFloat32()
	data.Pitch = reader.ReadFloat32()
	data.Roll = reader.ReadFloat32()

}
