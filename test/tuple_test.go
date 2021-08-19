package test

import (
	"testing"

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
	result := core.Point(4, -4, 3)
	expected := core.Tuple{X: 4, Y: -4, Z: 3, W: 1}
	test(t).that(result).isEqualTo(expected)
}

func Test_vector_creates_a_tuple_with_w_equal_0(t *testing.T) {
	result := core.Vector(4, -4, 3)
	expected := core.Tuple{X: 4, Y: -4, Z: 3, W: 0}
	test(t).that(result).isEqualTo(expected)
}

func Test_adding_two_tuples(t *testing.T) {
	t1 := core.Point(3, -2, 5)
	t2 := core.Vector(-2, 3, 1)
	result := t1.Plus(t2)
	expected := core.Point(1, 1, 6)
	test(t).that(result).isEqualTo(expected)
}

func Test_subtracting_two_points(t *testing.T) {
	t1 := core.Point(3, 2, 1)
	t2 := core.Point(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.Vector(-2, -4, -6)
	test(t).that(result).isEqualTo(expected)
}

func Test_subtracting_a_vector_from_a_point(t *testing.T) {
	t1 := core.Point(3, 2, 1)
	t2 := core.Vector(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.Point(-2, -4, -6)
	test(t).that(result).isEqualTo(expected)
}

func Test_subtracting_two_vectors(t *testing.T) {
	t1 := core.Vector(3, 2, 1)
	t2 := core.Vector(5, 6, 7)
	result := t1.Minus(t2)
	expected := core.Vector(-2, -4, -6)
	test(t).that(result).isEqualTo(expected)
}

func Test_subtracting_a_vector_from_the_zero_vector(t *testing.T) {
	zero := core.Vector(0, 0, 0)
	v := core.Vector(1, -2, 3)
	result := zero.Minus(v)
	expected := core.Vector(-1, 2, -3)
	test(t).that(result).isEqualTo(expected)
}

func Test_negating_a_tuple(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test(t).that(tuple.Negate()).isEqualTo(core.Tuple{X: -1, Y: 2, Z: -3, W: 4})
}

func Test_multiplying_a_tuple_by_a_scalar(t *testing.T) {
	tuple := core.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	test(t).that(tuple.Multiply(3.5)).isEqualTo(core.Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14})
}
