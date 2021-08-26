package test

import (
	"testing"

	"example.com/ray-tracer/graphics"
)

func Test_that_color_is_a_red_green_blue_tuple(t *testing.T) {
	c := graphics.NewColor(-0.5, 0.4, 1.7)
	test(t).that(c.Red()).isEqualTo(-0.5)
	test(t).that(c.Green()).isEqualTo(0.4)
	test(t).that(c.Blue()).isEqualTo(1.7)
}

func Test_adding_colors(t *testing.T) {
	c1 := graphics.NewColor(0.9, 0.6, 0.75)
	c2 := graphics.NewColor(0.7, 0.1, 0.25)
	test(t).that(c1.Plus(c2)).isEqualTo(graphics.NewColor(1.6, 0.7, 1))
}

func Test_subtracting_colors(t *testing.T) {
	c1 := graphics.NewColor(0.9, 0.6, 0.75)
	c2 := graphics.NewColor(0.7, 0.1, 0.25)
	test(t).that(c1.Minus(c2)).isEqualTo(graphics.NewColor(0.2, 0.5, 0.5))
}

func Test_multiplying_color_by_scalar(t *testing.T) {
	c := graphics.NewColor(0.2, 0.3, 0.4)
	test(t).that(c.Multiply(2)).isEqualTo(graphics.NewColor(0.4, 0.6, 0.8))
}

func Test_blending_colors(t *testing.T) {
	c1 := graphics.NewColor(1, 0.2, 0.4)
	c2 := graphics.NewColor(0.9, 1, 0.1)
	test(t).that(c1.BlendWith(c2)).isEqualTo(graphics.NewColor(0.9, 0.2, 0.04))
}
