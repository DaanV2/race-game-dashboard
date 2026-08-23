package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataSafetyCar struct {
	SafetyCarType uint8 // 0 = No Safety Car, 1 = Full Safety Car  2 = Virtual Safety Car, 3 = Formation Lap Safety Car
	EventType     uint8 // 0 = Deployed, 1 = Returning, 2 = Returned  3 = Resume Race
}

// GetSafetyCarType returns the SafetyCarType of *SafetyCar
func (data *EventDataSafetyCar) GetSafetyCarType() uint8 { return data.SafetyCarType }

// SetSafetyCarType stores the SafetyCarType of *SafetyCar
func (data *EventDataSafetyCar) SetSafetyCarType(v uint8) { data.SafetyCarType = v }

// GetEventType returns the EventType of *SafetyCar
func (data *EventDataSafetyCar) GetEventType() uint8 { return data.EventType }

// SetEventType stores the EventType of *SafetyCar
func (data *EventDataSafetyCar) SetEventType(v uint8) { data.EventType = v }

func (data *EventDataSafetyCar) Parse(reader *xbinary.LittleEndianReader) {
	data.SafetyCarType = reader.ReadUint8()
	data.EventType = reader.ReadUint8()

}
