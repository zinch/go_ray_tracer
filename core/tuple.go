package core

type Tuple struct {
	X float32
	Y float32
	Z float32
	W int
}

func (t Tuple) IsVector() bool {
	return t.W == 0
}

func (t Tuple) IsPoint() bool {
	return t.W == 1
}
