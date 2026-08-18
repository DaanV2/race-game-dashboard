package f12026s

type CarSetupData struct {
	m_frontWing              uint8   // Front wing aero
	m_rearWing               uint8   // Rear wing aero
	m_onThrottle             uint8   // Differential adjustment on throttle (percentage)
	m_offThrottle            uint8   // Differential adjustment off throttle (percentage)
	m_frontCamber            float32 // Front camber angle (suspension geometry)
	m_rearCamber             float32 // Rear camber angle (suspension geometry)
	m_frontToe               float32 // Front toe angle (suspension geometry)
	m_rearToe                float32 // Rear toe angle (suspension geometry)
	m_frontSuspension        uint8   // Front suspension
	m_rearSuspension         uint8   // Rear suspension
	m_frontAntiRollBar       uint8   // Front anti-roll bar
	m_rearAntiRollBar        uint8   // Front anti-roll bar
	m_frontSuspensionHeight  uint8   // Front ride height
	m_rearSuspensionHeight   uint8   // Rear ride height
	m_brakePressure          uint8   // Brake pressure (percentage)
	m_brakeBias              uint8   // Brake bias (percentage)
	m_engineBraking          uint8   // Engine braking (percentage)
	m_rearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
	m_rearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
	m_frontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
	m_frontRightTyrePressure float32 // Front right tyre pressure (PSI)
	m_ballast                uint8   // Ballast
	m_fuelLoad               float32 // Fuel load
}

type PacketCarSetupData struct {
	m_header             PacketHeader // Header
	m_carSetups          [24]CarSetupData
	m_nextFrontWingValue float32 // Value of front wing after next pit stop - player only
}
