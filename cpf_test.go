package cpfcnpj

import (
	"testing"
)

func TestValidCPFWithValidNumber1(t *testing.T) {
	r := ValidateCPF("64844696793")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithValidNumber2(t *testing.T) {
	r := ValidateCPF("62641322846")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithValidNumber3(t *testing.T) {
	r := ValidateCPF("87195726037")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithValidNumber4(t *testing.T) {
	r := ValidateCPF("71656686759")

	if !r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithInvalidNumber1(t *testing.T) {
	r := ValidateCPF("71656686753")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithInvalidNumber2(t *testing.T) {
	r := ValidateCPF("71656686739")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestValidCPFWithInvalidNumber3(t *testing.T) {
	r := ValidateCPF("71656686734")

	if r {
		t.Errorf("Invalid result: %v", r)
	}
}
