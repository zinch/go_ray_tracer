package test

import (
	"testing"

	"example.com/ray-tracer/graphics"
)

func Test_creating_a_canvas(t *testing.T) {
	c := graphics.NewCanvas(10, 20)
	test(t).that(c.Width).isEqualTo(10)
	test(t).that(c.Height).isEqualTo(20)
}
