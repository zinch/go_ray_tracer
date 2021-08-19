package test

import (
	"math"
	"testing"

	"example.com/ray-tracer/core"
)

func Test_computing_the_magnitude_of_a_vector(t *testing.T) {
	vector := core.Vector(1, 0, 0)
	test(t).that(vector.Magnitude()).isEqualTo(1)

	vector = core.Vector(0, 1, 0)
	test(t).that(vector.Magnitude()).isEqualTo(1)

	vector = core.Vector(0, 0, 1)
	test(t).that(vector.Magnitude()).isEqualTo(1)

	vector = core.Vector(1, 2, 3)
	test(t).that(vector.Magnitude()).isEqualTo(math.Sqrt(14))

	vector = core.Vector(-1, -2, -3)
	test(t).that(vector.Magnitude()).isEqualTo(math.Sqrt(14))
}

func Test_normalizing_a_vector(t *testing.T) {
	vector := core.Vector(4, 0, 0)
	test(t).that(vector.Normalize()).isEqualTo(core.Vector(1, 0, 0))

	vector = core.Vector(1, 2, 3)
	test(t).that(vector.Normalize()).isEqualTo(core.Vector(0.267261, 0.534522, 0.801783))
}

func Test_magnitude_of_a_normalized_vector_is_1(t *testing.T) {
	vector := core.Vector(1, 2, 3)
	test(t).that(vector.Normalize().Magnitude()).isEqualTo(1)
}

func Test_dot_product(t *testing.T) {
	vec1 := core.Vector(1, 2, 3)
	vec2 := core.Vector(2, 3, 4)
	test(t).that(vec1.Dot(vec2)).isEqualTo(20)
}
