package xsync

import (
	"iter"
	"sync"
)

type Map[K comparable, T any] struct {
	data *sync.Map
}

func NewMap[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		data: &sync.Map{},
	}
}

func (m *Map[K, T]) Get(key K) (reuslt T, ok bool) {
	var tmp any

	tmp, ok = m.data.Load(key)
	if !ok {
		return
	}
	result, ok := tmp.(T)

	return result, ok
}

func (m *Map[K, T]) Set(key K, value T) {
	m.data.Store(key, value)
}
func (m *Map[K, T]) Values() iter.Seq2[K, T] {
	return func(yield func(K, T) bool) {
		m.data.Range(func(keytmp, valuetmp any) bool {
			key, ok := keytmp.(K)
			if !ok {
				return true
			}

			value, ok := valuetmp.(T)
			if !ok {
				return true
			}

			return yield(key, value)
		})
	}
}
