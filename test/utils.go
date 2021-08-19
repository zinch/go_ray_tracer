package test

import (
	"testing"

	"example.com/ray-tracer/core"
)

func test(t *testing.T) Assert {
	return Assert{Test: t}
}

type Assert struct {
	Test *testing.T
}

func (assert Assert) that(t core.Tuple) TupleMatcher {
	return TupleMatcher{Assert: assert, Tuple: t}
}

type TupleMatcher struct {
	Assert
	Tuple core.Tuple
}

func (matcher TupleMatcher) isEqualTo(expected core.Tuple) {
	t := matcher.Assert.Test
	if !matcher.Tuple.Equals(expected) {
		t.Fatalf("%v must be equal to %v", matcher.Tuple, expected)
	}
}
