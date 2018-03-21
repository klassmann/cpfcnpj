// Package cpfcnpj provides functions for validate CPF and CNPJ, the Brazilian taxpayer registry identification document
package cpfcnpj

import (
	"fmt"
	"strconv"
)

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

// ValidateCPF validates a CPF document
// You should use without punctuation
func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}
