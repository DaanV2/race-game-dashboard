package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetTyresWear returns the TyresWear of *CarDamageData
func (data *CarDamageData) GetTyresWear() WheelMap[float32] { return data.TyresWear }

// SetTyresWear stores the TyresWear of *CarDamageData
func (data *CarDamageData) SetTyresWear(v WheelMap[float32]) { data.TyresWear = v }

// GetTyresDamage returns the TyresDamage of *CarDamageData
func (data *CarDamageData) GetTyresDamage() WheelMap[uint8] { return data.TyresDamage }

// SetTyresDamage stores the TyresDamage of *CarDamageData
func (data *CarDamageData) SetTyresDamage(v WheelMap[uint8]) { data.TyresDamage = v }

// GetBrakesDamage returns the BrakesDamage of *CarDamageData
func (data *CarDamageData) GetBrakesDamage() WheelMap[uint8] { return data.BrakesDamage }

// SetBrakesDamage stores the BrakesDamage of *CarDamageData
func (data *CarDamageData) SetBrakesDamage(v WheelMap[uint8]) { data.BrakesDamage = v }

// GetTyreBlisters returns the TyreBlisters of *CarDamageData
func (data *CarDamageData) GetTyreBlisters() WheelMap[uint8] { return data.TyreBlisters }

// SetTyreBlisters stores the TyreBlisters of *CarDamageData
func (data *CarDamageData) SetTyreBlisters(v WheelMap[uint8]) { data.TyreBlisters = v }

// GetFrontLeftWingDamage returns the FrontLeftWingDamage of *CarDamageData
func (data *CarDamageData) GetFrontLeftWingDamage() uint8 { return data.FrontLeftWingDamage }

// SetFrontLeftWingDamage stores the FrontLeftWingDamage of *CarDamageData
func (data *CarDamageData) SetFrontLeftWingDamage(v uint8) { data.FrontLeftWingDamage = v }

// GetFrontRightWingDamage returns the FrontRightWingDamage of *CarDamageData
func (data *CarDamageData) GetFrontRightWingDamage() uint8 { return data.FrontRightWingDamage }

// SetFrontRightWingDamage stores the FrontRightWingDamage of *CarDamageData
func (data *CarDamageData) SetFrontRightWingDamage(v uint8) { data.FrontRightWingDamage = v }

// GetRearWingDamage returns the RearWingDamage of *CarDamageData
func (data *CarDamageData) GetRearWingDamage() uint8 { return data.RearWingDamage }

// SetRearWingDamage stores the RearWingDamage of *CarDamageData
func (data *CarDamageData) SetRearWingDamage(v uint8) { data.RearWingDamage = v }

// GetFloorDamage returns the FloorDamage of *CarDamageData
func (data *CarDamageData) GetFloorDamage() uint8 { return data.FloorDamage }

// SetFloorDamage stores the FloorDamage of *CarDamageData
func (data *CarDamageData) SetFloorDamage(v uint8) { data.FloorDamage = v }

// GetDiffuserDamage returns the DiffuserDamage of *CarDamageData
func (data *CarDamageData) GetDiffuserDamage() uint8 { return data.DiffuserDamage }

// SetDiffuserDamage stores the DiffuserDamage of *CarDamageData
func (data *CarDamageData) SetDiffuserDamage(v uint8) { data.DiffuserDamage = v }

// GetSidepodDamage returns the SidepodDamage of *CarDamageData
func (data *CarDamageData) GetSidepodDamage() uint8 { return data.SidepodDamage }

// SetSidepodDamage stores the SidepodDamage of *CarDamageData
func (data *CarDamageData) SetSidepodDamage(v uint8) { data.SidepodDamage = v }

// GetDrsFault returns the DrsFault of *CarDamageData
func (data *CarDamageData) GetDrsFault() uint8 { return data.DrsFault }

// SetDrsFault stores the DrsFault of *CarDamageData
func (data *CarDamageData) SetDrsFault(v uint8) { data.DrsFault = v }

// GetErsFault returns the ErsFault of *CarDamageData
func (data *CarDamageData) GetErsFault() uint8 { return data.ErsFault }

// SetErsFault stores the ErsFault of *CarDamageData
func (data *CarDamageData) SetErsFault(v uint8) { data.ErsFault = v }

// GetGearBoxDamage returns the GearBoxDamage of *CarDamageData
func (data *CarDamageData) GetGearBoxDamage() uint8 { return data.GearBoxDamage }

// SetGearBoxDamage stores the GearBoxDamage of *CarDamageData
func (data *CarDamageData) SetGearBoxDamage(v uint8) { data.GearBoxDamage = v }

// GetEngineDamage returns the EngineDamage of *CarDamageData
func (data *CarDamageData) GetEngineDamage() uint8 { return data.EngineDamage }

// SetEngineDamage stores the EngineDamage of *CarDamageData
func (data *CarDamageData) SetEngineDamage(v uint8) { data.EngineDamage = v }

// GetEngineMGUHWear returns the EngineMGUHWear of *CarDamageData
func (data *CarDamageData) GetEngineMGUHWear() uint8 { return data.EngineMGUHWear }

// SetEngineMGUHWear stores the EngineMGUHWear of *CarDamageData
func (data *CarDamageData) SetEngineMGUHWear(v uint8) { data.EngineMGUHWear = v }

// GetEngineESWear returns the EngineESWear of *CarDamageData
func (data *CarDamageData) GetEngineESWear() uint8 { return data.EngineESWear }

// SetEngineESWear stores the EngineESWear of *CarDamageData
func (data *CarDamageData) SetEngineESWear(v uint8) { data.EngineESWear = v }

// GetEngineCEWear returns the EngineCEWear of *CarDamageData
func (data *CarDamageData) GetEngineCEWear() uint8 { return data.EngineCEWear }

// SetEngineCEWear stores the EngineCEWear of *CarDamageData
func (data *CarDamageData) SetEngineCEWear(v uint8) { data.EngineCEWear = v }

// GetEngineICEWear returns the EngineICEWear of *CarDamageData
func (data *CarDamageData) GetEngineICEWear() uint8 { return data.EngineICEWear }

// SetEngineICEWear stores the EngineICEWear of *CarDamageData
func (data *CarDamageData) SetEngineICEWear(v uint8) { data.EngineICEWear = v }

// GetEngineMGUKWear returns the EngineMGUKWear of *CarDamageData
func (data *CarDamageData) GetEngineMGUKWear() uint8 { return data.EngineMGUKWear }

// SetEngineMGUKWear stores the EngineMGUKWear of *CarDamageData
func (data *CarDamageData) SetEngineMGUKWear(v uint8) { data.EngineMGUKWear = v }

// GetEngineTCWear returns the EngineTCWear of *CarDamageData
func (data *CarDamageData) GetEngineTCWear() uint8 { return data.EngineTCWear }

// SetEngineTCWear stores the EngineTCWear of *CarDamageData
func (data *CarDamageData) SetEngineTCWear(v uint8) { data.EngineTCWear = v }

// GetEngineBlown returns the EngineBlown of *CarDamageData
func (data *CarDamageData) GetEngineBlown() uint8 { return data.EngineBlown }

// SetEngineBlown stores the EngineBlown of *CarDamageData
func (data *CarDamageData) SetEngineBlown(v uint8) { data.EngineBlown = v }

// GetEngineSeized returns the EngineSeized of *CarDamageData
func (data *CarDamageData) GetEngineSeized() uint8 { return data.EngineSeized }

// SetEngineSeized stores the EngineSeized of *CarDamageData
func (data *CarDamageData) SetEngineSeized(v uint8) { data.EngineSeized = v }

func (data *CarDamageData) Parse(reader *xbinary.LittleEndianReader) {
	data.TyresWear = reader.ReadFloat32x4()
	reader.Read(data.TyresDamage[:])
	reader.Read(data.BrakesDamage[:])
	reader.Read(data.TyreBlisters[:])
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
