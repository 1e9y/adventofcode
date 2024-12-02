package challenge

type TestCase[T any] struct {
	Input string
	Want  []T
}

func (t TestCase[T]) Name() string {
	if len(t.Input) <= 6 {
		return t.Input
	}

	return t.Input[:6]
}
