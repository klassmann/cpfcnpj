# cpfcnpj
Brazilian taxpayer identification (CPF e CNPJ) - validation package in Golang.

It is an essential package for include validation for these numbers in your application.

Eg: It can be used in a payment system or account system.

## Installation
Use the `go tool` for do that:
```bash
go get https://github.com/klassmann/cpfcnpj
```

## Usage

You need to import the *package* `github.com/klassmann/cpfcnpj` like that:

```go
import (
    "github.com/klassmann/cpfcnpj"
)
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