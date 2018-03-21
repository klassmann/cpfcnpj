package cpfcnpj

import "fmt"

var (
	cnpjFirstDigitTable  = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

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
