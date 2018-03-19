package cpfcnpj

import (
	"strconv"
)

func sumCPF(cpf string) int {
	sum := 0
	tamanho := len(cpf)
	for i, c := range cpf {
		s := string(c)
		n, err := strconv.Atoi(s)
		if err == nil {
			pos := (tamanho + 1) - i
			s := pos * n
			sum += s
		}
	}

	return sum
}

// ValidateCPF validates a CPF document
func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumCPF(firstPart)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumCPF(secondPart)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	final := secondPart + strconv.Itoa(d2)
	return final == cpf
}
