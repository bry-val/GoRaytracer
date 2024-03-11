package main

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	var tests = []struct {
		x, y, z float64
		want    *Tuple
	}{
		{1, 2, 3, &Tuple{1.0, 2.0, 3.0, 1.0}},
		{4, -3, 2, &Tuple{4.0, -3.0, 2.0, 1.0}},
		{1, 2, 5, &Tuple{1.0, 2.0, 5.0, 1.0}},
		{-1, 10, -3, &Tuple{-1.0, 10.0, -3.0, 1.0}},
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
