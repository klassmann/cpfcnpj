// Package cpfcnpj provides functions for validate CPF and CNPJ, the Brazilian taxpayer registry identification document
package cpfcnpj

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CPFFormatPattern is the pattern for string replacement
	// with Regex
	CPFFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
)

// CPF type definition
type CPF string

// NewCPF is a helper function to convert and clean a new CPF
// from a string
func NewCPF(s string) CPF {
	return CPF(Clean(s))
}

// IsValid returns if CPF is a valid CPF document
func (c *CPF) IsValid() bool {
	return ValidateCPF(string(*c))
}

// String returns a formatted CPF document
// 000.000.000-00
func (c *CPF) String() string {

	str := string(*c)

	if !c.IsValid() {
		return str
	}

	expr, err := regexp.Compile(CPFFormatPattern)

	if err != nil {
		return str
	}

	return expr.ReplaceAllString(str, "$1.$2.$3-$4")
}

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
