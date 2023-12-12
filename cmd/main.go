package main

import (
	"fmt"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/core/arithmetic"
	"github.com/matheus-alpe/go-hex-structure-example/internal/ports"
)

func main() {
    // ports
    var core ports.ArithmeticPort = arithmetic.NewAdapter()


    result, err := core.Addition(1, 4)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(result)
}
