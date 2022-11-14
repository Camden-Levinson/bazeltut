package go_hello_world

import "testing"

func TestAdd(t *testing.T) {
	expected := 3
	actual := add(1, 2)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
func TestSub(t *testing.T) {
	expected := 0
	actual := subtract(2, 2)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
func TestMul(t *testing.T) {
	expected := 40
	actual := multiple(10, 4)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
func TestDiv(t *testing.T) {
	expected := 5
	actual := divide(60, 12)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
