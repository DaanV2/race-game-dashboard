package f12024

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
	"github.com/daanv2/race-game-dashboard/pkg/f1game/f1common"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarTelemetryData

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
	EngineTemperature       uint16            // Engine temperature (celsius)
	TyresPressure           WheelMap[float32] // Tyres pressure (PSI)
	SurfaceType             WheelMap[uint8]   // Driving surface, see appendices
}

func (data *CarTelemetryData) SurfaceTypeID() WheelMap[f1common.SurfaceTypeID] {
	return TransformWheelMap[uint8, f1common.SurfaceTypeID](data.SurfaceType)
}

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
	data.BrakesTemperature = reader.ReadUint16x4()
	reader.Read(data.TyresSurfaceTemperature[:])
	reader.Read(data.TyresInnerTemperature[:])
	data.EngineTemperature = reader.ReadUint16()
	data.TyresPressure = reader.ReadFloat32x4()
	reader.Read(data.SurfaceType[:])
}
