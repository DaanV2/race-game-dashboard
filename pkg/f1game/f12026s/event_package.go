package f12026s

type EventCode string

const (
	EVENT_CODE_SESSION_STARTED      EventCode = "SSTA" // Sent when the session starts
	EVENT_CODE_SESSION_ENDED        EventCode = "SEND" // Sent when the session ends
	EVENT_CODE_FASTEST_LAP          EventCode = "FTLP" // When a driver achieves the fastest lap
	EVENT_CODE_RETIREMENT           EventCode = "RTMT" // When a driver retires
	EVENT_CODE_DRS_ENABLED          EventCode = "DRSE" // Race control have enabled DRS
	EVENT_CODE_DRS_DISABLED         EventCode = "DRSD" // Race control have disabled DRS
	EVENT_CODE_TEAM_MATE_IN_PITS    EventCode = "TMPT" // Your team mate has entered the pits
	EVENT_CODE_CHEQUERED_FLAG       EventCode = "CHQF" // The chequered flag has been waved
	EVENT_CODE_RACE_WINNER          EventCode = "RCWN" // The race winner is announced
	EVENT_CODE_PENALTY_ISSUED       EventCode = "PENA" // A penalty has been issued – details in event
	EVENT_CODE_SPEEDVTRAP_TRIGGERED EventCode = "SPTP" // Speed trap has been triggered by fastest speed
	EVENT_CODE_START_LIGHTS         EventCode = "STLG" // Start lights – number shown
	EVENT_CODE_LIGHTS_OUT           EventCode = "LGOT" // Lights out
	EVENT_CODE_DRIVE_THROUGH_SERVED EventCode = "DTSV" // Drive through penalty served
	EVENT_CODE_STOP_GO_SERVED       EventCode = "SGSV" // Stop go penalty served
	EVENT_CODE_FLASHBACK            EventCode = "FLBK" // Flashback activated
	EVENT_CODE_BUTTON_STATUS        EventCode = "BUTN" // Button status changed
	EVENT_CODE_RED_FLAG             EventCode = "RDFL" // Red flag shown
	EVENT_CODE_OVERTAKE             EventCode = "OVTK" // Overtake occurred
	EVENT_CODE_SAFETY_CAR           EventCode = "SCAR" // Safety car event – details in event
	EVENT_CODE_COLLISION            EventCode = "COLL" // Collision between two vehicles has occurred
)

type eventDataDetailsUnion interface {
	eventDataDetailsUnion()
}

// The event details packet is different for each type of event.
// Make sure only the correct type is interpreted.
type EventDataDetails struct {
	data eventDataDetailsUnion
}

var (
	_ eventDataDetailsUnion = &FastestLap{}
	_ eventDataDetailsUnion = &Retirement{}
	_ eventDataDetailsUnion = &DRSDisabled{}
	_ eventDataDetailsUnion = &TeamMateInPits{}
	_ eventDataDetailsUnion = &RaceWinner{}
	_ eventDataDetailsUnion = &Penalty{}
	_ eventDataDetailsUnion = &SpeedTrap{}
	_ eventDataDetailsUnion = &StartLights{}
	_ eventDataDetailsUnion = &DriveThroughPenaltyServed{}
	_ eventDataDetailsUnion = &StopGoPenaltyServed{}
	_ eventDataDetailsUnion = &Flashback{}
	_ eventDataDetailsUnion = &Buttons{}
	_ eventDataDetailsUnion = &Overtake{}
	_ eventDataDetailsUnion = &SafetyCar{}
	_ eventDataDetailsUnion = &Collision{}
)

type FastestLap struct {
	vehicleIdx uint8   // Vehicle index of car achieving fastest lap
	lapTime    float32 // Lap time is in seconds
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (f *FastestLap) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Retirement struct {
	vehicleIdx uint8 // Vehicle index of car retiring
	// Reason - 0 = invalid, 1 = retired, 2 = finished
	// 3 = terminal damage, 4 = inactive, 5 = not enough laps completed
	// 6 = black flagged, 7 = red flagged, 8 = mechanical failure
	// 9 = session skipped, 10 = session simulated
	reason uint8
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (r *Retirement) eventDataDetailsUnion() {
	panic("unimplemented")
}

type DRSDisabled struct {
	reason uint8 // 0 = Wet track, 1 = Safety car deployed, 2 = Red flag
	// 3 = Min lap not reached
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (d *DRSDisabled) eventDataDetailsUnion() {
	panic("unimplemented")
}

type TeamMateInPits struct {
	vehicleIdx uint8 // Vehicle index of team mate
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (t *TeamMateInPits) eventDataDetailsUnion() {
	panic("unimplemented")
}

type RaceWinner struct {
	vehicleIdx uint8 // Vehicle index of the race winner
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (r *RaceWinner) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Penalty struct {
	penaltyType      uint8 // Penalty type – see Appendices
	infringementType uint8 // Infringement type – see Appendices
	vehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	otherVehicleIdx  uint8 // Vehicle index of the other car involved
	time             uint8 // Time gained, or time spent doing action in seconds
	lapNum           uint8 // Lap the penalty occurred on
	placesGained     uint8 // Number of places gained by this
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (p *Penalty) eventDataDetailsUnion() {
	panic("unimplemented")
}

type SpeedTrap struct {
	vehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
	speed                      float32 // Top speed achieved in kilometres per hour
	isOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
	isDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
	fastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest in this session
	fastestSpeedInSession      float32 // Speed of the vehicle that is the fastest in this session
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (s *SpeedTrap) eventDataDetailsUnion() {
	panic("unimplemented")
}

type StartLights struct {
	numLights uint8 // Number of lights showing
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (s *StartLights) eventDataDetailsUnion() {
	panic("unimplemented")
}

type DriveThroughPenaltyServed struct {
	vehicleIdx uint8 // Vehicle index of the vehicle serving drive through
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (d *DriveThroughPenaltyServed) eventDataDetailsUnion() {
	panic("unimplemented")
}

type StopGoPenaltyServed struct {
	vehicleIdx uint8   // Vehicle index of the vehicle serving stop go
	stopTime   float32 // Time spent serving stop go in seconds
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (s *StopGoPenaltyServed) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Flashback struct {
	flashbackFrameIdentifier uint32  // Frame identifier flashed back to
	flashbackSessionTime     float32 // Session time flashed back to
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (f *Flashback) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Buttons struct {
	buttonStatus uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (b *Buttons) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Overtake struct {
	overtakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	beingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (o *Overtake) eventDataDetailsUnion() {
	panic("unimplemented")
}

type SafetyCar struct {
	safetyCarType uint8 // 0 = No Safety Car, 1 = Full Safety Car 2 = Virtual Safety Car, 3 = Formation Lap Safety Car
	eventType     uint8 // 0 = Deployed, 1 = Returning, 2 = Returned 3 = Resume Race
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (s *SafetyCar) eventDataDetailsUnion() {
	panic("unimplemented")
}

type Collision struct {
	vehicle1Idx uint8 // Vehicle index of the first vehicle involved in the collision
	vehicle2Idx uint8 // Vehicle index of the second vehicle involved in the collision
	severity    uint8 // Severity of the collision - 0 = low, 1 = medium, 2 = high
}

// eventDataDetailsUnion implements [eventDataDetailsUnion].
func (c *Collision) eventDataDetailsUnion() {
	panic("unimplemented")
}

type PacketEventData struct {
	m_header PacketHeader // Header

	m_eventStringCode [4]uint8         // Event string code, see below
	m_eventDetails    EventDataDetails // Event details - should be interpreted differently for each type
}
