package cpfcnpj

import "testing"

func TestValidCNPJWithInvalidNumber1(t *testing.T) {
	r := ValidateCNPJ("73687174000158")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCNPJWithInvalidNumber2(t *testing.T) {
	r := ValidateCNPJ("62009392000103")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCNPJWithInvalidNumber3(t *testing.T) {
	r := ValidateCNPJ("22796729000145")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCNPJWithValidNumber1(t *testing.T) {
	r := ValidateCNPJ("73687174000148")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCNPJWithValidNumber2(t *testing.T) {
	r := ValidateCNPJ("62009392000107")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCNPJWithValidNumber3(t *testing.T) {
	r := ValidateCNPJ("22796729000159")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}
