package greet

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "hello" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions", result)
	}
}
