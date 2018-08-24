package testHelper

import (
	"sort"
	"testing"
)

func AssertSliceEqual(a []int, b []int, t *testing.T) {
	if a == nil && b != nil || a != nil && b == nil {
		Fatal(a, b, t)
	}
	if len(a) != len(b) {
		Fatal(a, b, t)
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			Fatal(a, b, t)
		}
	}
}

func AssertSliceStringEqual(a []string, b []string, t *testing.T) {
	if a == nil && b != nil || a != nil && b == nil {
		Fatal(a, b, t)
	}
	if len(a) != len(b) {
		Fatal(a, b, t)
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			Fatal(a, b, t)
		}
	}
}

func AssertTrue(a bool, t *testing.T) {
	if !a {
		t.Fatal("Expect: true, result: false")
	}
}

func AssertFalse(a bool, t *testing.T) {
	if a {
		t.Fatal("Expect: false, result: true")
	}
}

func AssertEqual(a interface{}, b interface{}, t *testing.T) {
	if a != b {
		Fatal(a, b, t)
	}
}

func Fatal(expect interface{}, result interface{}, t *testing.T) {
	t.Fatalf("Expect: %d, result: %d", expect, result)
}
