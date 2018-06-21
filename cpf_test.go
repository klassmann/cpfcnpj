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

func TestCPFTypeIsValidWithValidNumber(t *testing.T) {
	var cpf CPF = "71656686759"

	if !cpf.IsValid() {
		t.Errorf("Invalid result: false")
	}
}

func TestCPFTypeIsValidWithInvalidNumber(t *testing.T) {
	var cpf CPF = "71656686734"

	if cpf.IsValid() {
		t.Errorf("Invalid result: true")
	}
}

func TestCPFTypeStringFormattingWithValidNumber(t *testing.T) {
	var cpf CPF = "71656686759"
	r := cpf.String()

	if r != "716.566.867-59" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestCPFTypeStringFormattingWithInvalidNumber(t *testing.T) {
	var cpf CPF = "71656686734"
	r := cpf.String()

	if r != "71656686734" {
		t.Errorf("Invalid result: %v", r)
	}
}

func TestCPFNewCPF(t *testing.T) {
	cpf := NewCPF("716.566.867-59")

	if cpf != "71656686759" {
		t.Errorf("Invalid result: %v", cpf)
	}
}
