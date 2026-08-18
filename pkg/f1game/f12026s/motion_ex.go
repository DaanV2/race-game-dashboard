package f12026s

type PacketMotionExData struct {
	m_header PacketHeader // Header

	// Extra player car ONLY data
	m_suspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	m_suspensionVelocity     [4]float32 // RL, RR, FL, FR
	m_suspensionAcceleration [4]float32 // RL, RR, FL, FR
	m_wheelSpeed             [4]float32 // Speed of each wheel
	m_wheelSlipRatio         [4]float32 // Slip ratio for each wheel
	m_wheelSlipAngle         [4]float32 // Slip angles for each wheel
	m_wheelLatForce          [4]float32 // Lateral forces for each wheel
	m_wheelLongForce         [4]float32 // Longitudinal forces for each wheel
	m_heightOfCOGAboveGround float32    // Height of centre of gravity above ground
	m_localVelocityX         float32    // Velocity in local space – metres/s
	m_localVelocityY         float32    // Velocity in local space
	m_localVelocityZ         float32    // Velocity in local space
	m_angularVelocityX       float32    // Angular velocity x-component – radians/s
	m_angularVelocityY       float32    // Angular velocity y-component
	m_angularVelocityZ       float32    // Angular velocity z-component
	m_angularAccelerationX   float32    // Angular acceleration x-component – radians/s/s
	m_angularAccelerationY   float32    // Angular acceleration y-component
	m_angularAccelerationZ   float32    // Angular acceleration z-component
	m_frontWheelsAngle       float32    // Current front wheels angle in radians
	m_wheelVertForce         [4]float32 // Vertical forces for each wheel
	m_frontAeroHeight        float32    // Front plank edge height above road surface
	m_rearAeroHeight         float32    // Rear plank edge height above road surface
	m_frontRollAngle         float32    // Roll angle of the front suspension
	m_rearRollAngle          float32    // Roll angle of the rear suspension
	m_chassisYaw             float32    // Yaw angle of the chassis relative to the direction of motion - radians
	m_chassisPitch           float32    // Pitch angle of the chassis relative to the direction of motion – radians
	m_wheelCamber            [4]float32 // Camber of each wheel in radians
	m_wheelCamberGain        [4]float32 // Camber gain for each wheel in radians, difference  between active camber and dynamic camber
}
