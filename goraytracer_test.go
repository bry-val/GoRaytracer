package main

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	var tests = []struct {
		x, y, z float64
		want    Tuple
	}{
		{1, 2, 3, Tuple{1.0, 2.0, 3.0, 1.0}},
		{4, -3, 2, Tuple{4.0, -3.0, 2.0, 1.0}},
		{1, 2, 5, Tuple{1.0, 2.0, 5.0, 1.0}},
		{-1, 10, -3, Tuple{-1.0, 10.0, -3.0, 1.0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f,%f", tt.x, tt.y, tt.z)
		t.Run(testname, func(t *testing.T) {
			ans := point(tt.x, tt.y, tt.z)
			if ans != tt.want {
				t.Errorf("\n\n***GOT*** \n%s\n\n***WANT***\n%s", ans, tt.want)
			}
		})
	}
}
func TestVector(t *testing.T) {
	var tests = []struct {
		x, y, z float64
		want    Tuple
	}{
		{1, 2, 3, Tuple{1.0, 2.0, 3.0, 0.0}},
		{4, -3, 2, Tuple{4.0, -3.0, 2.0, 0.0}},
		{1, 2, 5, Tuple{1.0, 2.0, 5.0, 0.0}},
		{-1, 10, -3, Tuple{-1.0, 10.0, -3.0, 0.0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f,%f", tt.x, tt.y, tt.z)
		t.Run(testname, func(t *testing.T) {
			ans := vector(tt.x, tt.y, tt.z)
			if ans != tt.want {
				t.Errorf("\n\n***GOT*** \n%s\n\n***WANT***\n%s", ans, tt.want)
			}
		})
	}
}
func TestEqual(t *testing.T) {
	var tests = []struct {
		a, b Tuple
		want bool
	}{
		{point(1, 2, 3), point(1, 2, 3), true},
		{point(1.1, 2, 3.1), point(1.4, 2, 3.001), false},
		//z component subtraction == 0.000049, test should return false
		{vector(1.0000001, 2.0, 3.000001), vector(1.0, 2.000001, 3.00005), false},
		{vector(1, 2, 3), point(1, 2, 3), false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := equal(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("\n\n***GOT*** \n%t\n\n***WANT***\n%t", ans, tt.want)
			}
		})
	}
}
func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b Tuple
		want Tuple
	}{
		{point(1, 2, 3), point(1, 2, 3), Tuple{2, 4, 6, 2}},
		{point(1, 2, 3), vector(1, 2, 3), Tuple{2, 4, 6, 1}},
		{vector(1, 2, 3), point(1, 2, 3), Tuple{2, 4, 6, 1}},
		{vector(1, 2, 3), vector(1, 2, 3), Tuple{2, 4, 6, 0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("\n\n***GOT*** \n%s\n\n***WANT***\n%s", ans, tt.want)
			}
		})
	}
}
func TestSubtract(t *testing.T) {
	var tests = []struct {
		a, b Tuple
		want Tuple
	}{
		{point(1, 2, 3), point(1, 2, 3), Tuple{0, 0, 0, 0}},
		{point(1, 2, 3), vector(1, 2, 3), Tuple{0, 0, 0, 1}},
		{vector(1, 2, 3), point(1, 2, 3), Tuple{0, 0, 0, -1}},
		{vector(1, 2, 3), vector(1, 2, 3), Tuple{0, 0, 0, 0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := subtract(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("\n\n***GOT*** \n%s\n\n***WANT***\n%s", ans, tt.want)
			}
		})
	}
}
