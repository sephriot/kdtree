package point3

import "fmt"

type Point3 struct {
	X float64
	Y float64
	Z float64
}

func (p Point3) Dimensions() int {
	return 3
}

func (p Point3) Dimension(i int) float64 {
	switch i {
	case 0:
		return p.X
	case 1:
		return p.Y
	case 2:
		return p.Z
	}

	return p.Z
}

func (p Point3) String() string {
	return fmt.Sprintf("(%f,%f,%f)", p.X, p.Y, p.Z)
}