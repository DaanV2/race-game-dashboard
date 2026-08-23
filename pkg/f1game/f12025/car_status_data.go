package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

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

// GetTractionControl returns the TractionControl of *CarStatusData
func (data *CarStatusData) GetTractionControl() uint8 { return data.TractionControl }

// SetTractionControl stores the TractionControl of *CarStatusData
func (data *CarStatusData) SetTractionControl(v uint8) { data.TractionControl = v }

// GetAntiLockBrakes returns the AntiLockBrakes of *CarStatusData
func (data *CarStatusData) GetAntiLockBrakes() uint8 { return data.AntiLockBrakes }

// SetAntiLockBrakes stores the AntiLockBrakes of *CarStatusData
func (data *CarStatusData) SetAntiLockBrakes(v uint8) { data.AntiLockBrakes = v }

// GetFuelMix returns the FuelMix of *CarStatusData
func (data *CarStatusData) GetFuelMix() uint8 { return data.FuelMix }

// SetFuelMix stores the FuelMix of *CarStatusData
func (data *CarStatusData) SetFuelMix(v uint8) { data.FuelMix = v }

// GetFrontBrakeBias returns the FrontBrakeBias of *CarStatusData
func (data *CarStatusData) GetFrontBrakeBias() uint8 { return data.FrontBrakeBias }

// SetFrontBrakeBias stores the FrontBrakeBias of *CarStatusData
func (data *CarStatusData) SetFrontBrakeBias(v uint8) { data.FrontBrakeBias = v }

// GetPitLimiterStatus returns the PitLimiterStatus of *CarStatusData
func (data *CarStatusData) GetPitLimiterStatus() uint8 { return data.PitLimiterStatus }

// SetPitLimiterStatus stores the PitLimiterStatus of *CarStatusData
func (data *CarStatusData) SetPitLimiterStatus(v uint8) { data.PitLimiterStatus = v }

// GetFuelInTank returns the FuelInTank of *CarStatusData
func (data *CarStatusData) GetFuelInTank() float32 { return data.FuelInTank }

// SetFuelInTank stores the FuelInTank of *CarStatusData
func (data *CarStatusData) SetFuelInTank(v float32) { data.FuelInTank = v }

// GetFuelCapacity returns the FuelCapacity of *CarStatusData
func (data *CarStatusData) GetFuelCapacity() float32 { return data.FuelCapacity }

// SetFuelCapacity stores the FuelCapacity of *CarStatusData
func (data *CarStatusData) SetFuelCapacity(v float32) { data.FuelCapacity = v }

// GetFuelRemainingLaps returns the FuelRemainingLaps of *CarStatusData
func (data *CarStatusData) GetFuelRemainingLaps() float32 { return data.FuelRemainingLaps }

// SetFuelRemainingLaps stores the FuelRemainingLaps of *CarStatusData
func (data *CarStatusData) SetFuelRemainingLaps(v float32) { data.FuelRemainingLaps = v }

// GetMaxRPM returns the MaxRPM of *CarStatusData
func (data *CarStatusData) GetMaxRPM() uint16 { return data.MaxRPM }

// SetMaxRPM stores the MaxRPM of *CarStatusData
func (data *CarStatusData) SetMaxRPM(v uint16) { data.MaxRPM = v }

// GetIdleRPM returns the IdleRPM of *CarStatusData
func (data *CarStatusData) GetIdleRPM() uint16 { return data.IdleRPM }

// SetIdleRPM stores the IdleRPM of *CarStatusData
func (data *CarStatusData) SetIdleRPM(v uint16) { data.IdleRPM = v }

// GetMaxGears returns the MaxGears of *CarStatusData
func (data *CarStatusData) GetMaxGears() uint8 { return data.MaxGears }

// SetMaxGears stores the MaxGears of *CarStatusData
func (data *CarStatusData) SetMaxGears(v uint8) { data.MaxGears = v }

// GetDrsAllowed returns the DrsAllowed of *CarStatusData
func (data *CarStatusData) GetDrsAllowed() uint8 { return data.DrsAllowed }

// SetDrsAllowed stores the DrsAllowed of *CarStatusData
func (data *CarStatusData) SetDrsAllowed(v uint8) { data.DrsAllowed = v }

// GetDrsActivationDistance returns the DrsActivationDistance of *CarStatusData
func (data *CarStatusData) GetDrsActivationDistance() uint16 { return data.DrsActivationDistance }

// SetDrsActivationDistance stores the DrsActivationDistance of *CarStatusData
func (data *CarStatusData) SetDrsActivationDistance(v uint16) { data.DrsActivationDistance = v }

// GetActualTyreCompound returns the ActualTyreCompound of *CarStatusData
func (data *CarStatusData) GetActualTyreCompound() uint8 { return data.ActualTyreCompound }

// SetActualTyreCompound stores the ActualTyreCompound of *CarStatusData
func (data *CarStatusData) SetActualTyreCompound(v uint8) { data.ActualTyreCompound = v }

// GetVisualTyreCompound returns the VisualTyreCompound of *CarStatusData
func (data *CarStatusData) GetVisualTyreCompound() uint8 { return data.VisualTyreCompound }

// SetVisualTyreCompound stores the VisualTyreCompound of *CarStatusData
func (data *CarStatusData) SetVisualTyreCompound(v uint8) { data.VisualTyreCompound = v }

// GetTyresAgeLaps returns the TyresAgeLaps of *CarStatusData
func (data *CarStatusData) GetTyresAgeLaps() uint8 { return data.TyresAgeLaps }

// SetTyresAgeLaps stores the TyresAgeLaps of *CarStatusData
func (data *CarStatusData) SetTyresAgeLaps(v uint8) { data.TyresAgeLaps = v }

// GetVehicleFiaFlags returns the VehicleFiaFlags of *CarStatusData
func (data *CarStatusData) GetVehicleFiaFlags() int8 { return data.VehicleFiaFlags }

// SetVehicleFiaFlags stores the VehicleFiaFlags of *CarStatusData
func (data *CarStatusData) SetVehicleFiaFlags(v int8) { data.VehicleFiaFlags = v }

// GetEnginePowerICE returns the EnginePowerICE of *CarStatusData
func (data *CarStatusData) GetEnginePowerICE() float32 { return data.EnginePowerICE }

// SetEnginePowerICE stores the EnginePowerICE of *CarStatusData
func (data *CarStatusData) SetEnginePowerICE(v float32) { data.EnginePowerICE = v }

// GetEnginePowerMGUK returns the EnginePowerMGUK of *CarStatusData
func (data *CarStatusData) GetEnginePowerMGUK() float32 { return data.EnginePowerMGUK }

// SetEnginePowerMGUK stores the EnginePowerMGUK of *CarStatusData
func (data *CarStatusData) SetEnginePowerMGUK(v float32) { data.EnginePowerMGUK = v }

// GetErsStoreEnergy returns the ErsStoreEnergy of *CarStatusData
func (data *CarStatusData) GetErsStoreEnergy() float32 { return data.ErsStoreEnergy }

// SetErsStoreEnergy stores the ErsStoreEnergy of *CarStatusData
func (data *CarStatusData) SetErsStoreEnergy(v float32) { data.ErsStoreEnergy = v }

// GetErsDeployMode returns the ErsDeployMode of *CarStatusData
func (data *CarStatusData) GetErsDeployMode() uint8 { return data.ErsDeployMode }

// SetErsDeployMode stores the ErsDeployMode of *CarStatusData
func (data *CarStatusData) SetErsDeployMode(v uint8) { data.ErsDeployMode = v }

// GetErsHarvestedThisLapMGUK returns the ErsHarvestedThisLapMGUK of *CarStatusData
func (data *CarStatusData) GetErsHarvestedThisLapMGUK() float32 { return data.ErsHarvestedThisLapMGUK }

// SetErsHarvestedThisLapMGUK stores the ErsHarvestedThisLapMGUK of *CarStatusData
func (data *CarStatusData) SetErsHarvestedThisLapMGUK(v float32) { data.ErsHarvestedThisLapMGUK = v }

// GetErsHarvestedThisLapMGUH returns the ErsHarvestedThisLapMGUH of *CarStatusData
func (data *CarStatusData) GetErsHarvestedThisLapMGUH() float32 { return data.ErsHarvestedThisLapMGUH }

// SetErsHarvestedThisLapMGUH stores the ErsHarvestedThisLapMGUH of *CarStatusData
func (data *CarStatusData) SetErsHarvestedThisLapMGUH(v float32) { data.ErsHarvestedThisLapMGUH = v }

// GetErsHarvestedLimitPerLap returns the ErsHarvestedLimitPerLap of *CarStatusData
func (data *CarStatusData) GetErsHarvestedLimitPerLap() float32 { return data.ErsHarvestedLimitPerLap }

// SetErsHarvestedLimitPerLap stores the ErsHarvestedLimitPerLap of *CarStatusData
func (data *CarStatusData) SetErsHarvestedLimitPerLap(v float32) { data.ErsHarvestedLimitPerLap = v }

// GetErsDeployedThisLap returns the ErsDeployedThisLap of *CarStatusData
func (data *CarStatusData) GetErsDeployedThisLap() float32 { return data.ErsDeployedThisLap }

// SetErsDeployedThisLap stores the ErsDeployedThisLap of *CarStatusData
func (data *CarStatusData) SetErsDeployedThisLap(v float32) { data.ErsDeployedThisLap = v }

// GetNetworkPaused returns the NetworkPaused of *CarStatusData
func (data *CarStatusData) GetNetworkPaused() uint8 { return data.NetworkPaused }

// SetNetworkPaused stores the NetworkPaused of *CarStatusData
func (data *CarStatusData) SetNetworkPaused(v uint8) { data.NetworkPaused = v }

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
