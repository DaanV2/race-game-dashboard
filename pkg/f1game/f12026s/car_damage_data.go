package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarDamageData

type CarDamageData struct {
	TyresWear            WheelMap[float32] // Tyre wear (percentage)
	TyresDamage          WheelMap[uint8]   // Tyre damage (percentage)
	BrakesDamage         WheelMap[uint8]   // Brakes damage (percentage)
	TyreBlisters         WheelMap[uint8]   // Tyre blisters value (percentage)
	FrontLeftWingDamage  uint8             // Front left wing damage (percentage)
	FrontRightWingDamage uint8             // Front right wing damage (percentage)
	RearWingDamage       uint8             // Rear wing damage (percentage)
	FloorDamage          uint8             // Floor damage (percentage)
	DiffuserDamage       uint8             // Diffuser damage (percentage)
	SidepodDamage        uint8             // Sidepod damage (percentage)
	DrsFault             uint8             // Indicator for DRS fault, 0 = OK, 1 = fault
	ErsFault             uint8             // Indicator for ERS fault, 0 = OK, 1 = fault
	GearBoxDamage        uint8             // Gear box damage (percentage)
	EngineDamage         uint8             // Engine damage (percentage)
	EngineMGUHWear       uint8             // Engine wear MGU-H (percentage)
	EngineESWear         uint8             // Engine wear ES (percentage)
	EngineCEWear         uint8             // Engine wear CE (percentage)
	EngineICEWear        uint8             // Engine wear ICE (percentage)
	EngineMGUKWear       uint8             // Engine wear MGU-K (percentage)
	EngineTCWear         uint8             // Engine wear TC (percentage)
	EngineBlown          uint8             // Engine blown, 0 = OK, 1 = fault
	EngineSeized         uint8             // Engine seized, 0 = OK, 1 = fault
}

func (data *CarDamageData) Parse(reader *xbinary.LittleEndianReader) {
	data.TyresWear = reader.ReadFloat32x4()
	data.TyresDamage = reader.ReadUint8x4()
	data.BrakesDamage = reader.ReadUint8x4()
	data.TyreBlisters = reader.ReadUint8x4()
	data.FrontLeftWingDamage = reader.ReadUint8()
	data.FrontRightWingDamage = reader.ReadUint8()
	data.RearWingDamage = reader.ReadUint8()
	data.FloorDamage = reader.ReadUint8()
	data.DiffuserDamage = reader.ReadUint8()
	data.SidepodDamage = reader.ReadUint8()
	data.DrsFault = reader.ReadUint8()
	data.ErsFault = reader.ReadUint8()
	data.GearBoxDamage = reader.ReadUint8()
	data.EngineDamage = reader.ReadUint8()
	data.EngineMGUHWear = reader.ReadUint8()
	data.EngineESWear = reader.ReadUint8()
	data.EngineCEWear = reader.ReadUint8()
	data.EngineICEWear = reader.ReadUint8()
	data.EngineMGUKWear = reader.ReadUint8()
	data.EngineTCWear = reader.ReadUint8()
	data.EngineBlown = reader.ReadUint8()
	data.EngineSeized = reader.ReadUint8()
}
