package xsync

import "sync"

type Pool[T any] struct {
	inner *sync.Pool
}

func NewPool[T any](newfn func() T) *Pool[T] {
	inner := &sync.Pool{}

	if newfn != nil {
		inner.New = func() any {
			return newfn()
		}
	}

	return &Pool[T]{
		inner: inner,
	}
}

func (p *Pool[T]) Get() T {
	var zero T

	v := p.inner.Get()
	if v == nil {
		return zero
	}

	if vt, ok := v.(T); ok {
		return vt
	}

	return zero
}

func (p *Pool[T]) Put(data T) {
	p.inner.Put(data)
}
