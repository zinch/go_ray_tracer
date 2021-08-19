package core

import "example.com/ray-tracer/math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
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
		math.AreEqual(t.W, other.W)
}

func (t Tuple) Plus(other Tuple) Tuple {
	return Tuple{X: t.X + other.X, Y: t.Y + other.Y, Z: t.Z + other.Z, W: t.W + other.W}
}

func (t Tuple) Minus(other Tuple) Tuple {
	return t.Plus(other.Negate())
}

func (t Tuple) Negate() Tuple {
	return Tuple{X: -t.X, Y: -t.Y, Z: -t.Z, W: -t.W}
}

func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{X: t.X * s, Y: t.Y * s, Z: t.Z * s, W: t.W * s}
}

func (t Tuple) Divide(s float64) Tuple {
	return t.Multiply(1 / s)
}

func (t Tuple) Magnitude() float64 {
	return 1
}

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0}
}
