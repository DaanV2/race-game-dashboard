package f12026s

// The game types, mapped to golangs type
type (
	// uint8  = uint8   // Unsigned 8-bit integer
	// int8   = int8    // Signed 8-bit integer
	// uint16 = uint16  // Unsigned 16-bit integer
	// int16  = int16   // Signed 16-bit integer
	// uint32 = uint32  // Unsigned 32-bit integer
	// uint64 = uint64  // Unsigned 64-bit integer
	//go:fix inline
	float = float32 // Floating point (32-bit)
	//go:fix inline
	double = float64 // Double-precision floating point (64-bit)
	//go:fix inline
	char = uint8 // Character
)
