package cpfcnpj

import (
	"fmt"
	"regexp"
)

var (
	cnpjFirstDigitTable  = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CNPJFormatPattern is the pattern for string replacement
	// with Regex
	CNPJFormatPattern string = `([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})`
)

// CNPJ type definition
type CNPJ string

// NewCNPJ is a helper function to convert and clean a new CNPJ
// from a string
func NewCNPJ(s string) CNPJ {
	return CNPJ(Clean(s))
}

// IsValid returns if CNPJ is a valid CNPJ document
func (c *CNPJ) IsValid() bool {
	return ValidateCNPJ(string(*c))
}

// String returns a formatted CNPJ document
// 00.000.000/0001-00
func (c *CNPJ) String() string {

	str := string(*c)

	if !c.IsValid() {
		return str
	}

	expr, err := regexp.Compile(CNPJFormatPattern)

	if err != nil {
		return str
	}

	return expr.ReplaceAllString(str, "$1.$2.$3/$4-$5")
}

// ValidateCNPJ validates a CNPJ document
// You should use without punctuation
func ValidateCNPJ(cnpj string) bool {

	if len(cnpj) != 14 {
		return false
	}

	firstPart := cnpj[:12]
	sum1 := sumDigit(firstPart, cnpjFirstDigitTable)
	rest1 := sum1 % 11
	d1 := 0

	if rest1 >= 2 {
		d1 = 11 - rest1
	}

	secondPart := fmt.Sprintf("%s%d", firstPart, d1)
	sum2 := sumDigit(secondPart, cnpjSecondDigitTable)
	rest2 := sum2 % 11
	d2 := 0

	if rest2 >= 2 {
		d2 = 11 - rest2
	}

	finalPart := fmt.Sprintf("%s%d", secondPart, d2)
	return finalPart == cnpj
}
