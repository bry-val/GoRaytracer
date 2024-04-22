package main

import (
	"fmt"
	"math"
)

const EPSILON float64 = 0.00001

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) String() string {
	kind := "Invalid"
	if t.w == 1.0 {
		kind = "Point"
	} else if t.w == 0.0 {
		kind = "Vector"
	}
	return fmt.Sprintf("x: %f\ny: %f\nz: %f\nType: %s\n", t.x, t.y, t.z, kind)
}

// a tuple is a point when w is 1
// a tuple is a vector when w is 0
func point(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1.0}
}

func vector(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0.0}
}

func add(a, b Tuple) Tuple {
	return Tuple{a.x + b.x,
		a.y + b.y,
		a.z + b.z,
		a.w + b.w,
	}
}
func subtract(a, b Tuple) Tuple {
	return Tuple{a.x - b.x,
		a.y - b.y,
		a.z - b.z,
		a.w - b.w,
	}
}

func equal(a, b Tuple) bool {
	return math.Abs(a.x-b.x) < EPSILON &&
		math.Abs(a.y-b.y) < EPSILON &&
		math.Abs(a.z-b.z) < EPSILON &&
		math.Abs(a.w-b.w) < EPSILON
}

func main() {
	fmt.Println(subtract(point(1, 2, 3), point(1, 2, 3)))
	fmt.Println(subtract(point(1, 2, 3), vector(1, 2, 3)))
	fmt.Println(subtract(vector(1, 2, 3), vector(1, 2, 3)))
	fmt.Println(subtract(vector(1, 3, 3), point(1, 2, 3)))

}
