package methods

type Number struct {
	// Public
	Value int
	// Private
	value int
}

func (n Number) GetNumber() int {
	n.value = 2
	return n.Value + n.value
}
