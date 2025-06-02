package collections

func Reduce[A, B any](list []A, fn func(B, A) B, initialValue B) B {
	res := initialValue

	for _, v := range list {
		res = fn(res, v)
	}
	return res
}
