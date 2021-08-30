package core

import (
	"math"
	"testing"

	"example.com/ray-tracer/core"
	"example.com/ray-tracer/test"
)

func Test_computing_the_magnitude_of_a_vector(t *testing.T) {
	vector := core.NewVector(1, 0, 0)
	test.With(t).That(vector.Magnitude()).IsEqualTo(1)

	vector = core.NewVector(0, 1, 0)
	test.With(t).That(vector.Magnitude()).IsEqualTo(1)

	vector = core.NewVector(0, 0, 1)
	test.With(t).That(vector.Magnitude()).IsEqualTo(1)

	vector = core.NewVector(1, 2, 3)
	test.With(t).That(vector.Magnitude()).IsEqualTo(math.Sqrt(14))

	vector = core.NewVector(-1, -2, -3)
	test.With(t).That(vector.Magnitude()).IsEqualTo(math.Sqrt(14))
}

func Test_normalizing_a_vector(t *testing.T) {
	vector := core.NewVector(4, 0, 0)
	test.With(t).That(vector.Normalize()).IsEqualTo(core.NewVector(1, 0, 0))

	vector = core.NewVector(1, 2, 3)
	test.With(t).That(vector.Normalize()).IsEqualTo(core.NewVector(0.267261, 0.534522, 0.801783))
}

func Test_magnitude_of_a_normalized_vector_is_1(t *testing.T) {
	vector := core.NewVector(1, 2, 3)
	test.With(t).That(vector.Normalize().Magnitude()).IsEqualTo(1)
}

func Test_dot_product(t *testing.T) {
	vec1 := core.NewVector(1, 2, 3)
	vec2 := core.NewVector(2, 3, 4)
	test.With(t).That(vec1.Dot(vec2)).IsEqualTo(20)
}

func Test_cross_product(t *testing.T) {
	vec1 := core.NewVector(1, 2, 3)
	vec2 := core.NewVector(2, 3, 4)
	test.With(t).That(vec1.Cross(vec2)).IsEqualTo(core.NewVector(-1, 2, -1))
	test.With(t).That(vec2.Cross(vec1)).IsEqualTo(core.NewVector(1, -2, 1))
}
