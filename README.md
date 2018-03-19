# cpfcnpj
Brazilian taxpayer registry identification document validation package


## Installation

    go get https://github.com/klassmann/cpfcnpj
  
## Usage
    
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

/// CNPJ
func ValidatingCNPJ() {
    r := cpfcnpj.ValidateCNPJ("64493884000146")

    if !r {
        fmt.Printf("The document is invalid.")
    }
}
```

