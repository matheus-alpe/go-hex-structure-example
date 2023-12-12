package main

import (
	"fmt"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/core/arithmetic"
	"github.com/matheus-alpe/go-hex-structure-example/internal/ports"
)

func main() {
    // ports
    var core ports.ArithmeticPort = arithmetic.NewAdapter()

    fmt.Println(core)
}
