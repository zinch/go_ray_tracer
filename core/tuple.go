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

func Point(x float32, y float32, z float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1}
}

func Vector(x float32, y float32, z float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0}
}

func AreEqual(t1 Tuple, t2 Tuple) bool {
	return math.AreEqual(t1.X, t2.X) ||
		math.AreEqual(t1.Y, t2.Y) ||
		math.AreEqual(t1.Z, t2.Z) ||
		t1.W == t2.W
}
