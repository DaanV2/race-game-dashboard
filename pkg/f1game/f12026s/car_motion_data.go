package f12026s

type CarMotionData struct {
	m_worldPositionX     float32 // World space X position - metres
	m_worldPositionY     float32 // World space Y position
	m_worldPositionZ     float32 // World space Z position
	m_worldVelocityX     float32 // Velocity in world space X – metres/s
	m_worldVelocityY     float32 // Velocity in world space Y
	m_worldVelocityZ     float32 // Velocity in world space Z
	m_worldForwardDirX   int16   // World space forward X direction (normalised)
	m_worldForwardDirY   int16   // World space forward Y direction (normalised)
	m_worldForwardDirZ   int16   // World space forward Z direction (normalised)
	m_worldRightDirX     int16   // World space right X direction (normalised)
	m_worldRightDirY     int16   // World space right Y direction (normalised)
	m_worldRightDirZ     int16   // World space right Z direction (normalised)
	m_gForceLateral      int16   // Lateral G-Force component (quantised)
	m_gForceLongitudinal int16   // Longitudinal G-Force component (quantised)
	m_gForceVertical     int16   // Vertical G-Force component (quantised)
	m_yaw                float32 // Yaw angle in radians
	m_pitch              float32 // Pitch angle in radians
	m_roll               float32 // Roll angle in radians
}

type PacketMotionData struct {
	m_header PacketHeader // Header

	m_carMotionData [24]CarMotionData // Data for all cars on track
}
