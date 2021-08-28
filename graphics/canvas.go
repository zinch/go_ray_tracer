package graphics

type Canvas struct {
	Width  int
	Height int
}

func NewCanvas(width int, height int) Canvas {
	return Canvas{width, height}
}
