package test

import (
	"bytes"
	"strings"
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

func Test_constructing_the_PPM_header(t *testing.T) {
	c := graphics.NewCanvas(5, 3)
	var buf bytes.Buffer
	c.ToPPM(&buf)
	header := strings.Join(strings.Split(buf.String(), "\n")[:3], "\n")
	expected := `P3
5 3
255`
	test(t).that(header).isEqualTo(expected)
}

func Test_constructing_the_PPM_pixel_data(t *testing.T) {
	c := graphics.NewCanvas(5, 3)
	c.WritePixel(0, 0, graphics.NewColor(1.5, 0, 0))
	c.WritePixel(2, 1, graphics.NewColor(0, 0.5, 0))
	c.WritePixel(4, 2, graphics.NewColor(-0.5, 0, 1))
	var buf bytes.Buffer
	c.ToPPM(&buf)

	pixelData := strings.Join(strings.Split(buf.String(), "\n")[3:], "\n")
	expected := ` 255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	test(t).that(pixelData).isEqualTo(expected)
}
