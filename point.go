package kdtree

import "math"

const (
	floatEqualityThreshold = 1e-9
)

type Point interface {
	Dimensions() int
	Dimension(i int) float64
}

func equals(p1, p2 Point) bool {
	if p1.Dimensions() != p2.Dimensions() {
		return false
	}
	for i := 0; i < p1.Dimensions(); i++ {
		if math.Abs(p1.Dimension(i)-p2.Dimension(i)) > floatEqualityThreshold {
			return false
		}
	}
	return true
}

func distance2(p1, p2 Point) float64 {
	if p1.Dimensions() != p2.Dimensions() {
		return math.MaxFloat64
	}
	d := 0.0
	for i := 0; i < p1.Dimensions(); i++ {
		d += (p1.Dimension(i)-p2.Dimension(i))*(p1.Dimension(i)-p2.Dimension(i))
	}

	return d
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(distance2(p1,p2))
}
