package ref

func Ref[T any](val T) *T {
	return &val
}
