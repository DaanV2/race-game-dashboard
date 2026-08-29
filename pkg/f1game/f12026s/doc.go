// This package contains all the models needed to read the telemetry data from f1 game: f12026
//
// Either use the [PacketHandler] (see example) to setup a pipeline of events or manually:
//
// Example:
//
//	var header PacketHeader
//	header.Parse(reader)
//
//	// Do stuff with other packets here
package f12026s
