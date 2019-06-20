package point3

import "testing"

func TestPoint3_String(t *testing.T) {
	p := Point3{}
	if p.String() != "(0.000000,0.000000,0.000000)" {
		t.Fail()
	}
}

func TestPoint3_Dimension(t *testing.T) {
	p := Point3{0,1,2}
	if p.Dimension(0) != 0 {
		t.Fail()
	}

	if p.Dimension(1) != 1 {
		t.Fail()
	}

	if p.Dimension(2) != 2 {
		t.Fail()
	}

	if p.Dimension(3) != 2 {
		t.Fail()
	}
}
