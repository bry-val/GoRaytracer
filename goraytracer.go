package main

import "fmt"

const EPSILON float64 = 0.00001

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) String() string {
	kind := "Vector"
	if t.w == 1.0 {
		kind = "Point"
	}
	return fmt.Sprintf("x: %f\ny: %f\nz: %f\nType: %s\n", t.x, t.y, t.z, kind)
}

func point(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1.0}
}

func vector(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0.0}
}

func main() {
	fmt.Println(point(1, 2, 3))
	fmt.Println(vector(3, 1, 3))
	fmt.Println(vector(3, 1, 3))

}
