package point2

import "fmt"

type Point2 struct {
	X float64
	Y float64
}

func (p Point2) Dimensions() int {
	return 2
}

func (p Point2) Dimension(i int) float64 {
	if i == 0 {
		return p.X
	}
	return p.Y
}

func (p Point2) String() string {
	return fmt.Sprintf("(%f,%f)", p.X, p.Y)
}