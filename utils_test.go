package cpfcnpj

import "testing"

func TestSumDigitOnlyZeroes(t *testing.T) {
	v := sumDigit("000", []int{1, 2, 3})

	if v != 0 {
		t.Errorf("Invalid result: %d", v)
	}
}

func TestSumDigitOnlyOnes(t *testing.T) {
	v := sumDigit("111", []int{1, 2, 3})

	if v != 6 {
		t.Errorf("Invalid result: %d", v)
	}
}

func TestClean1(t *testing.T) {
	s := Clean("111.222.333-99")

	if s != "11122233399" {
		t.Errorf("Invalid result: %s", s)
	}
}

func TestClean2(t *testing.T) {
	s := Clean("10.963.268/0001-82")

	if s != "10963268000182" {
		t.Errorf("Invalid result: %s", s)
	}
}
