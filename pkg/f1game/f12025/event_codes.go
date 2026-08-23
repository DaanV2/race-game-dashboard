package f12025

type EventCode string

const (
	EVENT_CODE_BUTTON_STATUS        EventCode = "BUTN" // Button status changed
	EVENT_CODE_CHEQUERED_FLAG       EventCode = "CHQF" // The chequered flag has been waved
	EVENT_CODE_COLLISION            EventCode = "COLL" // Collision between two vehicles has occurred
	EVENT_CODE_DRIVE_THROUGH_SERVED EventCode = "DTSV" // Drive through penalty served
	EVENT_CODE_DRS_DISABLED         EventCode = "DRSD" // Race control have disabled DRS
	EVENT_CODE_DRS_ENABLED          EventCode = "DRSE" // Race control have enabled DRS
	EVENT_CODE_FASTEST_LAP          EventCode = "FTLP" // When a driver achieves the fastest lap
	EVENT_CODE_FLASHBACK            EventCode = "FLBK" // Flashback activated
	EVENT_CODE_LIGHTS_OUT           EventCode = "LGOT" // Lights out
	EVENT_CODE_OVERTAKE             EventCode = "OVTK" // Overtake occurred
	EVENT_CODE_PENALTY_ISSUED       EventCode = "PENA" // A penalty has been issued – details in event
	EVENT_CODE_RACE_WINNER          EventCode = "RCWN" // The race winner is announced
	EVENT_CODE_RED_FLAG             EventCode = "RDFL" // Red flag shown
	EVENT_CODE_RETIREMENT           EventCode = "RTMT" // When a driver retires
	EVENT_CODE_SAFETY_CAR           EventCode = "SCAR" // Safety car event – details in event
	EVENT_CODE_SESSION_ENDED        EventCode = "SEND" // Sent when the session ends
	EVENT_CODE_SESSION_STARTED      EventCode = "SSTA" // Sent when the session starts
	EVENT_CODE_SPEED_TRAP_TRIGGERED EventCode = "SPTP" // Speed trap has been triggered by fastest speed
	EVENT_CODE_START_LIGHTS         EventCode = "STLG" // Start lights – number shown
	EVENT_CODE_STOP_GO_SERVED       EventCode = "SGSV" // Stop go penalty served
	EVENT_CODE_TEAM_MATE_IN_PITS    EventCode = "TMPT" // Your team mate has entered the pits
)
