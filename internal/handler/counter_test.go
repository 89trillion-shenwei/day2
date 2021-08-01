package handler

import (
	"testing"
)


func TestCalculate(t *testing.T) {
	s := "324+24*3243/100"
	ret := Calculate(Infix2ToPostfix(Str2Strs(s)))
	var a int
	a = 1102
	if ret == a {
		t.Log("succes")
	} else {
		t.Log("failed")
		t.Error("failed")
	}
}
