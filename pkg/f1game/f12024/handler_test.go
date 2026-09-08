package f12024_test

import "github.com/daanv2/race-game-dashboard/pkg/f1game/f12024"

func ExamplePacketHandler() {
	handler := &f12024.PacketHandler{}

	handler.Motion.Register(func(data *f12024.PacketMotionData) {
		// Do stuff with the packet
	})
	handler.Motion.Register(func(data *f12024.PacketMotionData) {
		// Do stuff with the packet async
		go func() {
			// ...
		}()
	})

	// Byte parses:
	var buf []byte

	handler.Ingest(buf)
}
