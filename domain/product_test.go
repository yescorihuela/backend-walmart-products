package domain

import "testing"

var image string = "https://i.picsum.photos/id/557/200/300.jpg?hmac=eC86bsSOhqQjoHHnj3yzH5wMTIY9S3ys6cQjU1_QLGc"

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

func TestToDTOWithPalindromeCriteria(t *testing.T) {
	mockProduct := NewProduct(121, "adda", "Coffee Maker", image, 32000)
	// Check discount application
	mockDTO := mockProduct.ToDTO("adda")
	expected := discount(mockProduct.Price, APPLIED_DISCOUNT)
	got := mockDTO.Price
	if got != expected {
		t.Errorf("Invalid applied discount")
	}

	// Check has_discount flag
	expectedHasDiscount := true
	gotHasDiscount := mockDTO.HasDiscount
	if gotHasDiscount != expectedHasDiscount {
		t.Errorf("Invalid evaluation of criteria")
	}
}

func TestToDTOWithoutPalindromeCriteria(t *testing.T) {
	mockProduct := NewProduct(121, "adda", "Coffee Maker", image, 32000)
	// Check no discount application
	mockDTO := mockProduct.ToDTO("hola")
	expectedDiscount := discount(mockProduct.Price, APPLIED_DISCOUNT)
	gotDiscount := mockDTO.Price
	if gotDiscount == expectedDiscount {
		t.Errorf("Invalid applied discount")
	}

	// Check has_discount flag
	expectedHasDiscount := false
	gotHasDiscount := mockDTO.HasDiscount
	if gotHasDiscount != expectedHasDiscount {
		t.Errorf("Invalid evaluation of criteria")
	}

}
