package collections

func Reduce[T any](list []T, fn func(accum, item T) T, initialValue T) T {
	res := initialValue

	for _, v := range list {
		res = fn(res, v)
	}
	return res
}
