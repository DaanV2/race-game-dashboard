package f12026s

type EventCode string

const (
	EVENT_CODE_BUTTON_STATUS        = "BUTN" // Button status changed
	EVENT_CODE_CHEQUERED_FLAG       = "CHQF" // The chequered flag has been waved
	EVENT_CODE_COLLISION            = "COLL" // Collision between two vehicles has occurred
	EVENT_CODE_DRIVE_THROUGH_SERVED = "DTSV" // Drive through penalty served
	EVENT_CODE_DRS_DISABLED         = "DRSD" // Race control have disabled DRS
	EVENT_CODE_DRS_ENABLED          = "DRSE" // Race control have enabled DRS
	EVENT_CODE_FASTEST_LAP          = "FTLP" // When a driver achieves the fastest lap
	EVENT_CODE_FLASHBACK            = "FLBK" // Flashback activated
	EVENT_CODE_LIGHTS_OUT           = "LGOT" // Lights out
	EVENT_CODE_OVERTAKE             = "OVTK" // Overtake occurred
	EVENT_CODE_PENALTY_ISSUED       = "PENA" // A penalty has been issued – details in event
	EVENT_CODE_RACE_WINNER          = "RCWN" // The race winner is announced
	EVENT_CODE_RED_FLAG             = "RDFL" // Red flag shown
	EVENT_CODE_RETIREMENT           = "RTMT" // When a driver retires
	EVENT_CODE_SAFETY_CAR           = "SCAR" // Safety car event – details in event
	EVENT_CODE_SESSION_ENDED        = "SEND" // Sent when the session ends
	EVENT_CODE_SESSION_STARTED      = "SSTA" // Sent when the session starts
	EVENT_CODE_SPEED_TRAP_TRIGGERED = "SPTP" // Speed trap has been triggered by fastest speed
	EVENT_CODE_START_LIGHTS         = "STLG" // Start lights – number shown
	EVENT_CODE_STOP_GO_SERVED       = "SGSV" // Stop go penalty served
	EVENT_CODE_TEAM_MATE_IN_PITS    = "TMPT" // Your team mate has entered the pits
)
