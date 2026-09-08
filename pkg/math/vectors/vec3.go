package vectors

import (
	"golang.org/x/exp/constraints"
)

type Vec3[TVal constraints.Float | constraints.Integer] struct {
	X TVal
	Y TVal
	Z TVal
}

func NewVec3[TVal constraints.Float | constraints.Integer](x, y, z TVal) Vec3[TVal] {
	return Vec3[TVal]{X: x, Y: y, Z: z}
}
