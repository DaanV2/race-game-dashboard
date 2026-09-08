package f12026s_test

import "github.com/daanv2/race-game-dashboard/pkg/f1game/f12026s"

func ExamplePacketHandler() {
	handler := &f12026s.PacketHandler{}

	handler.Motion.Register(func(data *f12026s.PacketMotionData) {
		// Do stuff with the packet
	})
	handler.Motion.Register(func(data *f12026s.PacketMotionData) {
		// Do stuff with the packet async
		go func() {
			// ...
		}()
	})

	// Byte parses:
	var buf []byte

	handler.Ingest(buf)
}
