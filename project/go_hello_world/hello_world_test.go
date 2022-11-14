package go_hello_world

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	if actual != 3.0 {
		t.Errorf("Expected %f but got %s", 3.0, actual)
	}
}
func TestSub(t *testing.T) {
	actual := Subtract(2, 2)
	if actual != 0.0 {
		t.Errorf("Expected %f but got %s", 0.0, actual)
	}
}
func TestMul(t *testing.T) {
	actual := Multiple(10, 4)
	if actual != 40.0 {
		t.Errorf("Expected %f but got %s", 40.0, actual)
	}
}
func TestDiv(t *testing.T) {
	actual := Divide(60, 12)
	if actual != 5.0 {
		t.Errorf("Expected %f but got %s", 5.0, actual)
	}
}
func TestGreeter(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
