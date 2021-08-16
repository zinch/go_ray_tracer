package math

import m "math"

func AreEqual(x float32, y float32) bool {
	return m.Abs(float64(x)-float64(y)) < 1e-6
}
