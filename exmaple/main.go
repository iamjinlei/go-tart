package main

import (
	"fmt"

	"github.com/iamjinlei/go-tart"
)

func main() {
	sma := tart.NewSma(5)
	for i := 0; i < 20; i++ {
		fmt.Printf("sma[%v] = %.4f\n", i, sma.Update(float64(i%7)))
	}
}
