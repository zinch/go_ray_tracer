package test

import (
	"fmt"
	"reflect"
	"testing"

	"example.com/ray-tracer/core"
	"example.com/ray-tracer/graphics"
	"example.com/ray-tracer/math"
)

func test(t *testing.T) Assert {
	return Assert{Test: t}
}

type Assert struct {
	Test *testing.T
}

var tupleType = reflect.TypeOf((*core.Tuple)(nil))

func (assert Assert) that(value interface{}) Matcher {
	if tuple, ok := value.(core.Tuple); ok {
		return TupleMatcher{Assert: assert, Tuple: tuple}
	}
	if color, ok := value.(graphics.Color); ok {
		return ColorMatcher{Assert: assert, Color: color}
	}
	if number, ok := value.(float64); ok {
		return Float64Matcher{Assert: assert, Value: number}
	}
	panic(fmt.Sprintf("%s does not have a matcher", reflect.TypeOf(value).Name()))
}

type Matcher interface {
	isEqualTo(value interface{})
}

type TupleMatcher struct {
	Assert
	Tuple core.Tuple
}

func (matcher TupleMatcher) isEqualTo(value interface{}) {
	expected := value.(core.Tuple)
	t := matcher.Assert.Test
	if !matcher.Tuple.Equals(expected) {
		t.Fatalf("%v must be equal to %v", matcher.Tuple, expected)
	}
}

type Float64Matcher struct {
	Assert
	Value float64
}

func (matcher Float64Matcher) isEqualTo(value interface{}) {
	var expected float64
	if intVal, ok := value.(int); ok {
		expected = float64(intVal)
	} else {
		expected = value.(float64)
	}
	t := matcher.Assert.Test
	if !math.AreEqual(matcher.Value, expected) {
		t.Fatalf("%f must be close to to %f", matcher.Value, expected)
	}
}

type ColorMatcher struct {
	Assert
	Color graphics.Color
}

func (matcher ColorMatcher) isEqualTo(value interface{}) {
	expected := value.(graphics.Color)
	t := matcher.Assert.Test
	if !matcher.Color.Equals(expected) {
		t.Fatalf("%v must be equal to %v", matcher.Color, expected)
	}
}
