package ctrl

import "testing"

func TestCouter(t *testing.T) {
	s :="2132  2+132/13*13232"
	ret := Couter(Analy(s))
	var a int
	a+=132
	a/=13
	a*=13232
	a+=21322
	if ret==a {
		t.Log("succes")
	}else {
		t.Log("failed")
		t.Error("failed")
	}
}