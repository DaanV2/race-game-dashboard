package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type CarTelemetryData struct {
	Speed                   uint16            // Speed of car in kilometres per hour
	Throttle                float32           // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32           // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32           // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8             // Amount of clutch applied (0 to 100)
	Gear                    int8              // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16            // Engine RPM
	DRS                     uint8             // 0 = off, 1 = on
	RevLightsPercent        uint8             // Rev lights indicator (percentage)
	RevLightsBitValue       uint16            // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       WheelMap[uint16]  // Brakes temperature (celsius)
	TyresSurfaceTemperature WheelMap[uint8]   // Tyres surface temperature (celsius)
	TyresInnerTemperature   WheelMap[uint8]   // Tyres inner temperature (celsius)
	EngineTemperature       uint8             // Engine temperature (celsius)
	TyresPressure           WheelMap[float32] // Tyres pressure (PSI)
	SurfaceType             WheelMap[uint8]   // Driving surface, see appendices
}

// GetSpeed returns the Speed of *CarTelemetryData
func (data *CarTelemetryData) GetSpeed() uint16 { return data.Speed }

// SetSpeed stores the Speed of *CarTelemetryData
func (data *CarTelemetryData) SetSpeed(v uint16) { data.Speed = v }

// GetThrottle returns the Throttle of *CarTelemetryData
func (data *CarTelemetryData) GetThrottle() float32 { return data.Throttle }

// SetThrottle stores the Throttle of *CarTelemetryData
func (data *CarTelemetryData) SetThrottle(v float32) { data.Throttle = v }

// GetSteer returns the Steer of *CarTelemetryData
func (data *CarTelemetryData) GetSteer() float32 { return data.Steer }

// SetSteer stores the Steer of *CarTelemetryData
func (data *CarTelemetryData) SetSteer(v float32) { data.Steer = v }

// GetBrake returns the Brake of *CarTelemetryData
func (data *CarTelemetryData) GetBrake() float32 { return data.Brake }

// SetBrake stores the Brake of *CarTelemetryData
func (data *CarTelemetryData) SetBrake(v float32) { data.Brake = v }

// GetClutch returns the Clutch of *CarTelemetryData
func (data *CarTelemetryData) GetClutch() uint8 { return data.Clutch }

// SetClutch stores the Clutch of *CarTelemetryData
func (data *CarTelemetryData) SetClutch(v uint8) { data.Clutch = v }

// GetGear returns the Gear of *CarTelemetryData
func (data *CarTelemetryData) GetGear() int8 { return data.Gear }

// SetGear stores the Gear of *CarTelemetryData
func (data *CarTelemetryData) SetGear(v int8) { data.Gear = v }

// GetEngineRPM returns the EngineRPM of *CarTelemetryData
func (data *CarTelemetryData) GetEngineRPM() uint16 { return data.EngineRPM }

// SetEngineRPM stores the EngineRPM of *CarTelemetryData
func (data *CarTelemetryData) SetEngineRPM(v uint16) { data.EngineRPM = v }

// GetDrs returns the Drs of *CarTelemetryData
func (data *CarTelemetryData) GetDrs() uint8 { return data.DRS }

// SetDrs stores the Drs of *CarTelemetryData
func (data *CarTelemetryData) SetDrs(v uint8) { data.DRS = v }

// GetRevLightsPercent returns the RevLightsPercent of *CarTelemetryData
func (data *CarTelemetryData) GetRevLightsPercent() uint8 { return data.RevLightsPercent }

// SetRevLightsPercent stores the RevLightsPercent of *CarTelemetryData
func (data *CarTelemetryData) SetRevLightsPercent(v uint8) { data.RevLightsPercent = v }

// GetRevLightsBitValue returns the RevLightsBitValue of *CarTelemetryData
func (data *CarTelemetryData) GetRevLightsBitValue() uint16 { return data.RevLightsBitValue }

// SetRevLightsBitValue stores the RevLightsBitValue of *CarTelemetryData
func (data *CarTelemetryData) SetRevLightsBitValue(v uint16) { data.RevLightsBitValue = v }

// GetBrakesTemperature returns the BrakesTemperature of *CarTelemetryData
func (data *CarTelemetryData) GetBrakesTemperature() WheelMap[uint16] { return data.BrakesTemperature }

// SetBrakesTemperature stores the BrakesTemperature of *CarTelemetryData
func (data *CarTelemetryData) SetBrakesTemperature(v WheelMap[uint16]) { data.BrakesTemperature = v }

// GetTyresSurfaceTemperature returns the TyresSurfaceTemperature of *CarTelemetryData
func (data *CarTelemetryData) GetTyresSurfaceTemperature() WheelMap[uint8] {
	return data.TyresSurfaceTemperature
}

// SetTyresSurfaceTemperature stores the TyresSurfaceTemperature of *CarTelemetryData
func (data *CarTelemetryData) SetTyresSurfaceTemperature(v WheelMap[uint8]) {
	data.TyresSurfaceTemperature = v
}

// GetTyresInnerTemperature returns the TyresInnerTemperature of *CarTelemetryData
func (data *CarTelemetryData) GetTyresInnerTemperature() WheelMap[uint8] {
	return data.TyresInnerTemperature
}

// SetTyresInnerTemperature stores the TyresInnerTemperature of *CarTelemetryData
func (data *CarTelemetryData) SetTyresInnerTemperature(v WheelMap[uint8]) {
	data.TyresInnerTemperature = v
}

// GetEngineTemperature returns the EngineTemperature of *CarTelemetryData
func (data *CarTelemetryData) GetEngineTemperature() uint8 { return data.EngineTemperature }

// SetEngineTemperature stores the EngineTemperature of *CarTelemetryData
func (data *CarTelemetryData) SetEngineTemperature(v uint8) { data.EngineTemperature = v }

// GetTyresPressure returns the TyresPressure of *CarTelemetryData
func (data *CarTelemetryData) GetTyresPressure() WheelMap[float32] { return data.TyresPressure }

// SetTyresPressure stores the TyresPressure of *CarTelemetryData
func (data *CarTelemetryData) SetTyresPressure(v WheelMap[float32]) { data.TyresPressure = v }

// GetSurfaceType returns the SurfaceType of *CarTelemetryData
func (data *CarTelemetryData) GetSurfaceType() WheelMap[uint8] { return data.SurfaceType }

// SetSurfaceType stores the SurfaceType of *CarTelemetryData
func (data *CarTelemetryData) SetSurfaceType(v WheelMap[uint8]) { data.SurfaceType = v }

func (data *CarTelemetryData) Parse(reader *xbinary.LittleEndianReader) {
	data.Speed = reader.ReadUint16()
	data.Throttle = reader.ReadFloat32()
	data.Steer = reader.ReadFloat32()
	data.Brake = reader.ReadFloat32()
	data.Clutch = reader.ReadUint8()
	data.Gear = reader.ReadInt8()
	data.EngineRPM = reader.ReadUint16()
	data.DRS = reader.ReadUint8()
	data.RevLightsPercent = reader.ReadUint8()
	data.RevLightsBitValue = reader.ReadUint16()
	data.BrakesTemperature = xbinary.Readx4(reader.ReadUint16)
	data.TyresSurfaceTemperature = xbinary.Readx4(reader.ReadUint8)
	data.TyresInnerTemperature = xbinary.Readx4(reader.ReadUint8)
	data.EngineTemperature = reader.ReadUint8()
	data.TyresPressure = xbinary.Readx4(reader.ReadFloat32)
	data.SurfaceType = xbinary.Readx4(reader.ReadUint8)

}
