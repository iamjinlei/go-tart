package main

import (
	"fmt"

	"github.com/iamjinlei/go-tart"
)

func main() {
	sma := tart.NewSma(5)
	for i := 0; i < 20; i++ {
		val := sma.Update(float64(i % 7))
		if sma.Valid() {
			fmt.Printf("sma[%v] = %.4f\n", i, val)
		} else {
			fmt.Printf("sma[%v] = unavail.\n", i)
		}
	}
}
