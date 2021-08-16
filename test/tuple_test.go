package core

import (
    "testing"
    "example.com/ray-tracer/core"
)

func TestTupleCreation(t *testing.T) {
    var t1 = core.Tuple{1, 2, 3, 1}
    var t2 = core.Tuple{1, 2, 3, 1}
    if t1.W != t2.W {
        t.Errorf("Tuples not equal")
    }
}
