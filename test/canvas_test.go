package test

import (
	"bytes"
	"testing"

	"example.com/ray-tracer/graphics"
)

func Test_creating_a_canvas(t *testing.T) {
	c := graphics.NewCanvas(10, 20)
	test(t).that(c.Width).isEqualTo(10)
	test(t).that(c.Height).isEqualTo(20)
}

func Test_writing_pixels_to_canvas(t *testing.T) {
	c := graphics.NewCanvas(10, 20)
	red := graphics.NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	test(t).that(c.PixelAt(2, 3)).isEqualTo(red)
}

func Test_saving_in_PPM_format(t *testing.T) {
	c := graphics.NewCanvas(5, 3)
	var buf bytes.Buffer
	c.ToPPM(&buf)
	expected := `P3`
	test(t).that(buf.String()).isEqualTo(expected)
}
