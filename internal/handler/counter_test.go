package handler

import (
	"testing"
)

func TestCouter(t *testing.T) {
	s := "1+1"
	e1, e2, err := Analy(s)
	if err != nil {
		t.Log(err.Error())
		t.Error("failed")
	}
	ret := Couter(e1, e2)
	var a int
	a = 2
	if ret == a {
		t.Log("succes")
	} else {
		t.Log("failed")
		t.Error("failed")
	}
}
