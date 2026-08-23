package f12025

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

type ButtonFlag uint32

const (
	BF_Cross_or_A        ButtonFlag = 0x00000001 // Cross or A
	BF_Triangle_or_Y     ButtonFlag = 0x00000002 // Triangle or Y
	BF_Circle_or_B       ButtonFlag = 0x00000004 // Circle or B
	BF_Square_or_X       ButtonFlag = 0x00000008 // Square or X
	BF_D_pad_Left        ButtonFlag = 0x00000010 // D-pad Left
	BF_D_pad_Right       ButtonFlag = 0x00000020 // D-pad Right
	BF_D_pad_Up          ButtonFlag = 0x00000040 // D-pad Up
	BF_D_pad_Down        ButtonFlag = 0x00000080 // D-pad Down
	BF_Options_or_Menu   ButtonFlag = 0x00000100 // Options or Menu
	BF_L1_or_LB          ButtonFlag = 0x00000200 // L1 or LB
	BF_R1_or_RB          ButtonFlag = 0x00000400 // R1 or RB
	BF_L2_or_LT          ButtonFlag = 0x00000800 // L2 or LT
	BF_R2_or_RT          ButtonFlag = 0x00001000 // R2 or RT
	BF_Left_Stick_Click  ButtonFlag = 0x00002000 // Left Stick Click
	BF_Right_Stick_Click ButtonFlag = 0x00004000 // Right Stick Click
	BF_Right_Stick_Left  ButtonFlag = 0x00008000 // Right Stick Left
	BF_Right_Stick_Right ButtonFlag = 0x00010000 // Right Stick Right
	BF_Right_Stick_Up    ButtonFlag = 0x00020000 // Right Stick Up
	BF_Right_Stick_Down  ButtonFlag = 0x00040000 // Right Stick Down
	BF_Special           ButtonFlag = 0x00080000 // Special  = 0x00080000
	BF_UDP_Action_1      ButtonFlag = 0x00100000 // UDP Action 1
	BF_UDP_Action_2      ButtonFlag = 0x00200000 // UDP Action 2
	BF_UDP_Action_3      ButtonFlag = 0x00400000 // UDP Action 3
	BF_UDP_Action_4      ButtonFlag = 0x00800000 // UDP Action 4
	BF_UDP_Action_5      ButtonFlag = 0x01000000 // UDP Action 5
	BF_UDP_Action_6      ButtonFlag = 0x02000000 // UDP Action 6
	BF_UDP_Action_7      ButtonFlag = 0x04000000 // UDP Action 7
	BF_UDP_Action_8      ButtonFlag = 0x08000000 // UDP Action 8
	BF_UDP_Action_9      ButtonFlag = 0x10000000 // UDP Action 9
	BF_UDP_Action_10     ButtonFlag = 0x20000000 // UDP Action 10
	BF_UDP_Action_11     ButtonFlag = 0x40000000 // UDP Action 11
	BF_UDP_Action_12     ButtonFlag = 0x80000000 // UDP Action 12
)

func (b ButtonFlag) HasButtons(bts ButtonFlag) bool {
	return (b & bts) == bts
}

type EventDataButtons struct {
	ButtonStatus ButtonFlag // Bit flags specifying which buttons are being pressed  currently - see appendices
}

// GetButtonStatus returns the ButtonStatus of *Buttons
func (data *EventDataButtons) GetButtonStatus() ButtonFlag { return data.ButtonStatus }

// SetButtonStatus stores the ButtonStatus of *Buttons
func (data *EventDataButtons) SetButtonStatus(v ButtonFlag) { data.ButtonStatus = v }

func (data *EventDataButtons) Parse(reader *xbinary.LittleEndianReader) {
	data.ButtonStatus = ButtonFlag(reader.ReadUint32())
}
