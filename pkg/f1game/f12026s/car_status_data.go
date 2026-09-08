package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

//go:generate go run github.com/daanv2/race-game-dashboard/tools/gen/accessors -type CarStatusData

type CarStatusData struct {
	TractionControl         uint8   // Traction control - 0 = off, 1 = medium, 2 = full
	AntiLockBrakes          uint8   // 0 (off) - 1 (on)
	FuelMix                 uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8   // Front brake bias (percentage)
	PitLimiterStatus        uint8   // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32 // Current fuel mass
	FuelCapacity            float32 // Fuel capacity
	FuelRemainingLaps       float32 // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16  // Cars idle RPM
	MaxGears                uint8   // Maximum number of gears
	DrsAllowed              uint8   // 0 = not allowed, 1 = allowed
	DrsActivationDistance   uint16  // 0 = DRS not available, non-zero - DRS will be available  in [X] metres
	ActualTyreCompound      uint8   // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1  21 = C0, 22 = C6, 7 = inter, 8 = wet  F1 Classic - 9 = dry, 10 = wet  F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard  15 = wet
	VisualTyreCompound      uint8   // F1 visual (can be different from actual compound)  16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet  F1 Classic – same as above  F2 ‘20, 15 = wet, 19 – super soft, 20 = soft  21 = medium, 22 = hard
	TyresAgeLaps            uint8   // Age in laps of the current set of tyres
	VehicleFiaFlags         int8    // -1 = invalid/unknown, 0 = none, 1 = green  2 = blue, 3 = yellow
	EnginePowerICE          float32 // Engine power output of ICE (W)
	EnginePowerMGUK         float32 // Engine power output of MGU-K (W)
	ErsStoreEnergy          float32 // ERS energy store in Joules
	ErsDeployMode           uint8   // ERS deployment mode, 0 = none, 1 = medium  2 = hotlap, 3 = boost
	ErsHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	ErsHarvestedLimitPerLap float32 // ERS energy harvest limit for this lap
	ErsDeployedThisLap      float32 // ERS energy deployed this lap
	NetworkPaused           uint8   // Whether the car is paused in a network game
}

func (data *CarStatusData) Parse(reader *xbinary.LittleEndianReader) {
	data.TractionControl = reader.ReadUint8()
	data.AntiLockBrakes = reader.ReadUint8()
	data.FuelMix = reader.ReadUint8()
	data.FrontBrakeBias = reader.ReadUint8()
	data.PitLimiterStatus = reader.ReadUint8()
	data.FuelInTank = reader.ReadFloat32()
	data.FuelCapacity = reader.ReadFloat32()
	data.FuelRemainingLaps = reader.ReadFloat32()
	data.MaxRPM = reader.ReadUint16()
	data.IdleRPM = reader.ReadUint16()
	data.MaxGears = reader.ReadUint8()
	data.DrsAllowed = reader.ReadUint8()
	data.DrsActivationDistance = reader.ReadUint16()
	data.ActualTyreCompound = reader.ReadUint8()
	data.VisualTyreCompound = reader.ReadUint8()
	data.TyresAgeLaps = reader.ReadUint8()
	data.VehicleFiaFlags = reader.ReadInt8()
	data.EnginePowerICE = reader.ReadFloat32()
	data.EnginePowerMGUK = reader.ReadFloat32()
	data.ErsStoreEnergy = reader.ReadFloat32()
	data.ErsDeployMode = reader.ReadUint8()
	data.ErsHarvestedThisLapMGUK = reader.ReadFloat32()
	data.ErsHarvestedThisLapMGUH = reader.ReadFloat32()
	data.ErsHarvestedLimitPerLap = reader.ReadFloat32()
	data.ErsDeployedThisLap = reader.ReadFloat32()
	data.NetworkPaused = reader.ReadUint8()

}
