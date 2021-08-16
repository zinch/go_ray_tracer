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

func Point(x float32, y float32, z float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1}
}
