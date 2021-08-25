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
