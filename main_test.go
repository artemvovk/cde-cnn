package main

import "testing"

func TestStringMaker(t *testing.T) {
	teststring := GetAString()
	if teststring != "hello world" {
		t.Errorf("Did not get the right string. Got %s", GetAString())
	}
}
