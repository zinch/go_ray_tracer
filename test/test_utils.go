package test

import "math"

func areEqual(x float32, y float32) bool {
	return math.Abs(float64(x)-float64(y)) < 1e-6
}
