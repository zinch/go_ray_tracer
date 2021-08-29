package graphics

type Canvas struct {
	Width  int
	Height int
	colors []Color
}

func NewCanvas(width int, height int) Canvas {
	colors := make([]Color, width*height)
	return Canvas{width, height, colors}
}

func (canvas *Canvas) WritePixel(x int, y int, color Color) {
	canvas.colors[canvas.pixelIndex(x, y)] = color
}

func (canvas *Canvas) PixelAt(x int, y int) Color {
	return canvas.colors[canvas.pixelIndex(x, y)]
}

func (canvas *Canvas) pixelIndex(x int, y int) int {
	return x*canvas.Width + y
}
