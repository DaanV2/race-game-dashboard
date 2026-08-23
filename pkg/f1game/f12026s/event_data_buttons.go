package f12026s

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type EventDataButtons struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed  currently - see appendices
}

// GetButtonStatus returns the ButtonStatus of *Buttons
func (data *EventDataButtons) GetButtonStatus() uint32 { return data.ButtonStatus }

// SetButtonStatus stores the ButtonStatus of *Buttons
func (data *EventDataButtons) SetButtonStatus(v uint32) { data.ButtonStatus = v }

func (data *EventDataButtons) Parse(reader *xbinary.LittleEndianReader) {
	data.ButtonStatus = reader.ReadUint32()
}
