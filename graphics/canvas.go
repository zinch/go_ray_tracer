package graphics

import (
	"bufio"
	"fmt"
	"io"
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

func (c *canvas) WritePixel(x int, y int, color Color) {
	c.colors[c.pixelIndex(x, y)] = color
}

func (c *canvas) PixelAt(x int, y int) Color {
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

	if _, err = writer.WriteString("P3\n"); err != nil {
		return err
	}

	if _, err = writer.WriteString(fmt.Sprintf("%d %d", c.Width, c.Height)); err != nil {
		return err
	}

	return nil
}
