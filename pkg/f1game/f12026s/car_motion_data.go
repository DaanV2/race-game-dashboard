package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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
	GForceLateral      int16   // Lateral G-Force component (quantised)
	GForceLongitudinal int16   // Longitudinal G-Force component (quantised)
	GForceVertical     int16   // Vertical G-Force component (quantised)
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

// GetWorldPositionX returns the WorldPositionX of *CarMotionData
func (data *CarMotionData) GetWorldPositionX() float32 { return data.WorldPositionX }

// SetWorldPositionX stores the WorldPositionX of *CarMotionData
func (data *CarMotionData) SetWorldPositionX(v float32) { data.WorldPositionX = v }

// GetWorldPositionY returns the WorldPositionY of *CarMotionData
func (data *CarMotionData) GetWorldPositionY() float32 { return data.WorldPositionY }

// SetWorldPositionY stores the WorldPositionY of *CarMotionData
func (data *CarMotionData) SetWorldPositionY(v float32) { data.WorldPositionY = v }

// GetWorldPositionZ returns the WorldPositionZ of *CarMotionData
func (data *CarMotionData) GetWorldPositionZ() float32 { return data.WorldPositionZ }

// SetWorldPositionZ stores the WorldPositionZ of *CarMotionData
func (data *CarMotionData) SetWorldPositionZ(v float32) { data.WorldPositionZ = v }

// GetWorldVelocityX returns the WorldVelocityX of *CarMotionData
func (data *CarMotionData) GetWorldVelocityX() float32 { return data.WorldVelocityX }

// SetWorldVelocityX stores the WorldVelocityX of *CarMotionData
func (data *CarMotionData) SetWorldVelocityX(v float32) { data.WorldVelocityX = v }

// GetWorldVelocityY returns the WorldVelocityY of *CarMotionData
func (data *CarMotionData) GetWorldVelocityY() float32 { return data.WorldVelocityY }

// SetWorldVelocityY stores the WorldVelocityY of *CarMotionData
func (data *CarMotionData) SetWorldVelocityY(v float32) { data.WorldVelocityY = v }

// GetWorldVelocityZ returns the WorldVelocityZ of *CarMotionData
func (data *CarMotionData) GetWorldVelocityZ() float32 { return data.WorldVelocityZ }

// SetWorldVelocityZ stores the WorldVelocityZ of *CarMotionData
func (data *CarMotionData) SetWorldVelocityZ(v float32) { data.WorldVelocityZ = v }

// GetWorldForwardDirX returns the WorldForwardDirX of *CarMotionData
func (data *CarMotionData) GetWorldForwardDirX() int16 { return data.WorldForwardDirX }

// SetWorldForwardDirX stores the WorldForwardDirX of *CarMotionData
func (data *CarMotionData) SetWorldForwardDirX(v int16) { data.WorldForwardDirX = v }

// GetWorldForwardDirY returns the WorldForwardDirY of *CarMotionData
func (data *CarMotionData) GetWorldForwardDirY() int16 { return data.WorldForwardDirY }

// SetWorldForwardDirY stores the WorldForwardDirY of *CarMotionData
func (data *CarMotionData) SetWorldForwardDirY(v int16) { data.WorldForwardDirY = v }

// GetWorldForwardDirZ returns the WorldForwardDirZ of *CarMotionData
func (data *CarMotionData) GetWorldForwardDirZ() int16 { return data.WorldForwardDirZ }

// SetWorldForwardDirZ stores the WorldForwardDirZ of *CarMotionData
func (data *CarMotionData) SetWorldForwardDirZ(v int16) { data.WorldForwardDirZ = v }

// GetWorldRightDirX returns the WorldRightDirX of *CarMotionData
func (data *CarMotionData) GetWorldRightDirX() int16 { return data.WorldRightDirX }

// SetWorldRightDirX stores the WorldRightDirX of *CarMotionData
func (data *CarMotionData) SetWorldRightDirX(v int16) { data.WorldRightDirX = v }

// GetWorldRightDirY returns the WorldRightDirY of *CarMotionData
func (data *CarMotionData) GetWorldRightDirY() int16 { return data.WorldRightDirY }

// SetWorldRightDirY stores the WorldRightDirY of *CarMotionData
func (data *CarMotionData) SetWorldRightDirY(v int16) { data.WorldRightDirY = v }

// GetWorldRightDirZ returns the WorldRightDirZ of *CarMotionData
func (data *CarMotionData) GetWorldRightDirZ() int16 { return data.WorldRightDirZ }

// SetWorldRightDirZ stores the WorldRightDirZ of *CarMotionData
func (data *CarMotionData) SetWorldRightDirZ(v int16) { data.WorldRightDirZ = v }

// GetGForceLateral returns the GForceLateral of *CarMotionData
func (data *CarMotionData) GetGForceLateral() int16 { return data.GForceLateral }

// SetGForceLateral stores the GForceLateral of *CarMotionData
func (data *CarMotionData) SetGForceLateral(v int16) { data.GForceLateral = v }

// GetGForceLongitudinal returns the GForceLongitudinal of *CarMotionData
func (data *CarMotionData) GetGForceLongitudinal() int16 { return data.GForceLongitudinal }

// SetGForceLongitudinal stores the GForceLongitudinal of *CarMotionData
func (data *CarMotionData) SetGForceLongitudinal(v int16) { data.GForceLongitudinal = v }

// GetGForceVertical returns the GForceVertical of *CarMotionData
func (data *CarMotionData) GetGForceVertical() int16 { return data.GForceVertical }

// SetGForceVertical stores the GForceVertical of *CarMotionData
func (data *CarMotionData) SetGForceVertical(v int16) { data.GForceVertical = v }

// GetYaw returns the Yaw of *CarMotionData
func (data *CarMotionData) GetYaw() float32 { return data.Yaw }

// SetYaw stores the Yaw of *CarMotionData
func (data *CarMotionData) SetYaw(v float32) { data.Yaw = v }

// GetPitch returns the Pitch of *CarMotionData
func (data *CarMotionData) GetPitch() float32 { return data.Pitch }

// SetPitch stores the Pitch of *CarMotionData
func (data *CarMotionData) SetPitch(v float32) { data.Pitch = v }

// GetRoll returns the Roll of *CarMotionData
func (data *CarMotionData) GetRoll() float32 { return data.Roll }

// SetRoll stores the Roll of *CarMotionData
func (data *CarMotionData) SetRoll(v float32) { data.Roll = v }

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
	data.GForceLateral = reader.ReadInt16()
	data.GForceLongitudinal = reader.ReadInt16()
	data.GForceVertical = reader.ReadInt16()
	data.Yaw = reader.ReadFloat32()
	data.Pitch = reader.ReadFloat32()
	data.Roll = reader.ReadFloat32()

}
