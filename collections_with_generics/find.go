package collections

func Find[T any](list []T, testFn func(item T) bool) (val T, found bool) {
	for _, v := range list {
		if testFn(v) {
			return v, true
		}
	}
	return
}
