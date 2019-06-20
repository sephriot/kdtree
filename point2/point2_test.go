package point2

import "testing"

func TestPoint2_String(t *testing.T) {
	p := Point2{}
	if p.String() != "(0.000000,0.000000)" {
		t.Log(p)
		t.Fail()
	}
}