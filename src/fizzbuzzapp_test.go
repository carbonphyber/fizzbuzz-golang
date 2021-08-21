package fizzbuzzapp

import (
	"testing"
)

// TestFizzbuzzNegative calls fizzbuzzapp.Fizzbuzz with a negative integer, checking
// for an error.
func TestFizzbuzzNegative(t *testing.T) {
	ret, err := Fizzbuzz(-1)
	if err == nil {
		t.Fatalf(`Fizzbuzz(-1) = %v, %v, want nil, error`, ret, err)
	}
}

// TestFizzbuzzTooLarge calls fizzbuzzapp.Fizzbuzz with a very large integer, checking
// for an error.
func TestFizzbuzzTooLarge(t *testing.T) {
	ret, err := Fizzbuzz(10000)
	if err == nil {
		t.Fatalf(`Fizzbuzz(-1) = %v, %v, want nil, error`, ret, err)
	}
}

// TestFizzbuzz1 calls fizzbuzzapp.Fizzbuzz with 1, checking
// for "1", nil.
func TestFizzbuzz1(t *testing.T) {
	ret, err := Fizzbuzz(1)
	expected := "1"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(1) = %v, %v, want "1", nil`, *ret, err)
	}
}

// TestFizzbuzz2 calls fizzbuzzapp.Fizzbuzz with 2, checking
// for "2", nil.
func TestFizzbuzz2(t *testing.T) {
	ret, err := Fizzbuzz(2)
	expected := "2"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(2) = %v, %v, want "2", nil`, *ret, err)
	}
}

// TestFizzbuzz3 calls fizzbuzzapp.Fizzbuzz with 3, checking
// for "fizz", nil.
func TestFizzbuzz3(t *testing.T) {
	ret, err := Fizzbuzz(3)
	expected := "fizz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(3) = %v, %v, want "fizz", nil`, *ret, err)
	}
}

// TestFizzbuzz2 calls fizzbuzzapp.Fizzbuzz with 4, checking
// for "4", nil.
func TestFizzbuzz4(t *testing.T) {
	ret, err := Fizzbuzz(4)
	expected := "4"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(4) = %v, %v, want "4", nil`, *ret, err)
	}
}

// TestFizzbuzz5 calls fizzbuzzapp.Fizzbuzz with 5, checking
// for "2", nil.
func TestFizzbuzz5(t *testing.T) {
	ret, err := Fizzbuzz(5)
	expected := "buzz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(5) = %v, %v, want "buzz", nil`, *ret, err)
	}
}

// TestFizzbuzz6 calls fizzbuzzapp.Fizzbuzz with 6, checking
// for "fizz", nil.
func TestFizzbuzz6(t *testing.T) {
	ret, err := Fizzbuzz(6)
	expected := "fizz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(6) = %v, %v, want "fizz", nil`, *ret, err)
	}
}

// TestFizzbuzz5 calls fizzbuzzapp.Fizzbuzz with 10, checking
// for "buzz", nil.
func TestFizzbuzz10(t *testing.T) {
	ret, err := Fizzbuzz(10)
	expected := "buzz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(10) = %v, %v, want "buzz", nil`, *ret, err)
	}
}

// TestFizzbuzz15 calls fizzbuzzapp.Fizzbuzz with 15, checking
// for "fizzbuzz", nil.
func TestFizzbuzz15(t *testing.T) {
	ret, err := Fizzbuzz(15)
	expected := "fizzbuzz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(15) = %v, %v, want "fizzbuzz", nil`, *ret, err)
	}
}

// TestFizzbuzz30 calls fizzbuzzapp.Fizzbuzz with 30, checking
// for "fizzbuzz", nil.
func TestFizzbuzz30(t *testing.T) {
	ret, err := Fizzbuzz(30)
	expected := "fizzbuzz"
	if *ret != expected || err != nil {
		t.Fatalf(`Fizzbuzz(30) = %v, %v, want "fizzbuzz", nil`, *ret, err)
	}
}
