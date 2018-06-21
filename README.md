# cpfcnpj
Brazilian taxpayer identification (CPF e CNPJ) - validation package in Golang.

It is an essential package for include validation for these numbers in your application.

Eg: It can be used in a payment system or account system.

## Information about CPF and CNPJ

- [Article - TransferWise](https://transferwise.com/gb/blog/cpf-cnpj-meaning-brazil)
- [Wikipedia - CNPJ](https://en.wikipedia.org/wiki/CNPJ)
- [Wikipedia - CPF](https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas)

## Installation
Use the `go tool` for do that:
```bash
$ go get github.com/klassmann/cpfcnpj
```

## Usage

You need to import the *package* `github.com/klassmann/cpfcnpj` like that:

```go
import (
    "github.com/klassmann/cpfcnpj"
)
```

## Example

```go
package main

import (
	"fmt"

	"github.com/klassmann/cpfcnpj"
)

func main() {

	// This will initialize a new CNPJ and clear the string
	cnpj := cpfcnpj.NewCNPJ("70.082.591/0001-79")

	// Verifies if it is a valid document
	if !cnpj.IsValid() {
		panic(fmt.Errorf("it is not a valid document: %v", cnpj))
	}

	// Cleaned output
	fmt.Printf("%v\n", cnpj)
	// Output: 70082591000179

	// Formatted output
	fmt.Println(cnpj.String())
	// Output: 70.082.591/0001-79
}
```

### Function: `cpfcnpj.ValidateCPF(s string) bool`
Validates the CPF document
```go
import (
    "fmt"
    "github.com/klassmann/cpfcnpj"
)

/// CPF
func ValidatingCPF() {
    r := cpfcnpj.ValidateCPF("71656686734")

    if !r {
        fmt.Printf("The document is invalid.")
    }
}
```

### Function: `cpfcnpj.ValidateCNPJ(s string) bool`
Validates the CNPJ document
```go
/// CNPJ
func ValidatingCNPJ() {
    r := cpfcnpj.ValidateCNPJ("64493884000146")

    if !r {
        fmt.Printf("The document is invalid.")
    }
}
```

### Function: `cpfcnpj.Clean(s string) string`
Clean the formatted document
```go
cpf := cpfcnpj.Clean("111.222.333-99")
cnpj := cpfcnpj.Clean("10.963.268/0001-82")
```

## License
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat)](https://opensource.org/licenses/Apache-2.0)