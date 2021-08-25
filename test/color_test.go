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
