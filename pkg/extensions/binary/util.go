package xbinary

func Readx4[T any](callbf func() T) [4]T {
	return [4]T{
		callbf(),
		callbf(),
		callbf(),
		callbf(),
	}
}

func Readx8[T any](callbf func() T) [8]T {
	return [8]T{
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
	}
}

func Readx12[T any](callbf func() T) [12]T {
	return [12]T{
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
	}
}

func Readx24[T any](callbf func() T) [24]T {
	return [24]T{
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
	}
}

func Readx32[T any](callbf func() T) [32]T {
	return [32]T{
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
		callbf(),
	}
}
