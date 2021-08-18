package core

import "example.com/ray-tracer/math"

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

func (t Tuple) Equals(other Tuple) bool {
	return math.AreEqual(t.X, other.X) &&
		math.AreEqual(t.Y, other.Y) &&
		math.AreEqual(t.Z, other.Z) &&
		t.W == other.W
}

func Point(x float32, y float32, z float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1}
}

func Vector(x float32, y float32, z float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0}
}
