package safe

func MustAt[T any](xs []T, i int) T {
	if i < 0 || i >= len(xs) {
		panic("index out of range")
	}
	return xs[i]
}
