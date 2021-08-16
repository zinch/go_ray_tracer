package core

import (
	"testing"

	"example.com/ray-tracer/core"
)

func Test_tuple_with_w_equal_1_is_a_point(t *testing.T) {
	var tuple = core.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
	if !tuple.IsPoint() {
		t.Fatalf("%v must be a point", tuple)
	}
    if tuple.IsVector() {
		t.Fatalf("%v cannot be a vector", tuple)
    }
}

func Test_tuple_with_w_equal_0_is_a_vector(t *testing.T) {
	var tuple = core.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0}
	if tuple.IsPoint() {
		t.Fatalf("%v cannot be a point", tuple)
	}
    if !tuple.IsVector() {
		t.Fatalf("%v must be a vector", tuple)
    }
}
