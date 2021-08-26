package graphics

import "example.com/ray-tracer/core"

type Color struct {
	t core.Tuple
}

func NewColor(red float64, green float64, blue float64) Color {
	return Color{t: core.Tuple{X: red, Y: green, Z: blue}}
}

func (c Color) Red() float64 {
	return c.t.X
}

func (c Color) Green() float64 {
	return c.t.Y
}

func (c Color) Blue() float64 {
	return c.t.Z
}

func (c Color) Plus(other Color) Color {
	return NewColor(c.Red()+other.Red(), c.Green()+other.Green(), c.Blue()+other.Blue())
}

func (c Color) Minus(other Color) Color {
	return NewColor(c.Red()-other.Red(), c.Green()-other.Green(), c.Blue()-other.Blue())
}

func (c Color) Multiply(s float64) Color {
	return NewColor(c.Red()*s, c.Green()*s, c.Blue()*s)
}

func (c Color) Equals(other Color) bool {
	return c.t.Equals(other.t)
}
