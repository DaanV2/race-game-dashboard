package generics

import (
	"golang.org/x/exp/constraints"
)

func CopySlice[TDest, TSrc constraints.Integer | constraints.Float](dst []TDest, src []TSrc) {
	l := min(len(dst), len(src))

	for i := range l {
		dst[i] = TDest(src[i])
	}
}
