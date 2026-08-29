package generics

import "reflect"

func New[T any]() T {
	var zero T
	typ := reflect.TypeFor[T]() // Get the reflect.Type of T

	// If T is a pointer type (e.g., *MyStruct)
	if typ.Kind() == reflect.Pointer {
		val := reflect.New(typ.Elem())

		return val.Interface().(T)
	}

	return zero
}
