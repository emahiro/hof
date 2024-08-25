package hof

import "iter"

func Filter[T any](src []T, fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range src {
			if fn(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Map[T, U any](src []T, fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for _, v := range src {
			if !yield(fn(v)) {
				break
			}
		}
	}
}

func Map2[T comparable, U any](src map[T]U, fn func(T, U) U) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for k, v := range src {
			if !yield(k, fn(k, v)) {
				break
			}
		}
	}
}

func Chunk[T any](src []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i < len(src); i += size {
			end := i + size
			if end > len(src) {
				end = len(src)
			}
			if !yield(src[i:end]) {
				break
			}
		}
	}
}

func Reduce[T, U any](src []T, init U, fn func(U, T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		acc := init
		for _, v := range src {
			acc = fn(acc, v)
		}
		yield(acc)
	}
}

func Reduce2[K comparable, V, U any](src map[K]V, init U, fn func(U, K, V) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		acc := init
		for k, v := range src {
			acc = fn(acc, k, v)
		}
		yield(acc)
	}
}
