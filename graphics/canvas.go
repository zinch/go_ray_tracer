package graphics

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type canvas struct {
	Width  int
	Height int
	colors []Color
}

func NewCanvas(width int, height int) canvas {
	colors := make([]Color, width*height)
	return canvas{width, height, colors}
}

func (c *canvas) WritePixel(y int, x int, color Color) {
	c.colors[c.pixelIndex(x, y)] = color
}

func (c *canvas) PixelAt(y int, x int) Color {
	return c.colors[c.pixelIndex(x, y)]
}

func (c *canvas) pixelIndex(x int, y int) int {
	return x*c.Width + y
}

func (c *canvas) ToPPM(out io.Writer) (err error) {
	writer := bufio.NewWriter(out)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	header := `P3
%d %d
255
`

	if _, err = writer.WriteString(fmt.Sprintf(header, c.Width, c.Height)); err != nil {
		return err
	}

	cramp := func(x float64) float64 {
		return math.Min(math.Max(x, 0), 1)
	}

	formatColor := func(c Color) string {
		red := int(math.Round(cramp(c.Red()) * 255))
		green := int(math.Round(cramp(c.Green()) * 255))
		blue := int(math.Round(cramp(c.Blue()) * 255))
		return fmt.Sprintf("%d %d %d ", red, green, blue)
	}

	for x := 0; x < c.Height; x++ {
		var b strings.Builder
		for y := 0; y < c.Width; y++ {
			color := c.PixelAt(y, x)
			fmt.Fprint(&b, formatColor(color))
		}
		line := b.String()
		line = line[:len(line)-1]
		if _, err = writer.WriteString(fmt.Sprintln(line)); err != nil {
			return err
		}
	}

	return nil
}
