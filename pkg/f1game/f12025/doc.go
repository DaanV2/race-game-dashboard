// This package contains all the models needed to read the telemetry data from f1 game: f12025
//
// Either use the [PacketHandler] (see example) to setup a pipeline of events or manually:
//
// Example:
//
//	var header PacketHeader
//	header.Parse(reader)
//
//	// Do stuff with other packets here
package f12025
