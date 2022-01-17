package domain

import "testing"

func TestPalindrome(t *testing.T) {
	expected := true
	got := isPalindrome("11211")
	if got != expected {
		t.Errorf("Expected value (%v) is not the same as got value (%v)", expected, got)
	}
}

func TestDiscount(t *testing.T) {
	expected := float32(1000.00)
	got := discount(2000, 50)
	if got != expected {
		t.Errorf("Expected value (%v) is not the same as got value (%v)", expected, got)
	}
}
