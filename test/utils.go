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
		return TupleMatcher{Assert: assert, Value: tuple}
	}
	if color, ok := value.(graphics.Color); ok {
		return ColorMatcher{Assert: assert, Value: color}
	}
	if number, ok := value.(float64); ok {
		return Float64Matcher{Assert: assert, Value: number}
	}
	if number, ok := value.(int); ok {
		return IntMatcher{Assert: assert, Value: number}
	}
	if str, ok := value.(string); ok {
		return StringMatcher{Assert: assert, Value: str}
	}
	panic(fmt.Sprintf("%s does not have a matcher", reflect.TypeOf(value).Name()))
}

type Matcher interface {
	isEqualTo(value interface{})
}

type TupleMatcher struct {
	Assert
	Value core.Tuple
}

func (matcher TupleMatcher) isEqualTo(value interface{}) {
	expected := value.(core.Tuple)
	t := matcher.Assert.Test
	if !matcher.Value.Equals(expected) {
		t.Fatalf("%v must be equal to %v", matcher.Value, expected)
	}
}

type Float64Matcher struct {
	Assert
	Value float64
}

func (matcher Float64Matcher) isEqualTo(value interface{}) {
	t := matcher.Assert.Test
	var intVal int
	var expected float64
	var ok bool
	if intVal, ok = value.(int); ok {
		expected = float64(intVal)
	} else if expected, ok = value.(float64); !ok {
		t.Fatalf("%T is not float64", value)
	}
	if !math.AreEqual(matcher.Value, expected) {
		t.Fatalf("%f must be close to to %f", matcher.Value, expected)
	}
}

type IntMatcher struct {
	Assert
	Value int
}

func (matcher IntMatcher) isEqualTo(value interface{}) {
	t := matcher.Assert.Test
	var expected int
	var ok bool
	if expected, ok = value.(int); !ok {
		t.Fatalf("%T is not int", value)
	}
	if expected != matcher.Value {
		t.Fatalf("%d must be equal to %d", matcher.Value, expected)
	}
}

type ColorMatcher struct {
	Assert
	Value graphics.Color
}

func (matcher ColorMatcher) isEqualTo(value interface{}) {
	expected := value.(graphics.Color)
	t := matcher.Assert.Test
	if !matcher.Value.Equals(expected) {
		t.Fatalf("%v must be equal to %v", matcher.Value, expected)
	}
}

type StringMatcher struct {
	Assert
	Value string
}

func (matcher StringMatcher) isEqualTo(value interface{}) {
	expected := value.(string)
	t := matcher.Assert.Test
	if matcher.Value != expected {
		t.Fatalf("%v must be equal to %v", matcher.Value, expected)
	}
}
