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
