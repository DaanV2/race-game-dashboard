package f12025_test

import "github.com/daanv2/race-game-dashboard/pkg/f1game/f12025"

func ExamplePacketHandler() {
	handler := &f12025.PacketHandler{}

	handler.Motion.Register(func(data *f12025.PacketMotionData) {
		// Do stuff with the packet
	})
	handler.Motion.Register(func(data *f12025.PacketMotionData) {
		// Do stuff with the packet async
		go func() {
			// ...
		}()
	})

	// Byte parses:
	var buf []byte

	handler.Ingest(buf)
}
