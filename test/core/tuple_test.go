package core

import (
	"testing"

	"example.com/ray-tracer/test"

	"example.com/ray-tracer/core"
)

func Test_tuple_with_w_equal_1_is_a_point(t *testing.T) {
	tuple := core.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
	if !tuple.IsPoint() {
		t.Fatalf("%v must be a point", tuple)
	}
	if tuple.IsVector() {
		t.Fatalf("%v cannot be a vector", tuple)
	}
}

func Test_tuple_with_w_equal_0_is_a_vector(t *testing.T) {
	tuple := core.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0}
	if tuple.IsPoint() {
		t.Fatalf("%v cannot be a point", tuple)
	}
	if !tuple.IsVector() {
		t.Fatalf("%v must be a vector", tuple)
	}
}

func Test_point_creates_a_tuple_with_w_equal_1(t *testing.T) {
	result := core.NewPoint(4, -4, 3)
	expected := core.Tuple{X: 4, Y: -4, Z: 3, W: 1}
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_vector_creates_a_tuple_with_w_equal_0(t *testing.T) {
	result := core.NewVector(4, -4, 3)
	expected := core.Tuple{X: 4, Y: -4, Z: 3, W: 0}
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_adding_two_tuples(t *testing.T) {
	t1 := core.NewPoint(3, -2, 5)
	t2 := core.NewVector(-2, 3, 1)
	result := t1.Plus(t2)
	expected := core.NewPoint(1, 1, 6)
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_subtracting_two_points(t *testing.T) {
	t1 := core.NewPoint(3, 2, 1)
	t2 := core.NewPoint(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.NewVector(-2, -4, -6)
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_subtracting_a_vector_from_a_point(t *testing.T) {
	t1 := core.NewPoint(3, 2, 1)
	t2 := core.NewVector(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.NewPoint(-2, -4, -6)
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_subtracting_two_vectors(t *testing.T) {
	t1 := core.NewVector(3, 2, 1)
	t2 := core.NewVector(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.NewVector(-2, -4, -6)
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_subtracting_a_vector_from_the_zero_vector(t *testing.T) {
	zero := core.NewVector(0, 0, 0)
	v := core.NewVector(1, -2, 3)
	result := zero.Minus(v)
	expected := core.NewVector(-1, 2, -3)
	test.With(t).That(result).IsEqualTo(expected)
}

func Test_negating_a_tuple(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test.With(t).That(tuple.Negate()).IsEqualTo(core.Tuple{X: -1, Y: 2, Z: -3, W: 4})
}

func Test_multiplying_a_tuple_by_a_scalar(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test.With(t).That(tuple.Multiply(3.5)).IsEqualTo(core.Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14})
}

func Test_multiplying_a_tuple_by_a_fraction(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test.With(t).That(tuple.Multiply(0.5)).IsEqualTo(core.Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2})
}

func Test_dividing_a_tuple_by_a_scalar(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test.With(t).That(tuple.Divide(2)).IsEqualTo(core.Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2})
}
