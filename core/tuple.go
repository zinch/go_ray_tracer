package core

import (
	"math"

	rayMath "example.com/ray-tracer/math"
)

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
	return rayMath.AreEqual(t.X, other.X) &&
		rayMath.AreEqual(t.Y, other.Y) &&
		rayMath.AreEqual(t.Z, other.Z) &&
		rayMath.AreEqual(t.W, other.W)
}

func (t Tuple) Plus(other Tuple) Tuple {
	return Tuple{t.X + other.X, t.Y + other.Y, t.Z + other.Z, t.W + other.W}
}

func (t Tuple) Minus(other Tuple) Tuple {
	return t.Plus(other.Negate())
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t Tuple) Divide(s float64) Tuple {
	return t.Multiply(1 / s)
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tuple) Normalize() Tuple {
	magnitude := t.Magnitude()
	return Tuple{t.X / magnitude, t.Y / magnitude, t.Z / magnitude, t.W / magnitude}
}

func (t Tuple) Dot(other Tuple) float64 {
	return t.X*other.X + t.Y*other.Y + t.Z*other.Z + t.W*other.W
}

func (t Tuple) Cross(other Tuple) Tuple {
	return NewVector(
		t.Y*other.Z-t.Z*other.Y,
		t.Z*other.X-t.X*other.Z,
		t.X*other.Y-t.Y*other.X,
	)
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
