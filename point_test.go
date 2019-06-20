package kdtree

import (
	"github.com/sephriot/kdtree/point2"
	"github.com/sephriot/kdtree/point3"
	"math"
	"testing"
)

func TestEquals(t *testing.T)  {
	p1 := point2.Point2{}
	p2 := point2.Point2{}
	if !equals(p1,p2) {
		t.Fail()
	}
	p3 := point3.Point3{}
	if equals(p1, p3) {
		t.Fail()
	}
	p4 := point3.Point3{}
	if !equals(p3, p4) {
		t.Fail()
	}
}

func TestDistance2(t *testing.T) {
	p1 := point2.Point2{}
	p2 := point2.Point2{}
	if distance(p1, p2) > floatEqualityThreshold {
		t.Fail()
	}
	p3 := point2.Point2{X:1}
	if distance(p1, p3) - 1.0 > floatEqualityThreshold {
		t.Fail()
	}
	p4 := point3.Point3{}
	if distance2(p1, p4) < math.MaxFloat64 {
		t.Fail()
	}
}