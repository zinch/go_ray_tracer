package math

import m "math"

func AreEqual(x float64, y float64) bool {
	return m.Abs(float64(x)-float64(y)) < 1e-6
}
