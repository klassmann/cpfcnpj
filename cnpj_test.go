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

func TestCNPJTypeIsValidWithValidNumber(t *testing.T) {
	var cnpj CNPJ = "22796729000159"

	if !cnpj.IsValid() {
		t.Errorf("Invalid result: false")
	}
}

func TestCNPJTypeIsValidWithInvalidNumber(t *testing.T) {
	var cnpj CNPJ = "62009392000103"

	if cnpj.IsValid() {
		t.Errorf("Invalid result: true")
	}
}

func TestCNPJTypeStringFormattingWithValidNumber(t *testing.T) {
	var cnpj CNPJ = "22796729000159"
	r := cnpj.String()

	if r != "22.796.729/0001-59" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestCNPJTypeStringFormattingWithInvalidNumber(t *testing.T) {
	var cnpj CNPJ = "62009392000103"
	r := cnpj.String()

	if r != "62009392000103" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestCNPJNewCNPJ(t *testing.T) {
	cnpj := NewCNPJ("22.796.729/0001-59")

	if cnpj != "22796729000159" {
		t.Errorf("Invalid result: %v", cnpj)
	}
}
