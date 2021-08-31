package canvas

import (
	"bytes"
	"strings"
	"testing"

	"example.com/ray-tracer/graphics"
	"example.com/ray-tracer/test"
)

func Test_creating_a_canvas(t *testing.T) {
	c := graphics.NewCanvas(10, 20)
	test.With(t).That(c.Width).IsEqualTo(10)
	test.With(t).That(c.Height).IsEqualTo(20)
}

func Test_writing_pixels_to_canvas(t *testing.T) {
	c := graphics.NewCanvas(10, 20)
	red := graphics.NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	test.With(t).That(c.PixelAt(2, 3)).IsEqualTo(red)
}

func Test_constructing_the_PPM_header(t *testing.T) {
	c := graphics.NewCanvas(5, 3)
	var buf bytes.Buffer
	c.ToPPM(&buf)
	header := strings.Join(strings.Split(buf.String(), "\n")[:3], "\n")
	expected := `P3
5 3
255`
	test.With(t).That(header).IsEqualTo(expected)
}

func Test_constructing_the_PPM_pixel_data(t *testing.T) {
	c := graphics.NewCanvas(5, 3)
	c.WritePixel(0, 0, graphics.NewColor(1.5, 0, 0))
	c.WritePixel(2, 1, graphics.NewColor(0, 0.5, 0))
	c.WritePixel(4, 2, graphics.NewColor(-0.5, 0, 1))
	var buf bytes.Buffer
	c.ToPPM(&buf)

	pixelData := strings.Join(strings.Split(buf.String(), "\n")[3:], "\n")
	expected := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	test.With(t).That(pixelData).IsEqualTo(expected)
}

func Test_splitting_long_lines_in_PPM_files(t *testing.T) {
	c := graphics.NewCanvas(10, 2)
	for y := 0; y < 10; y++ {
		for x := 0; x < 2; x++ {
			c.WritePixel(y, x, graphics.NewColor(1, 0.8, 0.6))
		}
	}
	var buf bytes.Buffer
	c.ToPPM(&buf)

	pixelData := strings.Join(strings.Split(buf.String(), "\n")[3:7], "\n")
	expected := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`
	test.With(t).That(pixelData).IsEqualTo(expected)
}
